package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/response"
)

//根据user_agent数据可对不同版本数据做特殊处理

//强制请求必须带user-agent头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取user-agent信息
		if len(c.Request.UserAgent()) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}
		c.Next()
	}
}
