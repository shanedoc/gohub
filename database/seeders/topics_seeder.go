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

	seed.Add("SeedTopicsTable", func(db *gorm.DB) {

		topics := factories.MakeTopics(10)

		result := db.Table("topics").Create(&topics)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
