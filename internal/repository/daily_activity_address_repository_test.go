package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Yintc123/BlockScope/internal/config"
	"github.com/Yintc123/BlockScope/internal/db"
	"github.com/Yintc123/BlockScope/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestDailyActiveAddressRepository_Integration(t *testing.T) {
	// 1. 初始化配置與 DB
	cfg, _ := config.LoadConfig("test")
	dbConn, err := db.NewDB(cfg.DB)
	if err != nil {
		t.Skip("跳過整合測試： 無法連線資料庫")
	}

	// 確保表結構符合最新的 domain 定義
	dbConn.AutoMigrate(&domain.DailyActiveAddress{})

	repo := NewDailyActiveAddressRepository(dbConn)
	ctx := context.Background()

	// 測試用資料：使用不同時段的 Time 來驗證 type:date 是否自動忽略時間
	testDate := time.Date(2024, 5, 20, 15, 30, 0, 0, time.Local)
	testChain := "sol"

	// 每次測試前清理該資料，確保測試獨立性
	t.Cleanup(func() {
		dbConn.Where("chain = ?", testChain).Delete(&domain.DailyActiveAddress{})
	})

	t.Run("成功建立資料並讀取", func(t *testing.T) {
		addr := &domain.DailyActiveAddress{
			Date:  testDate,
			Chain: testChain,
			Count: 500000,
		}

		err := repo.Create(ctx, addr)
		assert.NoError(t, err)
		assert.NotZero(t, addr.ID)

		// testDate 再加一小時
		var nextHour time.Time = testDate.Add(time.Hour)
		// 驗證讀取：即便查詢時給的時間不同，，只要日期符合就預期能找到結果
		found, err := repo.FindByDate(ctx, nextHour, testChain)
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, int64(500000), found.Count)
	})

	t.Run("觸發唯一索引衝突", func(t *testing.T) {
		duplicate := &domain.DailyActiveAddress{
			Date:  testDate,
			Chain: testChain,
			Count: 999,
		}

		err := repo.Create(ctx, duplicate)
		// 預期會報錯，由於已有相同的 Date + Chain
		assert.Error(t, err, "預期會觸發 Unique Index 衝突錯誤")
	})
}
