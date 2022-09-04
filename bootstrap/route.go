package bootstrap

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/routes"
)

//路由初始化
func SetupRoute(route *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleware(route)
	//注册路由api
	routes.RegisterAPIRoutes(route)
	//配置404
	setup404Handler(route)
}

func registerGlobalMiddleware(route *gin.Engine) {
	route.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(route *gin.Engine) {
	//处理404请求
	route.NoRoute(func(c *gin.Context) {
		//获取http accept头信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//html
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			//default json 返回
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
