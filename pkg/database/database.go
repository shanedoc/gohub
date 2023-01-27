//package 数据库操作
package database

import (
	"database/sql"
	"fmt"

	"github.com/shanedoc/gohub/pkg/config"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

//连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	//使用gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	//处理错误
	if err != nil {
		fmt.Println(err.Error())
	}
	//获取底层的sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return dbname
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		deleteAllSqliteTables()
	default:
		panic("error")
	}
	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}

	//读取所有数据
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'").Error
	if err != nil {
		return err
	}
	//删除所有的数据
	for _, t := range tables {
		err := DB.Migrator().DropTable(t)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	tables := []string{}
	//读取所有的数据库
	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}
	//暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")
	//删除所有的表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	//开启外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
