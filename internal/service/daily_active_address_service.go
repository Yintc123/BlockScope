package service

import (
	"context"
	"time"

	"github.com/Yintc123/BlockScope/internal/domain"
	"github.com/Yintc123/BlockScope/internal/repository"
)

type DailyActiveAddressService interface {
	GetDailyActiveAddress(
		ctx context.Context,
		date time.Time,
		chain string,
	) (*domain.DailyActiveAddress, error)
}

type dailyActiveAddressService struct {
	// interface 本身就是 reference type，所以不用加 *
	repo repository.DailyActiveAddressRepository
}

func NewDailyActiveAddressService(
	repo repository.DailyActiveAddressRepository,
) DailyActiveAddressService {
	return &dailyActiveAddressService{repo: repo}
}

func (service *dailyActiveAddressService) GetDailyActiveAddress(
	ctx context.Context,
	date time.Time,
	chain string,
) (*domain.DailyActiveAddress, error) {
	return service.repo.FindByDate(ctx, date, chain)
}
