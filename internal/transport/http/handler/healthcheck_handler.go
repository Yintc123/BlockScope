package handler

import (
	"context"
	"net/http"

	"github.com/Yintc123/BlockScope/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
	service *service.HealthcheckService
}

func NewHealthcheckHandler(service *service.HealthcheckService) *HealthcheckHandler {
	return &HealthcheckHandler{service: service}
}

// 回傳 API server 和 DB 的狀態
func (handler *HealthcheckHandler) Check(c *gin.Context) {
	ctx := context.Background()

	dbErr := handler.service.CheckDB(ctx)
	serverAlive := handler.service.CheckServer()

	var status string = "ok"
	if !serverAlive || dbErr != nil {
		status = "fail"
	}

	var dbStatus string = "ok"
	if dbErr != nil {
		dbStatus = "fail: " + dbErr.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"checks": gin.H{
			"API server": serverAlive,
			"DB":         dbStatus,
		},
	})
}
