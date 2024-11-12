package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"1U/config"
	"1U/contract"
	"1U/internal/logger"
	"1U/internal/price"
	"math"
)

type VRFClient struct {
	client       *ethclient.Client
	contract     *contract.RandomNumber
	auth         *bind.TransactOpts
	fromAddress  common.Address
	pollInterval time.Duration
	config       *config.VRFConfig
	priceService *price.PriceService
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
	startTime := time.Now()
	logger.Info("------------------------请求随机数阶段-----------------------")
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

	logger.Infof("交易已发送: %s", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return nil, fmt.Errorf("等待交易确认失败: %w", err)
	}

	gasUsed := receipt.GasUsed
	gasCost := new(big.Float).Quo(
		new(big.Float).SetInt(new(big.Int).Mul(gasPrice, big.NewInt(int64(gasUsed)))),
		big.NewFloat(math.Pow10(18)),
	)
	gasCostFloat, _ := gasCost.Float64()

	logger.Infof("RequestRandomNumber接口Gas费用: %.8f MATIC", gasCostFloat)

	var requestId *big.Int
	for _, log := range receipt.Logs {
		event, err := c.contract.ParseRequestedRandomness(*log)
		if err == nil && event != nil {
			logger.Infof("获取到requestId: %s", event.RequestId.String())
			requestId = event.RequestId
			break
		}
	}

	if requestId == nil {
		return nil, fmt.Errorf("未能从交易日志中获取requestId")
	}

	logger.Infof("请求随机数阶段耗时: %s", time.Since(startTime))
	logger.Info("---------------------------end-----------------------------")
	return requestId, nil
}

func (c *VRFClient) WaitForRandomNumber(ctx context.Context, requestId *big.Int) ([]*big.Int, error) {
	startTime := time.Now()
	logger.Info("------------------------获取随机数阶段-----------------------")
	logger.Info("等待随机数结果")

	ticker := time.NewTicker(c.pollInterval)
	defer ticker.Stop()

	var fulfillmentTx *types.Transaction
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
				blockNumber, err := c.client.BlockNumber(ctx)
				if err != nil {
					return nil, fmt.Errorf("获取当前区块号失败: %w", err)
				}

				filterOpts := &bind.FilterOpts{
					Start: blockNumber - 100,
					End:   &blockNumber,
				}

				iter, err := c.contract.FilterRandomnessFulfilled(filterOpts, []*big.Int{requestId}, nil)
				if err != nil {
					return nil, fmt.Errorf("过滤随机数完成事件失败: %w", err)
				}
				defer iter.Close()

				if iter.Next() {
					fulfillmentTx, _, err = c.client.TransactionByHash(ctx, iter.Event.Raw.TxHash)
					if err != nil {
						logger.Info("获取回调交易详情失败: %v", err)
					} else {
						receipt, err := c.client.TransactionReceipt(ctx, fulfillmentTx.Hash())
						if err != nil {
							logger.Info("获取回调交易收据失败: %v", err)
						} else {
							gasPrice := fulfillmentTx.GasPrice()
							gasUsed := receipt.GasUsed

							callbackGasCost := new(big.Float).Quo(
								new(big.Float).SetInt(new(big.Int).Mul(gasPrice, big.NewInt(int64(gasUsed)))),
								big.NewFloat(math.Pow10(18)),
							)
							callbackGasCostFloat, _ := callbackGasCost.Float64()

							logger.Infof("WaitForRandomNumber接口Gas费用: %.8f MATIC", callbackGasCostFloat)
							logger.Infof("获取随机数阶段总耗时: %s", time.Since(startTime))
							logger.Info("---------------------------end-----------------------------")
							return request.RandomNumbers, nil
						}
					}
				}
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
