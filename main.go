package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/bootstrap"
	btsConfig "github.com/shanedoc/gohub/config"
	"github.com/shanedoc/gohub/pkg/config"
)

func init() {
	//加载config目录下配置文件
	btsConfig.Initialize()
}

func main() {
	//初始化gin
	var env string

	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	//初始化logger
	bootstrap.SetupLogger()

	//设置gin运行模式,本地开发建议debug模式
	gin.SetMode(gin.ReleaseMode)

	// new 一个 Gin Engine 实例
	router := gin.New()

	//初始化db
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

}
