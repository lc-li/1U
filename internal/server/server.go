package server

import (
	"fmt"
	"net/http"

	"1U/config"
	"1U/internal/handler"
	"1U/internal/logger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	// 初始化Gin引擎
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())

	// 设置路由
	handler.RegisterRoutes(router)

	return &Server{
		Config: cfg,
		Router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("127.0.0.1:%d", s.Config.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.Router,
		ReadTimeout:  s.Config.Server.ReadTimeout,
		WriteTimeout: s.Config.Server.WriteTimeout,
	}

	logger.Infof("服务器正在运行，监听地址 %s", addr)
	return srv.ListenAndServe()
}
