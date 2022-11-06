package middlewares

import (
	"net"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/shanedoc/gohub/pkg/response"
	"go.uber.org/zap"
)

//记录gin运行过程出现的panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//获取用户信息
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				//连接中断,客户端中断为正常情况不需要记录错误信息
				var brokenpiple bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*net.OpError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenpiple = true
						}
					}
				}
				//链接中断
				if brokenpiple {
					logger.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					//断开链接,无法写状态
					return
				}

				//如果不是链接中断,就开始记录堆栈信息
				logger.Error("recovery from panic",
					zap.Time("time", time.Now()),               // 记录时间
					zap.Any("error", err),                      // 记录错误信息
					zap.String("request", string(httpRequest)), // 请求信息
					zap.Stack("stacktrace"),
				)

				//返回500状态码
				response.Abort500(c)
			}
		}()
		c.Next()
	}
}
