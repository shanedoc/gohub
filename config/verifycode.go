package config

import "github.com/shanedoc/gohub/pkg/config"

func init() {
	//验证码配置文件
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			//验证码长度
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),

			//过期时间,单位:分钟
			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15),

			//debug模式下过期时间,单位:分钟
			"debug_expire_time": 10080,

			//local开发测试验证码,默认为123456
			"debug_code": 123456,

			//方便本地api调用测试
			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
