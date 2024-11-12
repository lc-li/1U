package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

func InitLogger(level, format, output string) error {
	// 配置日志级别
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 配置输出
	var sink zapcore.WriteSyncer
	if output == "stdout" {
		sink = zapcore.AddSync(os.Stdout)
	} else {
		file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		sink = zapcore.AddSync(file)
	}

	// 配置编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, sink, zapLevel)
	logger := zap.New(core)
	globalLogger = logger.Sugar()

	return nil
}

func GetLogger() *zap.SugaredLogger {
	if globalLogger == nil {
		// 如果没有初始化，使用默认配置
		logger, _ := zap.NewProduction()
		globalLogger = logger.Sugar()
	}
	return globalLogger
}

// 一些便捷方法
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Infof(template string, args ...interface{}) {
	GetLogger().Infof(template, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	GetLogger().Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	GetLogger().Fatalf(template, args...)
}
