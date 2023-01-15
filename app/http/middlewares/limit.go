package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/app"
	"github.com/shanedoc/gohub/pkg/limiter"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/shanedoc/gohub/pkg/response"
	"github.com/spf13/cast"
)

/*
	limitIP为全局中间件 根据ip限制访问
	eg:
		5 reqs/second: "5-S"
		10 reqs/minute: "10-M"
*/

func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		//针对ip限流
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

//限流中间件 使用单独路由
func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		//针对路由增加访问次数
		c.Set("limiter-once", true)
		//ip+route
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	//获取超额情况
	rote, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}
	//设置头信息
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rote.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rote.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rote.Reset))
	//超额
	if rote.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "请求频繁",
		})
		return false
	}
	return true
}
