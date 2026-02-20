package domain

import "time"

type DailyActiveAddress struct {
	ID uint `gorm:"primaryKey"`
	// 強制資料庫使用 date 類型，並在寫入時自動處理時區轉換
	Date  time.Time `gorm:"type:date;index:idx_date_chain,unique"`
	Chain string    `gorm:"index:idx_date_chain,unique"`
	Count int64
}
