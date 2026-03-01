package middleware

import (
	"log"
	"net/http"

	"github.com/Yintc123/BlockScope/internal/domain"
	"github.com/gin-gonic/gin"
)

// ErrorHandler 統一處理 Gin Handler 拋出的錯誤，每一次的請求都會到 middleware 檢查是否有錯誤訊息
func ErrorHandler() gin.HandlerFunc {
	// 這裡是攔截 gin 內部的 error 處理機制嗎？讓自定義的 error 來接手嗎？
	return func(ctx *gin.Context) {
		// 初始化時，先執行後續的 handler
		ctx.Next()
		// 之後每一次處理 Gin Handler 拋出的錯誤都從此開始接續執行

		// 檢查是否有錯誤發生，如果 handler 有呼叫 ctx.Error(err)，將錯誤放到 ctx 的錯誤陣列中
		if len(ctx.Errors) > 0 {
			// 取得最後一個錯誤(通常是最具體的一個)
			err := ctx.Errors.Last().Err

			// 判斷是否為 domain 定義的 AppError
			if appErr, ok := err.(*domain.AppError); ok {
				// 如果是 500 錯誤，記錄真實的 RawErr 到伺服器 Log
				if appErr.Code == http.StatusInternalServerError {
					log.Printf("Internal Error: %v", appErr.RawErr)
				}

				// 回應給前端（appErr 結構體會自動轉為 JSON）
				ctx.JSON(appErr.Code, appErr)
				return
			}

			// 處理未知錯誤（這通常是嚴重的程式 Bug，必須記錄 Log）
			log.Printf("internal Error: %v", err)
			// 處理未預期的錯誤(例如未被包裝過的系統錯誤)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unknown System Error."})
		}
	}
}
