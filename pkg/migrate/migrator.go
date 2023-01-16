package migrate

import (
	"github.com/shanedoc/gohub/pkg/database"
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
		Folder:   "databases/migrations/",
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
