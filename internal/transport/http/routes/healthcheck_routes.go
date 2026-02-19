package http

import (
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRootRoutes(
	router *gin.Engine,
	healthcheckHandler *handler.HealthcheckHandler,
) {
	router.GET("/health", healthcheckHandler.Check)
}
