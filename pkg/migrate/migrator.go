package migrate

import (
	"os"

	"github.com/shanedoc/gohub/pkg/console"
	"github.com/shanedoc/gohub/pkg/database"
	"github.com/shanedoc/gohub/pkg/file"
	"gorm.io/gorm"
)

// Package migrate 处理数据库迁移

//Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

//Migration 对应数据的 migrations表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

//创建 Migrator 用来执行迁移操作
func NewMigrator() *Migrator {
	//初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	//不存在就创建
	migrator.createMigrationsTable()
	return migrator
}

//创建 migrations表
func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}
	//不存在再创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

//up-执行所有未迁移过的文件
func (migrator *Migrator) Up() {
	//读取所有迁移文件确保按时间排序
	migrateFiles := migrator.readAllMigrationFiles()
	//获取当前批次的值
	batch := migrator.getBatch()
	//获取所有要迁移的数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	//可以通过此值来判断数据库是否已是最新
	refund := false

	//对迁移文件进行遍历,如果没有执行过就执行up回调
	for _, mfile := range migrateFiles {
		//对比文件名称查看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			refund = true
		}
	}
	if !refund {
		console.Success("database is up to date.")
	}
}

//rollback 回滚操作
func (migrator *Migrator) Rollback() {
	//获取最后一批的迁移数据
	lastMigration := &Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)
	//回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

//回退迁移，按照倒序执行迁移的 down 方法
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	//标记是否执行了迁移回退的操作
	runed := false
	for _, _migration := range migrations {
		//提示
		console.Warning("rollback" + _migration.Migration)
		//执行迁移文件down
		// 通过迁移文件的名称获取『MigrationFile』对象
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}
		runed = true
		//回退成功删除记录
		migrator.DB.Delete(&_migration)

		//打印运行状态
		console.Success("finish " + mfile.FileName)
	}
	return runed
}

// 从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {

	// 读取 database/migrations/ 目录下的所有文件
	// 默认是会按照文件名称进行排序
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {

		// 去除文件后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())

		// 通过迁移文件的名称获取『MigrationFile』对象
		mfile := getMigrationFile(fileName)

		// 加个判断，确保迁移文件可用，再放进 migrateFiles 数组中
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	// 返回排序好的『MigrationFile』数组
	return migrateFiles
}

//获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	//初识默认值
	batch := 1
	//获取最后执行的迁移记录
	lastMigration := &Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

//执行迁移,执行迁移up的方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// 执行 up 区块的 SQL
	if mfile.Up != nil {
		// 友好提示
		console.Warning("migrating " + mfile.FileName)
		// 执行 up 方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}

	// 入库
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

//回滚所有迁移
func (migrator *Migrator) Reset() {
	migrations := []Migration{}
	//按照倒序读取所有的迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)
	//回滚所有迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

//Refresh 回滚所有迁移，并运行所有迁移
func (migrator *Migrator) Refresh() {
	//回滚
	migrator.Reset()
	//再次执行所有迁移
	migrator.Up()
}
