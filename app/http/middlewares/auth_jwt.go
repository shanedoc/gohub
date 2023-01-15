package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/app/models/user"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/jwt"
	"github.com/shanedoc/gohub/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取header中的token信息,解析验证
		claims, err := jwt.NewJWT().ParserToken(c)

		//jwt解析失败
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		//jwt解析成功 设置用户信息
		userModel := user.Get(claims.Id)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		//将用户信息存入gin.Context中,后续接口的数据会校验这部分信息
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)
		c.Next()
	}
}
