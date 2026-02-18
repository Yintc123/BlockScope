package http

import (
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterHealthcheckRoutes(
	router *gin.Engine,
	healthcheckHandler *handler.HealthcheckHandler,
) {
	router.GET("/health", healthcheckHandler.Check)
}
