package verifycode

import (
	"time"

	"github.com/shanedoc/gohub/pkg/app"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/redis"
)

//redisStore实现verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	//本地环境调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

//实现get方法
func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

//verfiy实现 verfiycode.Store interface的verfiy方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
