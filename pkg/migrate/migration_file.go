package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator, *sql.DB)

//migrationFiles 所有的迁移文件数组
var migrationFiles []MigrationFile

//MigrationFile 代表单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

//Add 新增一个迁移文件,所有迁移文件都需要调用此方法进行注册
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

//判断迁移是否被执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
