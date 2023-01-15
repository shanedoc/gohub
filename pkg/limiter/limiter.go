package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/shanedoc/gohub/pkg/redis"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

//限流处理逻辑

//获取limitor key ip
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

//获取路由+ip 对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

//将url中的 ’/‘ 转成 ‘-’
func routeToKeyString(url string) string {
	url = strings.ReplaceAll(url, "/", "-")
	url = strings.ReplaceAll(url, ":", "_")
	return url
}

//检查是否请求超额
func CheckRate(c *gin.Context, key, formatted string) (limiterlib.Context, error) {
	//实例化
	var content limiterlib.Context
	//定义请求的限制
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return content, err
	}
	//初始化存储 使用redis对象
	store, err := sredis.NewStoreWithOptions(redis.Reids.Client, limiterlib.StoreOptions{
		//为limiter设置前缀
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return content, err
	}

	//使用上面初始化的limiter.Rate和redis存储对象
	limitterObj := limiterlib.New(store, rate)

	//获取限流的数据
	if c.GetBool("limiter-once") {
		//peek取结果不增加访问次数
		return limitterObj.Peek(c, key)
	} else {
		//确保多个路由组调用limitip限流时 只增加一次访问次数
		c.Set("limiter-once", true)
		//取出结果并增加访问次数
		return limitterObj.Get(c, key)
	}

}
