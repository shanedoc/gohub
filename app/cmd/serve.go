package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/bootstrap"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/console"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/spf13/cobra"
)

//命令行运行web应用

var CmdServer = &cobra.Command{
	Use:   "serve",
	Short: "start web serve",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	//实例化
	router := gin.New()
	//初始化路由绑定
	bootstrap.SetupRoute(router)
	//运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server,error:" + err.Error())
	}

}
