package handler

import "github.com/Yintc123/BlockScope/internal/service"

// 分層封裝，Handlers 為複數形式，代表所有 handler 的集合
type Handlers struct {
	Healthcheck *HealthcheckHandler
	Stats       *StatsHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Healthcheck: NewHealthcheckHandler(services.Healthcheck),
		Stats:       NewStatsHandler(services.DailyActiveAddress),
	}
}
