package main

import (
	"fmt"
	"os"

	"github.com/shanedoc/gohub/app/cmd"
	"github.com/shanedoc/gohub/bootstrap"
	btsConfig "github.com/shanedoc/gohub/config"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/console"
	"github.com/spf13/cobra"
)

func init() {
	//加载config目录下配置文件
	btsConfig.Initialize()
}

func main() {
	//web入口调用serve命令
	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple web project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		//rootCmd所有命令都会执行以下代码

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
	}
	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServer,
	)
	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServer)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

/*func main() {
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

	//初始化redis
	bootstrap.SetupRedis()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	//verfiycode.NewVerifyCode().SendSMS("15652946160")

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

}*/
