package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"worker-test/config"
	"worker-test/contract"
	"worker-test/internal/logger"
)

type VRFClient struct {
	client       *ethclient.Client
	contract     *contract.RandomNumber
	auth         *bind.TransactOpts
	fromAddress  common.Address
	pollInterval time.Duration
	config       *config.VRFConfig
}

func NewVRFClient(ctx context.Context, cfg *config.Config) (*VRFClient, error) {
	logger.Info("初始化VRF客户端")

	client, err := ethclient.Dial(cfg.Ethereum.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("连接以太坊失败: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("加载私钥失败: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("无法获取公钥")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取chainID失败: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("创建交易选项失败: %w", err)
	}

	address := common.HexToAddress(cfg.Ethereum.ContractAddress)
	randomNumber, err := contract.NewRandomNumber(address, client)
	if err != nil {
		return nil, fmt.Errorf("加载合约失败: %w", err)
	}

	return &VRFClient{
		client:       client,
		contract:     randomNumber,
		auth:         auth,
		fromAddress:  fromAddress,
		pollInterval: cfg.VRF.PollInterval,
		config:       &cfg.VRF,
	}, nil
}

func (c *VRFClient) RequestRandomNumber(ctx context.Context) (*big.Int, error) {
	logger.Info("开始请求随机数")

	nonce, err := c.client.PendingNonceAt(ctx, c.fromAddress)
	if err != nil {
		return nil, fmt.Errorf("获取nonce失败: %w", err)
	}

	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取gas价格失败: %w", err)
	}

	c.auth.Nonce = big.NewInt(int64(nonce))
	c.auth.GasPrice = gasPrice

	tx, err := c.contract.RequestRandomWords(
		c.auth,
		c.config.NumWords,
		c.config.GasLimit,
		c.config.Confirmations,
	)
	if err != nil {
		return nil, fmt.Errorf("请求随机数失败: %w", err)
	}

	logger.Infof("交易已发送，hash: %s", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return nil, fmt.Errorf("等待交易确认失败: %w", err)
	}

	for _, log := range receipt.Logs {
		event, err := c.contract.ParseRequestedRandomness(*log)
		if err == nil && event != nil {
			logger.Infof("获取到requestId: %s", event.RequestId.String())
			return event.RequestId, nil
		}
	}

	return nil, fmt.Errorf("未能从交易日志中获取requestId")
}

func (c *VRFClient) WaitForRandomNumber(ctx context.Context, requestId *big.Int) ([]*big.Int, error) {
	logger.Info("等待随机数结果")

	ticker := time.NewTicker(c.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			request, err := c.contract.GetRandomRequest(&bind.CallOpts{Context: ctx}, requestId)
			if err != nil {
				return nil, fmt.Errorf("获取随机请求失败: %w", err)
			}

			if request.Fulfilled {
				logger.Info("随机数生成完成")
				return request.RandomNumbers, nil
			}

			logger.Info("等待随机数生成中...")
		}
	}
}

func (c *VRFClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
