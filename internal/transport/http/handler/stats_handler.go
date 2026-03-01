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

	// 綁定 Query 參數：把 HTTP Request 的 query string（URL 參數）解析後再填進 req struct
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 將 AppError 結構的物件放入 ctx.Errors 陣列中，return 之後會由 middleware 處理
		ctx.Error(domain.NewBadRequestError("Parameter binding failed: " + err.Error()))
		return
	}

	// 驗證結構(validator)
	if err := validator.Validator.Struct(req); err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 將 AppError 結構的物件放入 ctx.Errors 陣列中，return 之後會由 middleware 處理
		ctx.Error(domain.NewBadRequestError("Request validation failed: " + err.Error()))
		return
	}

	var date time.Time
	// 由於先前已經用 Validator 驗證 req，故基本上不會驗證錯誤
	date, _ = time.Parse("2006-01-02", req.Date)

	var result *domain.DailyActiveAddress
	var err error
	result, err = handler.service.GetDailyActiveAddress(ctx.Request.Context(), date, req.Chain)

	// 優先處理錯誤
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "db connection failed"})
		// 將 AppError 結構的物件放入 ctx.Errors 陣列中，return 之後會由 middleware 處理
		ctx.Error(domain.NewInternalError(err, "failed to get daily active address"))
		return
	}

	// 再處理查無資料
	if result == nil {
		// ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		// 將 AppError 結構的物件放入 ctx.Errors 陣列中，return 之後會由 middleware 處理
		ctx.Error(domain.NewNotFoundError("data not found"))
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
