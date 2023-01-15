package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

//console命令行辅助方法

//绿色输出成功消息
func Success(msg string) {
	colorOut(msg, "green")
}

// Error 打印一条报错消息，红色输出
func Error(msg string) {
	colorOut(msg, "red")
}

// Warning 打印一条提示消息，黄色输出
func Warning(msg string) {
	colorOut(msg, "yellow")
}

// Exit 打印一条报错消息，并退出 os.Exit(1)
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

//exitif 自动判断err是否为空
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func colorOut(msg, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
