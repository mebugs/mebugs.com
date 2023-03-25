package validate

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// 字符最大长度
func maxLength(lv validator.FieldLevel) bool {
	val := lv.Field().String()
	if val == "" {
		return true
	}
	i, _ := strconv.Atoi(lv.Param())
	l := len(val)
	return l <= i
}

// checkPhoneNum 校验手机号是否合法
// @param phone 手机号
// @return legal 手机号是否合法
func checkPhoneNum(phone string) (legal bool) {
	pattern := `^[1]([3-9])[0-9]{9}$`
	legal, _ = regexp.MatchString(pattern, phone)
	return
}
