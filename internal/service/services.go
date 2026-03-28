package service

import (
	"github.com/Yintc123/BlockScope/internal/repository"
	"gorm.io/gorm"
)

// 分層封裝，Services 為複數形式，代表所有 service 的集合
type Services struct {
	DailyActiveAddress DailyActiveAddressService
	Healthcheck        HealthcheckService
}

func NewServices(repos *repository.Repositories, db *gorm.DB) *Services {
	return &Services{
		DailyActiveAddress: NewDailyActiveAddressService(repos.DailyActiveAddress),
		Healthcheck:        NewHealthCheckService(db),
	}
}
