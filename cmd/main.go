package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"1U/config"
	"1U/internal/client"
	"1U/internal/logger"
)

func main() {
	// 解析命令行参数
	configPath := flag.String("config", "config/config.yaml", "配置文件路径")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	if err := logger.InitLogger(cfg.Log.Level, cfg.Log.Format, cfg.Log.Output); err != nil {
		panic(err)
	}

	// 创建上下文（支持优雅退出）
	ctx, cancel := context.WithTimeout(context.Background(), cfg.VRF.Timeout)
	defer cancel()

	// 处理信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Info("收到退出信号")
		cancel()
	}()

	// 创建VRF客户端
	vrfClient, err := client.NewVRFClient(ctx, cfg)
	if err != nil {
		logger.Fatalf("创建VRF客户端失败: %v", err)
	}
	defer vrfClient.Close()

	// 请求随机数
	requestId, err := vrfClient.RequestRandomNumber(ctx)
	if err != nil {
		logger.Fatalf("请求随机数失败: %v", err)
	}

	// 等待随机数结果
	randomNumbers, err := vrfClient.WaitForRandomNumber(ctx, requestId)
	if err != nil {
		logger.Fatalf("等待随机数失败: %v", err)
	}

	// 打印结果
	logger.Info("获取到随机数结果:")
	for i, num := range randomNumbers {
		logger.Infof("随机数 %d: %s", i+1, num.String())
	}

	for {
	}
}
