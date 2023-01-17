package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeModel(cmd *cobra.Command, args []string) {
	//格式化模型名称,返回一个model对象
	model := makeModelFromString(args[0])

	//确保模型目录存在 eg:app/models/user
	dir := fmt.Sprintf("app/models/%s/", model.PackageName)
	//确保目录的目标存在,确保父目录和子目录都会创建  第二个参数是创建目录权限:0777
	os.MkdirAll(dir, os.ModePerm)

	//替换变量
	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"_util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)
}
