package handler

import (
	"context"
	"net/http"

	"1U/config"
	"1U/internal/logger"
	"1U/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Config     *config.Config
	VRFService service.VRFService
}

func NewHandler(cfg *config.Config, vrfService service.VRFService) *Handler {
	return &Handler{
		Config:     cfg,
		VRFService: vrfService,
	}
}

func RegisterRoutes(router *gin.Engine) {
	// 初始化服务
	cfg := config.GetConfig()
	vrfService := service.GetVRFService()

	h := NewHandler(cfg, vrfService)

	// 健康检查
	router.GET("/health", h.HealthCheck)

	// 获取随机数
	router.GET("/random", h.GetRandomNumbers)

	// 其他路由可以继续添加
}

// 健康检查处理函数
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// 获取随机数处理函数
func (h *Handler) GetRandomNumbers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Config.VRF.Timeout)
	defer cancel()

	randomNumbers, err := h.VRFService.GetRandomNumbers(ctx)
	if err != nil {
		logger.Errorf("获取随机数失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"random_numbers": randomNumbers,
	})
}
