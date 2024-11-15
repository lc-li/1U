package main

import (
	"flag"

	"1U/config"
	"1U/internal/client"
	"1U/internal/logger"
	"1U/internal/server"
	"1U/internal/service"
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

	// 创建VRF客户端
	vrfClient, err := client.NewVRFClient(cfg)
	if err != nil {
		logger.Fatalf("创建VRF客户端失败: %v", err)
	}
	defer vrfClient.Close()

	// 初始化VRF服务
	if err := service.InitVRFService(cfg, vrfClient); err != nil {
		logger.Fatalf("初始化VRF服务失败: %v", err)
	}

	// 创建并运行服务器
	srv := server.NewServer(cfg)
	if err := srv.Run(); err != nil {
		logger.Fatalf("启动服务器失败: %v", err)
	}
}
