package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger
var currentLogFile *os.File
var currentLogDay string

func InitLogger(level, format, output string) error {
	// 配置日志级别
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 配置编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	var encoder zapcore.Encoder
	if format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 配置输出
	var sink zapcore.WriteSyncer
	if output == "stdout" {
		sink = zapcore.AddSync(os.Stdout)
	} else {
		logFile, err := openLogFile()
		if err != nil {
			return err
		}
		currentLogFile = logFile
		currentLogDay = time.Now().Format("2006-01-02")
		sink = zapcore.AddSync(logFile)

		// 启动日志轮转检查
		go checkLogRotation()
	}

	core := zapcore.NewCore(encoder, sink, zapLevel)
	logger := zap.New(core)
	globalLogger = logger.Sugar()

	return nil
}

func openLogFile() (*os.File, error) {
	now := time.Now()
	logDir := "logs"

	// 创建年月日目录
	yearDir := filepath.Join(logDir, now.Format("2006"))
	monthDir := filepath.Join(yearDir, now.Format("01"))
	if err := os.MkdirAll(monthDir, 0755); err != nil {
		return nil, fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 生成日志文件名
	logFileName := fmt.Sprintf("app-%s.log", now.Format("2006-01-02"))
	logFilePath := filepath.Join(monthDir, logFileName)

	// 打开日志文件
	return os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}

func checkLogRotation() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		today := time.Now().Format("2006-01-02")
		if today != currentLogDay {
			// 重新初始化日志器
			if err := InitLogger("info", "console", "file"); err != nil {
				fmt.Printf("轮转日志失败: %v\n", err)
				continue
			}

			// 关闭旧的日志文件
			if currentLogFile != nil {
				currentLogFile.Close()
			}
		}
	}
}

func GetLogger() *zap.SugaredLogger {
	if globalLogger == nil {
		// 如果没有初始化，使用默认配置
		logger, _ := zap.NewProduction()
		globalLogger = logger.Sugar()
	}
	return globalLogger
}

// 便捷方法
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
