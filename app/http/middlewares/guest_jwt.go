package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/jwt"
	"github.com/shanedoc/gohub/pkg/response"
)

//guset使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			//token获取成功
			_, err := jwt.NewJWT().ParserToken(c)
			if err != nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
