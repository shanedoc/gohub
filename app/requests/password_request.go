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

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

//手机号重置密码参数校验
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

//邮件重置密码参数校验
func ResetByEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	message := govalidator.MapData{
		"email": []string{
			"required:email必填",
			"min:emial长度需大于4",
			"max:email长度需小雨30",
			"email:email格式不正确,请填写有效的邮箱地址",
		},
		"verify_code": []string{
			"required:验证码为必填项",
			"digits:验证码长度必须为6位数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
	}

	err := validate(data, rules, message)
	_data := data.(*ResetByEmailRequest)
	err = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, err)
	return err
}
