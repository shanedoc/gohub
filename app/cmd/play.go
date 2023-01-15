package cmd

import (
	"github.com/spf13/cobra"
)

//创建play命令 方便本地调试接口

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

//调试完成后清除测试代码
func runPlay(cmd *cobra.Command, args []string) {

}
