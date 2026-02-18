package http

import (
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterStatsRoutes(
	router *gin.RouterGroup,
	statsHandler *handler.StatsHandler,
) {
	router.GET("/daily-active-address", statsHandler.GetDailyActiveAddress)
}
