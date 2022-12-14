package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/pkg/captcha"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/shanedoc/gohub/pkg/response"
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
