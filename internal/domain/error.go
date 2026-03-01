package domain

import "fmt"

// AppError 定義定義業務邏輯錯誤
type AppError struct {
	Code    int    `json:"-"`       // 用於 http 狀態碼，"-" 代表不序列化這個欄位
	Message string `json:"message"` // 回應前端的錯誤訊息
	RawErr  error  `json:"-"`       // 原始錯誤(僅用於 server 的 log)，"-" 代表不序列化這個欄位
}

// 註：json:"-"，ctx.JSON() 會忽略此欄位

// AppError 實作了 Error() string 方法，因此 AppError 型別完全符合 error 介面
func (e *AppError) Error() string {
	if e.RawErr != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.RawErr)
	}
	return e.Message
}

// *** 預定義錯誤工廠(工廠模式)，每一次都會建立新的 AppError 實例 ***
// 發送的數據格式錯誤、參數缺失或參數類型不正確
func NewBadRequestError(msg string) *AppError {
	return &AppError{
		Code:    400,
		Message: msg,
	}
}

// 用戶未認證或 Token 過期
func NewUnauthorizedError(msg string) *AppError {
	return &AppError{
		Code:    401,
		Message: msg,
	}
}

// 已認證，但無權限操作資源
func NewForbiddenError(msg string) *AppError {
	return &AppError{
		Code:    403,
		Message: msg,
	}
}

// 資源不存在，資料在資料庫中不存在
func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Code:    404,
		Message: msg,
	}
}

// 伺服器錯誤
func NewInternalError(rawErr error, msg string) *AppError {
	var defaultMessage string = "Internal Server Error"
	if msg != "" {
		defaultMessage = msg
	}
	return &AppError{
		Code:    500,
		Message: defaultMessage,
		RawErr:  rawErr,
	}
}
