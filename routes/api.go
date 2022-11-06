//注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/app/http/controllers/api/v1/auth"
)

//注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//v1路由组,把路由全部放到该路由组下
	v1 := r.Group("v1")
	{
		//auth路由
		authGroup := v1.Group("auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist) //校验手机号
			authGroup.POST("/signup/email/exist", suc.IsEmailExist) //校验邮件

			//发送验证码
			vcc := new(auth.VerifyController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)  //图片验证码,加限流
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone) //验证码校验
		}
	}

}
