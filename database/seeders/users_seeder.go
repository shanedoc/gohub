package seeders

import (
	"fmt"

	"github.com/shanedoc/gohub/database/factories"
	"github.com/shanedoc/gohub/pkg/console"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/shanedoc/gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	//添加seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		//创建10个用户
		users := factories.MakeUser(10)
		//批量创建用户 （注意批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)
		//记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}
		//打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
