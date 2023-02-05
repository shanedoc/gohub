//注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/app/http/controllers/api/v1/auth"
	"github.com/shanedoc/gohub/app/http/middlewares"
)

//注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//v1路由组,把路由全部放到该路由组下
	v1 := r.Group("v1")

	//全局限流中间件
	v1.Use(middlewares.LimitIP("200-H"))
	{
		//auth路由
		authGroup := v1.Group("auth")

		//测试时可提高参数设置
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			//使用手机号登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using_phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using_password", middlewares.GuestJWT(), lgc.LoginByPassword) //密码登录
			authGroup.POST("/login/refresh_token", middlewares.AuthJWT(), lgc.RefreshToken)      //令牌刷线

			//重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone) //使用手机号重置密码
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail) //使用邮件 重置密码

			//注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)     //校验手机号
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)     //校验邮件
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail) //使用邮件注册用户

			//发送验证码
			vcc := new(auth.VerifyController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)  //图片验证码,加限流
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone) //验证码校验
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail) //验证码校验

			//获取当前用户
			uc := new(controllers.UsersController)
			v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
			userGroup := v1.Group("/users")
			{
				userGroup.GET("", uc.Index)
				userGroup.PUT("", middlewares.AuthJWT(), uc.UserUpdateProfile)
			}

			//分类
			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("", cgc.Index)                                //分类列表
				cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)        //创建分类
				cgcGroup.POST("/:id", middlewares.AuthJWT(), cgc.Update)   //修改分类
				cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete) //删除分类
			}

			//话题
			tpc := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics")
			{
				tpcGroup.GET("", tpc.Index)                                //列表
				tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)        //创建帖子
				tpcGroup.POST("/:id", middlewares.AuthJWT(), tpc.Update)   //修改帖子
				tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete) //删除帖子
				tpcGroup.GET("/:id", middlewares.AuthJWT(), tpc.Show)      //查看帖子详情
			}

			//连接
			lnk := new(controllers.LinksController)
			lnkGroup := v1.Group("/links")
			{
				lnkGroup.GET("", lnk.Index)
			}

		}
	}

}
