package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Yintc123/BlockScope/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDailyActiveAddressService struct {
	// 匿名嵌入
	mock.Mock
}

// 實作 service interface 的方法
func (mockService *MockDailyActiveAddressService) GetDailyActiveAddress(
	ctx context.Context,
	date time.Time,
	chain string,
) (*domain.DailyActiveAddress, error) {
	args := mockService.Called(ctx, date, chain)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.DailyActiveAddress), args.Error(1)
}

func TestStatsHandler_GetDailyActiveAddress(t *testing.T) {
	// 設定 gin 為測試模式，減少日誌輸出
	gin.SetMode(gin.TestMode)

	t.Run("成功取得資料_200", func(t *testing.T) {
		// 於 t.Run 內初始化 mockService 和 mockHandler 進行測試隔離
		mockService := new(MockDailyActiveAddressService)
		mockHandler := NewStatsHandler(mockService)
		router := setupTestRouter(mockHandler)
		// 按照 routes 檔案的定義註冊路由
		// router := gin.New()
		// statsGroup := router.Group("stats")
		// statsGroup.GET("/daily-active-address", mockHandler.GetDailyActiveAddress)

		// 準備測試數據
		var targetDate time.Time = time.Date(2024, 5, 20, 0, 0, 0, 0, time.UTC)
		var mockResult *domain.DailyActiveAddress = &domain.DailyActiveAddress{
			ID:    1,
			Date:  targetDate,
			Chain: "btc",
			Count: 1000,
		}

		// 預期的 Mock 回應
		mockService.On("GetDailyActiveAddress", mock.Anything, targetDate, "btc").Return(mockResult, nil).Once()

		// 發送請求：使用實際路由路徑
		req := httptest.NewRequest("GET", "/stats/daily-active-address?date=2024-05-20&chain=btc", nil)
		respWriter := httptest.NewRecorder()
		router.ServeHTTP(respWriter, req)

		// 驗證結果
		assert.Equal(t, http.StatusOK, respWriter.Code)
		var resp map[string]interface{}
		json.Unmarshal(respWriter.Body.Bytes(), &resp)
		assert.Equal(t, float64(1000), resp["count"])
		assert.Contains(t, respWriter.Body.String(), `"count":1000`)
		mockService.AssertExpectations(t)
	})

	t.Run("日期格式錯誤_400", func(t *testing.T) {
		mockService := new(MockDailyActiveAddressService)
		mockHandler := NewStatsHandler(mockService)
		router := setupTestRouter(mockHandler)

		// router := gin.New()
		// statsGroup := router.Group("/stats")
		// statsGroup.GET("/daily-active-address", mockHandler.GetDailyActiveAddress)

		// 發送錯誤的日期格式
		req := httptest.NewRequest("GET", "/stats/daily-active-address?date=2024/05/20&chain=btc", nil)
		respWriter := httptest.NewRecorder()
		router.ServeHTTP(respWriter, req)

		// 驗證結果，回傳 BadRequest 400
		assert.Equal(t, http.StatusBadRequest, respWriter.Code)
		// 這裡會觸發 Validator 的錯誤
		assert.Contains(t, respWriter.Body.String(), "error")
	})

	t.Run("找不到資料_404", func(t *testing.T) {
		mockService := new(MockDailyActiveAddressService)
		mockHandler := NewStatsHandler(mockService)
		router := setupTestRouter(mockHandler)
		// router := gin.New()
		// statsGroup := router.Group("/stats")
		// statsGroup.GET("/daily-active-address", mockHandler.GetDailyActiveAddress)

		// 準備測試數據
		var targetDate time.Time = time.Date(2024, 5, 20, 0, 0, 0, 0, time.UTC)
		var targetChain string = "btc"
		var mockResult *domain.DailyActiveAddress = nil

		mockService.On("GetDailyActiveAddress", mock.Anything, targetDate, targetChain).Return(mockResult, nil).Once()

		req := httptest.NewRequest("GET", "/stats/daily-active-address?date=2024-05-20&chain=btc", nil)
		respWriter := httptest.NewRecorder()
		router.ServeHTTP(respWriter, req)

		assert.Equal(t, http.StatusNotFound, respWriter.Code)
		assert.Contains(t, respWriter.Body.String(), "error")
		mockService.AssertExpectations(t)
	})

	t.Run("連線資料庫失敗_500", func(t *testing.T) {
		mockService := new(MockDailyActiveAddressService)
		mockHandler := NewStatsHandler(mockService)
		router := setupTestRouter(mockHandler)
		// router := gin.New()
		// statsGroup := router.Group("/stats")
		// statsGroup.GET("/daily-active-address", mockHandler.GetDailyActiveAddress)
		var targetDate time.Time = time.Date(2024, 5, 20, 0, 0, 0, 0, time.UTC)
		var targetChain string = "btc"
		var mockResult *domain.DailyActiveAddress = nil

		mockService.On("GetDailyActiveAddress", mock.Anything, targetDate, targetChain).Return(mockResult, errors.New("db connection failed")).Once()

		req := httptest.NewRequest("GET", "/stats/daily-active-address?date=2024-05-20&chain=btc", nil)
		respWriter := httptest.NewRecorder()
		router.ServeHTTP(respWriter, req)

		assert.Equal(t, http.StatusInternalServerError, respWriter.Code)
		assert.Contains(t, respWriter.Body.String(), "db connection failed")
		mockService.AssertExpectations(t)
	})
}

// 輔助函式：快速建立測試用的 gin Engine 與路由
func setupTestRouter(handler *StatsHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	statsGroup := router.Group("/stats")
	statsGroup.GET("/daily-active-address", handler.GetDailyActiveAddress)
	return router
}
