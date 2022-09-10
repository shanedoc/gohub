package captcha

import (
	"errors"
	"time"

	"github.com/shanedoc/gohub/app/app"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/redis"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

//Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	//区分环境进行调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

//Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	value := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Delete(key)
	}
	return value
}

//Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verfiy(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
