package bootstrap

import (
	"fmt"

	"github.com/shanedoc/gohub/pkg/cache"
	"github.com/shanedoc/gohub/pkg/config"
)

//setCache缓存
func SetupCache() {
	//初始化缓存专用的redis client 使用专属的缓存DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)
	cache.InitWithCacheStore(rds)

}
