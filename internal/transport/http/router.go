package http

import (
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	route *gin.Engine,
	statsHandler *handler.StatsHandler,
	healthcheckHandler *handler.HealthcheckHandler,
) {
	route.GET("/stats/daily-active-address", statsHandler.GetDailyActiveAddress)
	route.GET("/health", healthcheckHandler.Check)
}
