package service

import (
	"context"
	"math/big"

	"1U/config"
	"1U/internal/client"
	"1U/internal/logger"
)

type VRFService interface {
	GetRandomNumbers(ctx context.Context) ([]*big.Int, error)
}

type vrfService struct {
	client *client.VRFClient
	config *config.Config
}

var VrfServiceInstance VRFService

func InitVRFService(cfg *config.Config, vrfClient *client.VRFClient) error {
	VrfServiceInstance = &vrfService{
		client: vrfClient,
		config: cfg,
	}
	return nil
}

func GetVRFService() VRFService {
	return VrfServiceInstance
}

func (s *vrfService) GetRandomNumbers(ctx context.Context) ([]*big.Int, error) {
	// 请求随机数
	requestId, err := s.client.RequestRandomNumber(ctx)
	if err != nil {
		return nil, err
	}

	// 等待随机数结果
	randomNumbers, err := s.client.WaitForRandomNumber(ctx, requestId)
	if err != nil {
		return nil, err
	}

	// 处理随机数结果
	s.processRandomNumbers(randomNumbers)

	return randomNumbers, nil
}

// 处理随机数的函数
func (s *vrfService) processRandomNumbers(randomNumbers []*big.Int) {
	// 您可以在这里添加对随机数的业务逻辑处理
	logger.Info("获取到随机数结果:")
	for i, num := range randomNumbers {
		logger.Infof("随机数 %d: %s", i+1, num.String())
	}
}
