package make

import (
	"fmt"

	"github.com/shanedoc/gohub/pkg/app"
	"github.com/shanedoc/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeMigration(cmd *cobra.Command, args []string) {
	//日期格式化
	timeStr := app.TimenowInTimezone().Format(args[0])
	model := makeModelFromString(args[0])
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, fileName, model, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}