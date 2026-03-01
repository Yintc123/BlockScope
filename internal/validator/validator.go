package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

// 啟動專案的時候，會直接執行 package 的 init()
func init() {
	fmt.Println("The package, validate, init starts...")
	Validator = validator.New()

	// 註冊自定義驗證 tag: "lowercase"
	_ = Validator.RegisterValidation("lowercase", validateAndLowercase)
}

// 驗證欄位是否為字串並將其轉為小寫
func validateAndLowercase(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	// 確保欄位為字串型別，使用 reflect.String 來檢查欄位的型別
	if field.Kind() != reflect.String {
		return false
	}

	value := field.String()
	// 將值轉為小寫
	lowercased := strings.ToLower(value)

	// 將小寫的值寫回結構體欄位中
	field.SetString(lowercased)

	// 驗證通過並且轉為小寫
	return true
}
