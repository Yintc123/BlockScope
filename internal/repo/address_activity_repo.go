package repo

import (
	"time"

	"github.com/Yintc123/BlockScope/internal/db"
	"github.com/Yintc123/BlockScope/internal/model"
)

func CountDailyActiveAddresses(date time.Time) (int64, error) {
	var (
		count int64
		err   error
	)

	err = db.DB.Model(&model.AddressActivity{}).
		Where("activity_date = ?", date).
		Distinct("address").
		Count(&count).
		Error

	return count, err
}
