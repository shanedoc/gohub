package bootstrap

import (
	"fmt"

	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/redis"
)

//setup redis
//初始化redis
func SetupRedis() {
	//建立redis链接
	redis.ConnectReids(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)

}
