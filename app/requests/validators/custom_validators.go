package validators

import (
	"github.com/shanedoc/gohub/pkg/captcha"
	"github.com/shanedoc/gohub/pkg/verifycode"
)

// 验证规则
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

//密码验证:两次验证码
func ValidatePasswordConfirm(password, password_confirm string, errs map[string][]string) map[string][]string {
	if password != password_confirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

//校验:手机/邮箱
func ValidateVerifyCode(email, code string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(email, code); !ok {
		errs["verifycode"] = append(errs["verifycode"], "验证码有误")
	}
	return errs
}
