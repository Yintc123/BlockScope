package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Yintc123/BlockScope/internal/domain"
	"gorm.io/gorm"
)

type DailyActiveAddressRepository interface {
	FindByDate(
		ctx context.Context,
		date time.Time,
		chain string,
	) (*domain.DailyActiveAddress, error)
}

type dailyActiveAddressRepo struct {
	db *gorm.DB
}

func NewDailyActiveAddressRepository(db *gorm.DB) DailyActiveAddressRepository {
	return &dailyActiveAddressRepo{db: db}
}

func (repo *dailyActiveAddressRepo) FindByDate(
	ctx context.Context,
	date time.Time,
	chain string,
) (*domain.DailyActiveAddress, error) {
	// 建立回傳物件的格式
	var result domain.DailyActiveAddress

	var err error = repo.db.WithContext(ctx).
		Model(&domain.DailyActiveAddress{}).
		Where("date= ? AND chain= ?", date, chain).
		First(&result).Error

	// 如果查不到資料，Gorm 會回應 ErrRecordNotFound·但查不到資料並非 DB 出錯
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	// 回應時建立一個新的 DailyActiveAddress 物件
	return &result, nil
}
