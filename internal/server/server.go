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

	// 添加 CORS 中间件
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 生产环境建议设置具体域名
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 设置路由
	handler.RegisterRoutes(router)

	return &Server{
		Config: cfg,
		Router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("0.0.0.0:%d", s.Config.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.Router,
		ReadTimeout:  s.Config.Server.ReadTimeout,
		WriteTimeout: s.Config.Server.WriteTimeout,
	}

	logger.Infof("服务器正在运行，监听地址 %s", addr)
	return srv.ListenAndServe()
}
