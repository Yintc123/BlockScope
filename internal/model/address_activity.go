package model

import "time"

type AddressActivity struct {
	ID           uint `gorm:"primaryKey"`
	Address      string
	TxHash       string
	ActivityDate time.Time
	createAt     time.Time
}
