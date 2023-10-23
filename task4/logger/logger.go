package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerSingleton - структура синглтона для логгера
type loggerSingleton struct {
	logger *zap.Logger
}

var instance *loggerSingleton = nil

// GetLoggerInstance - функция для получения экземпляра логгера (синглтона)
func GetLoggerInstance() *loggerSingleton {
	if instance == nil {
		instance = &loggerSingleton{
			logger: initLogger(),
		}
	}

	return instance
}

// initLogger - функция для инициализации логгера
func initLogger() *zap.Logger {
	// Создаем директорию для логов, если она не существует
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	// Определяем путь к файлу для логов
	logFilePath := filepath.Join(logDir, "logfile.log")

	config := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{logFilePath},
		ErrorOutputPaths: []string{"error.log"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "timestamp",
			LevelKey:     "level",
			MessageKey:   "message",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		return nil
	}

	return logger
}

// Info - функция для записи информационных сообщений
func (ls *loggerSingleton) Info(message string) {
	ls.logger.Info(message)
}

// Debug - функция для записи отладочных сообщений
func (ls *loggerSingleton) Debug(message string) {
	ls.logger.Debug(message)
}

// Error - функция для записи сообщений об ошибках
func (ls *loggerSingleton) Error(message string) {
	ls.logger.Error(message)
}
