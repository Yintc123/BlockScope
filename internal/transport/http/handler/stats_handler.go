package handler

import (
	"net/http"
	"time"

	"github.com/Yintc123/BlockScope/internal/domain"
	"github.com/Yintc123/BlockScope/internal/service"
	"github.com/Yintc123/BlockScope/internal/transport/http/request"
	"github.com/Yintc123/BlockScope/internal/validator"
	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	service service.DailyActiveAddressService
}

func NewStatsHandler(service service.DailyActiveAddressService) *StatsHandler {
	return &StatsHandler{service: service}
}

func (handler *StatsHandler) GetDailyActiveAddress(ctx *gin.Context) {
	var req request.DailyActiveAddressQuery

	// 把 HTTP Request 的 query string（URL 參數）解析後再填進 req struct
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var date time.Time
	// 由於先前已經用 Validator 驗證 req，故基本上不會驗證錯誤
	date, _ = time.Parse("2006-01-02", req.Date)

	var result *domain.DailyActiveAddress
	var err error
	result, err = handler.service.GetDailyActiveAddress(ctx.Request.Context(), date, req.Chain)
	if err != nil || result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    result.ID,
		"date":  result.Date,
		"chain": result.Chain,
		"count": result.Count,
	})
}

func (handler *StatsHandler) TestAPI(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
