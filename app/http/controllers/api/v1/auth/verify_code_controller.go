package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/app/requests"
	"github.com/shanedoc/gohub/pkg/captcha"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/shanedoc/gohub/pkg/response"
	"github.com/shanedoc/gohub/pkg/verifycode"
)

type VerifyController struct {
	v1.BaseAPIController
}

func (vc *VerifyController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	//记录错误日志
	logger.LogIf(err)
	//返回响应
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

//发送短信验证码
func (vc *VerifyController) SendUsingPhone(c *gin.Context) {
	//验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}
	//发送sms
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送失败")
	} else {
		response.Success(c)
	}
}

//发送email
func (vc *VerifyController) SendUsingEmail(c *gin.Context) {
	//验证表单
	request := requests.VerifyCodeEmailRequest{}
	logger.DebugJSON("发送邮件", "发件详情", request)
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}
	//发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送email验证码失败")
	} else {
		response.Success(c)
	}
}
