package model

import "time"

type AddressActivity struct {
	ID           uint   `gorm:"primaryKey"`
	Address      string `gorm:"index"`
	TxHash       string
	ActivityDate time.Time `gorm:"index"`
	createAt     time.Time
}
