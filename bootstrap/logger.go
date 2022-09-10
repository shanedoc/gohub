package bootstrap

import (
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/logger"
)

//初始化logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
