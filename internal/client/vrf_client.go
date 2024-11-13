package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"1U/config"
	"1U/contract"
	"1U/internal/logger"
	"math"
)

type NetworkStatus struct {
	client    *ethclient.Client
	contract  *contract.RandomNumber
	auth      *bind.TransactOpts
	isHealthy bool
}

type VRFClient struct {
	networks       map[string]*NetworkStatus
	currentNetwork string
	fromAddress    common.Address
	pollInterval   time.Duration
	config         *config.Config
	requestBlock   uint64
	retryInterval  time.Duration
	maxRetries     int
}

func NewVRFClient(ctx context.Context, cfg *config.Config) (*VRFClient, error) {
	logger.Info("初始化VRF客户端")

	// 初始化私钥
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, fmt.Errorf("加载私钥失败: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("无法获取公钥")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 创建VRF客户端实例
	client := &VRFClient{
		networks:      make(map[string]*NetworkStatus),
		fromAddress:   fromAddress,
		pollInterval:  cfg.VRF.PollInterval,
		config:        cfg,
		retryInterval: cfg.VRF.Retry.Interval,
		maxRetries:    cfg.VRF.Retry.MaxRetries,
	}

	// 并行初始化网络
	errChan := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	// 初始化主网络
	go func() {
		defer wg.Done()
		if err := client.initializeNetworkWithRetry(ctx, cfg.Networks.Primary, privateKey); err != nil {
			errChan <- fmt.Errorf("初始化主网络失败: %w", err)
			return
		}
		client.networks[cfg.Networks.Primary.Name].isHealthy = true
	}()

	// 初始化备用网络
	go func() {
		defer wg.Done()
		if err := client.initializeNetworkWithRetry(ctx, cfg.Networks.Fallback, privateKey); err != nil {
			errChan <- fmt.Errorf("初始化备用网络失败: %w", err)
			return
		}
		client.networks[cfg.Networks.Fallback.Name].isHealthy = true
	}()

	// 等待初始化完成
	wg.Wait()
	close(errChan)

	// 检查错误
	var errors []string
	for err := range errChan {
		errors = append(errors, err.Error())
	}

	// 如果两个网络都初始化失败，返回错误
	if len(client.networks) == 0 {
		return nil, fmt.Errorf("所有网络初始化失败: %v", strings.Join(errors, "; "))
	}

	// 设置当前网络为主网络（如果可用）
	if status, ok := client.networks[cfg.Networks.Primary.Name]; ok && status.isHealthy {
		client.currentNetwork = cfg.Networks.Primary.Name
	} else if status, ok := client.networks[cfg.Networks.Fallback.Name]; ok && status.isHealthy {
		client.currentNetwork = cfg.Networks.Fallback.Name
	}

	// 启动网络健康检查
	go client.startHealthCheck(ctx)

	return client, nil
}

func (c *VRFClient) initializeNetworkWithRetry(ctx context.Context, network config.NetworkConfig, privateKey *ecdsa.PrivateKey) error {
	var lastErr error
	for i := 0; i < c.maxRetries; i++ {
		if i > 0 {
			logger.Infof("重试初始化网络 %s，第 %d 次", network.Name, i+1)
			time.Sleep(c.retryInterval)
		}

		client, contractInstance, auth, err := initializeNetwork(ctx, network, privateKey)
		if err != nil {
			lastErr = err
			logger.Errorf("初始化网络 %s 失败: %v", network.Name, err)
			continue
		}

		c.networks[network.Name] = &NetworkStatus{
			client:    client,
			contract:  contractInstance,
			auth:      auth,
			isHealthy: true,
		}
		return nil
	}
	return fmt.Errorf("在 %d 次尝试后仍然失败: %w", c.maxRetries, lastErr)
}

func (c *VRFClient) startHealthCheck(ctx context.Context) {
	ticker := time.NewTicker(c.config.VRF.HealthCheck.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.checkNetworkHealth(ctx)
		}
	}
}

func (c *VRFClient) checkNetworkHealth(ctx context.Context) {
	for name, status := range c.networks {
		// 创建一个带超时的上下文
		checkCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		_, err := status.client.ChainID(checkCtx)
		cancel()

		wasHealthy := status.isHealthy
		status.isHealthy = err == nil

		if wasHealthy && !status.isHealthy {
			logger.Info("网络 %s 变为不健康状态", name)
			// 如果当前网络不健康，尝试切换到其他网络
			if name == c.currentNetwork {
				c.trySwitch()
			}
		} else if !wasHealthy && status.isHealthy {
			logger.Infof("网络 %s 恢复健康状态", name)
			// 如果主网络恢复，切换回主网络
			if name == c.config.Networks.Primary.Name && c.currentNetwork != name {
				c.switchToNetwork(name)
			}
		}
	}
}

func (c *VRFClient) trySwitch() {
	// 优先尝试切换到主网络
	if status, ok := c.networks[c.config.Networks.Primary.Name]; ok && status.isHealthy {
		c.switchToNetwork(c.config.Networks.Primary.Name)
		return
	}
	// 否则尝试切换到备用网络
	if status, ok := c.networks[c.config.Networks.Fallback.Name]; ok && status.isHealthy {
		c.switchToNetwork(c.config.Networks.Fallback.Name)
	}
}

