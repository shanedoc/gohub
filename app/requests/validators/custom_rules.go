package validators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

//初始化时执行,自定义表单验证规则
func init() {
	//自定义汉字长度校验规则
	//中文长度<8
	govalidator.AddCustomRule("max_cn", func(field, rule, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})
	//中文长度>2
	govalidator.AddCustomRule("min_cn", func(field, rule, message string, value interface{}) error {
		vallength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if vallength < l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})
}
