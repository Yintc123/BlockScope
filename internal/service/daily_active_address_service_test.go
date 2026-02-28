package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Yintc123/BlockScope/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 建立 MockRepository
type MockDailyActiveAddressRepository struct {
	// 匿名嵌入
	mock.Mock
}

// 實作 Repository 介面方法
func (mockRepo *MockDailyActiveAddressRepository) FindByDate(
	ctx context.Context,
	date time.Time,
	chain string,
) (*domain.DailyActiveAddress, error) {
	args := mockRepo.Called(ctx, date, chain)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.DailyActiveAddress), args.Error(1)
}

func (mockRepo *MockDailyActiveAddressRepository) Create(
	ctx context.Context,
	address *domain.DailyActiveAddress,
) error {
	args := mockRepo.Called(ctx, address)
	return args.Error(0)
}

// 測試案例，Table-Driven Tests (表格驅動測試)
func TestDailyActiveAddressService_GetDailyActiveAddress(t *testing.T) {
	testDate := time.Now()
	testChain := "eth"

	tests := []struct {
		name          string
		date          time.Time
		chain         string
		mockReturn    *domain.DailyActiveAddress
		mockErr       error
		expectedCount int64
		expectedErr   bool
	}{
		{
			name:  "成功取得資料",
			date:  testDate,
			chain: testChain,
			mockReturn: &domain.DailyActiveAddress{
				Date:  testDate,
				Chain: testChain,
				Count: 1000,
			},
			mockErr:       nil,
			expectedCount: 1000,
			expectedErr:   false,
		},
		{
			name:          "查無資料時回傳 nil",
			date:          testDate,
			chain:         testChain,
			mockReturn:    nil,
			mockErr:       nil,
			expectedCount: 0,
			expectedErr:   false,
		},
		{
			name:          "資料庫發生錯誤",
			date:          testDate,
			chain:         testChain,
			mockReturn:    nil,
			mockErr:       errors.New("db connection failed"),
			expectedCount: 0,
			expectedErr:   true,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			ctx := context.Background()
			// 在每一次回圈都建立一個新的 mock repo，避免測試資料殘留影響下一次測試
			// 等同 mockRepo := &MockDailyActiveAddressRepository{}
			mockRepo := new(MockDailyActiveAddressRepository)
			svc := NewDailyActiveAddressService(mockRepo)
			// 設定 Mock 的期望行為
			mockRepo.On("FindByDate", ctx, testcase.date, testcase.chain).Return(testcase.mockReturn, testcase.mockErr).Once()

			// 執行測試，測試 service 的 GetDailyActiveAddress 的業務羅技是否正確
			res, err := svc.GetDailyActiveAddress(ctx, testcase.date, testcase.chain)

			// 驗證結果
			if testcase.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if testcase.mockReturn != nil {
					assert.Equal(t, testcase.expectedCount, res.Count)
				} else {
					assert.Nil(t, res)
				}
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
