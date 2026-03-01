package request

import (
	"testing"

	"github.com/Yintc123/BlockScope/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestDailyActiveAddressQuery_Validation(t *testing.T) {
	// 使用專案中定義的 validator 實例
	validate := validator.Validator

	testcases := []struct {
		name    string
		input   DailyActiveAddressQuery
		wantErr bool
	}{
		{
			name: "ValidInput",
			input: DailyActiveAddressQuery{
				Date:  "2024-05-25",
				Chain: "btc",
			},
			wantErr: false,
		},
		{
			name: "InvalidDateFormat",
			input: DailyActiveAddressQuery{
				Date:  "2024/05/25",
				Chain: "btc",
			},
			wantErr: true,
		},
		{
			name: "UnsupportedChainName",
			input: DailyActiveAddressQuery{
				Date:  "2024-05-20",
				Chain: "sui",
			},
			wantErr: true,
		},
		{
			name: "MissingDateField",
			input: DailyActiveAddressQuery{
				Chain: "btc",
			},
			wantErr: true,
		},
		{
			name: "MissingChainField",
			input: DailyActiveAddressQuery{
				Date: "2024-05-20",
			},
			wantErr: true,
		},
		{
			name:    "EmptyInput",
			input:   DailyActiveAddressQuery{},
			wantErr: true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
