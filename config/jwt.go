package config

import "github.com/shanedoc/gohub/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			//过期时间 单位:分钟 一般不超过2h
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),
			//允许刷新时间,单位:分钟,从生成签名开始计时
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),
			//debug模式下过期时间
			"debug_expire_time": 86400,
		}
	})
}
