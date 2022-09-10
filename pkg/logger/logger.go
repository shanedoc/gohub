package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shanedoc/gohub/app/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//zap log

var Logger *zap.Logger

//logigger初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)
	//设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}
	//初始化core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)
	//初始化logger
	Logger = zap.New(core,
		zap.AddCaller(),                   //调用文件和行号
		zap.AddCallerSkip(1),              //调用文件时去掉一层
		zap.AddStacktrace(zap.ErrorLevel), //error时显示stacktrace
	)

	//将自定义的logger替换成全局logger
	//zap.L().Fatal()调用时,使用自定义的Logger
	zap.ReplaceGlobals(Logger)

}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {

	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	// 本地环境配置
	if app.IsLocal() {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

//日志记录媒介,os.Stdout和日志
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	if logType == "daily" {
		logname := time.Now().Format("2006-12-01.log")
		filename = strings.ReplaceAll(filename, "log.log", logname)
	}

	//滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	//配置输出介质
	if app.IsLocal() {
		//local  终端打印写入文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		//记录文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

//调试专用方法,不会中断程序执行,会将warning信息打印在中断上,
//第一个参数使用json.Marshal,第二个参数可选
//logger.Dump(user.User{Name:"test"})
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	//判断第二个参数
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

//error不为空时,记录error等级日志
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occured:", zap.Error(err))
	}
}

func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Warn Occured:", zap.Error(err))
	}
}

func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Info Occured:", zap.Error(err))
	}
}

//debug调试日志
// logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

//info
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

//warn
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

//error
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

//fatal
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

//语法糖
// DebugString 记录一条字符串类型的 debug 日志，调用示例：
//         logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

//语法糖
// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
//         logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
	//将数据格式转成json格式
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
