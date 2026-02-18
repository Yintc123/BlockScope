package http

import (
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	statsHandler *handler.StatsHandler,
	healthcheckHandler *handler.HealthcheckHandler,
) {
	// /stats Group
	stats := router.Group("/stats")
	{
		stats.GET("/daily-active-address", statsHandler.GetDailyActiveAddress)
	}
	router.GET("/health", healthcheckHandler.Check)
}
