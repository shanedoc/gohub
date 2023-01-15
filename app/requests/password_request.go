package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/app/requests/validators"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verfiy_code,omitempty" valid:"verfiy_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

//重置密码参数校验
func ResetByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verfiy_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号码为必填项,参数名为phone",
			"digits:手机号码长度为11位数字",
		},
		"verify_code": []string{
			"required:验证码为必填项,参数名为verify_code",
			"digits:验证码长度为6位数字",
		},
		"password": []string{
			"required:密码为必填项,参数名为password",
			"digits:密码长度最少6位长度",
		},
	}
	err := validate(data, rules, message)

	//检查验证码
	_data := data.(*ResetByPhoneRequest)
	err = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, err)
	return err
}
