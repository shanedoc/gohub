package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/bootstrap"
)

func main() {
	//初始化gin
	//r := gin.Default()
	//new Gin框架实例
	router := gin.New()

	//初始化路由并绑定
	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run()
	if err != nil {
		//处理错误信息
		fmt.Println(err.Error())
	}

}
