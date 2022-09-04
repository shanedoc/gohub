package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化gin
	//r := gin.Default()
	//new Gin框架实例
	r := gin.New()

	//注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	//注册路由
	r.GET("/", func(c *gin.Context) {
		//json格式响应
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	//处理404请求
	r.NoRoute(func(c *gin.Context) {
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

	//监听服务8080端口
	r.Run(":8080")
}
