package migrations

import (
	"database/sql"

	"github.com/shanedoc/gohub/app/models"
	"github.com/shanedoc/gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {
	type User struct {
		models.BaseModel

		City         string `gorm:"type:varchar(10);"`
		Introduction string `gorm:"type:varchar(255);"`
		Avator       string `gorm:type:"varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2023_02_03_173707_add_fields_to_user", up, down)
}
