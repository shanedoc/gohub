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
		cmd.CmdKey,
		cmd.CmdPlay,
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
