package seed

import (
	"gorm.io/gorm"
)

//处理数据库填充逻辑

//存放所有的seeder
var seeders []Seeder

// 按照顺序执行Seeder数组
// 支持一些必须按顺序执行的 seeder，例如 topic 创建的
// 时必须依赖于 user, 所以 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

type SeedFunc func(*gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeedFunc
	Name string
}

//Add注册到seeders数组
func Add(name string, fn SeedFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

//SetRunOrder 设置『按顺序执行的 Seeder 数组』
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