func (c *VRFClient) switchToNetwork(network string) {
	if status, ok := c.networks[network]; ok && status.isHealthy {
		c.currentNetwork = network
		logger.Infof("切换到网络: %s", network)
	}
}

func (c *VRFClient) getCurrentNetworkStatus() *NetworkStatus {
	return c.networks[c.currentNetwork]
}

func (c *VRFClient) RequestRandomNumber(ctx context.Context) (*big.Int, error) {
	status := c.getCurrentNetworkStatus()
	if status == nil {
		return nil, fmt.Errorf("没有可用的网络")
	}

	requestId, err := c.tryRequestRandomNumber(ctx, status)
	if err != nil {
		logger.Errorf("当前网络请求失败: %v, 尝试切换网络", err)
		c.trySwitch()

		// 获取新的网络状态
		status = c.getCurrentNetworkStatus()
		if status == nil {
			return nil, fmt.Errorf("没有可用的备用网络")
		}

		return c.tryRequestRandomNumber(ctx, status)
	}

	return requestId, nil
}

func (c *VRFClient) tryRequestRandomNumber(ctx context.Context, status *NetworkStatus) (*big.Int, error) {
	startTime := time.Now()
	logger.Infof("在%s网络上请求随机数", c.currentNetwork)

	// 更新nonce和gas价格
	nonce, err := status.client.PendingNonceAt(ctx, c.fromAddress)
	if err != nil {
		return nil, fmt.Errorf("获取nonce失败: %w", err)
	}

	gasPrice, err := status.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取gas价格失败: %w", err)
	}

	status.auth.Nonce = big.NewInt(int64(nonce))
	status.auth.GasPrice = gasPrice

	// 发送交易
	tx, err := status.contract.RequestRandomWords(
		status.auth,
		c.config.VRF.NumWords,
		c.config.VRF.GasLimit,
		c.config.VRF.Confirmations,
	)
	if err != nil {
		return nil, fmt.Errorf("请求随机数失败: %w", err)
	}

	logger.Infof("交易已发送: %s", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, status.client, tx)
	if err != nil {
		return nil, fmt.Errorf("等待交易确认失败: %w", err)
	}

	c.requestBlock = receipt.BlockNumber.Uint64()
	logger.Infof("请求随机数的区块号: %d", c.requestBlock)

	gasUsed := receipt.GasUsed
	gasCost := new(big.Float).Quo(
		new(big.Float).SetInt(new(big.Int).Mul(gasPrice, big.NewInt(int64(gasUsed)))),
		big.NewFloat(math.Pow10(18)),
	)
	gasCostFloat, _ := gasCost.Float64()

	logger.Infof("RequestRandomNumber接口Gas费用: %.8f MATIC", gasCostFloat)

	var requestId *big.Int
	for _, log := range receipt.Logs {
		event, err := status.contract.ParseRequestedRandomness(*log)
		if err == nil && event != nil {
			logger.Infof("获取到requestId: %s", event.RequestId.String())
			requestId = event.RequestId
			break
		}
	}

	if requestId == nil {
		return nil, fmt.Errorf("未能从交易日志中获取requestId")
	}

	logger.Infof("在%s网络上请求随机数成功，耗时: %s", c.currentNetwork, time.Since(startTime))
	return requestId, nil
}

func (c *VRFClient) WaitForRandomNumber(ctx context.Context, requestId *big.Int) ([]*big.Int, error) {
	startTime := time.Now()
	logger.Info("------------------------获取随机数阶段-----------------------")
	logger.Info("等待随机数结果")

	ticker := time.NewTicker(c.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			// 获取当前网络状态
			status := c.getCurrentNetworkStatus()
			if status == nil {
				logger.Errorf("当前没有可用的网络，尝试切换网络")
				c.trySwitch()
				continue
			}

			request, err := status.contract.GetRandomRequest(&bind.CallOpts{Context: ctx}, requestId)
			if err != nil {
				logger.Errorf("获取随机数请求失败: %v, 尝试切换网络", err)
				c.trySwitch()
				continue
			}

			if request.Fulfilled {
				logger.Infof("获取随机数阶段总耗时: %s", time.Since(startTime))
				logger.Info("---------------------------end-----------------------------")
				return request.RandomNumbers, nil
			}

			logger.Info("等待随机数生成中...")
		}
	}
}

func (c *VRFClient) Close() {
	for _, status := range c.networks {
		if status.client != nil {
			status.client.Close()
		}
	}
}

func initializeNetwork(ctx context.Context, network config.NetworkConfig, privateKey *ecdsa.PrivateKey) (*ethclient.Client, *contract.RandomNumber, *bind.TransactOpts, error) {
	// 连接网络
	client, err := ethclient.Dial(network.RPCURL)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("连接网络失败: %w", err)
	}

	// 获取chainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		client.Close()
		return nil, nil, nil, fmt.Errorf("获取chainID失败: %w", err)
	}

	// 创建auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		client.Close()
		return nil, nil, nil, fmt.Errorf("创建交易选项失败: %w", err)
	}

	// 加载合约
	address := common.HexToAddress(network.ContractAddress)
	contractInstance, err := contract.NewRandomNumber(address, client)
	if err != nil {
		client.Close()
		return nil, nil, nil, fmt.Errorf("加载合约失败: %w", err)
	}

	return client, contractInstance, auth, nil
}
