package handler

import (
	"net/http"

	"github.com/Yintc123/BlockScope/internal/service"
	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
	service service.HealthcheckService
}

func NewHealthcheckHandler(service service.HealthcheckService) *HealthcheckHandler {
	return &HealthcheckHandler{service: service}
}

// 回傳 API server 和 DB 的狀態
func (handler *HealthcheckHandler) Check(ctx *gin.Context) {
	// ctxBackground := context.Background()
	ctxRequest := ctx.Request.Context()

	// dbErr := handler.service.CheckDB(ctxBackground)
	dbErr := handler.service.CheckDB(ctxRequest)
	serverAlive := handler.service.CheckServer()

	var status string = "ok"
	if !serverAlive || dbErr != nil {
		status = "fail"
	}

	var dbStatus string = "ok"
	if dbErr != nil {
		dbStatus = "fail: " + dbErr.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
		"checks": gin.H{
			"API server": serverAlive,
			"DB":         dbStatus,
		},
	})
}
