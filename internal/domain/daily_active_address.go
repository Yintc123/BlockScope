package domain

import "time"

type DailyActiveAddress struct {
	ID    uint      `gorm:"primaryKey"`
	Date  time.Time `gorm:"index:idx_date_chain,unique"`
	Chain string    `gorm:"index:idx_date_chain,unique"`
	Count int64
}
