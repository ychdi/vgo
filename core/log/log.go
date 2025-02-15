package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// logger 日志类
var logger *zap.Logger

// InitLog 初始化日志
func InitLog() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	defaultLogLevel := zapcore.DebugLevel
	path, err := os.Getwd()
	if err != nil {
		panic("日志目录获取失败" + err.Error())
	}
	logFile, _ := os.OpenFile(path+"/storage/logs/vgo.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)
	writer := zapcore.AddSync(logFile)
	lg := zap.New(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	defer func(lg *zap.Logger) {
		err := lg.Sync()
		if err != nil {
			fmt.Println("日志初始化错误", err)
		}
	}(lg)
	logger = lg
}

// GetLogger 获取日志类
func GetLogger() *zap.Logger {
	return logger
}
