// request DTO(Data Transfer Object)
// 1. 僅需描述前端需要帶什麼資料給 API server，不需包含任何業務邏輯
// 2. 封裝前端帶來的資料，驗證和綁定
package request

type DailyActiveAddressQuery struct {
	Date  string `json:"date" form:"date" validate:"required,datetime=2006-01-02"`
	Chain string `json:"chain" form:"chain" validate:"required,oneof=eth btc sol"`
}
