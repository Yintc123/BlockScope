package repository

import "gorm.io/gorm"

// 分層封裝，Repositories 為複數形式，代表所有 Repo 的集合
type Repositories struct {
	DailyActiveAddress DailyActiveAddressRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		DailyActiveAddress: NewDailyActiveAddressRepository(db),
	}
}
