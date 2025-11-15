package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	File string `env:"LOG_FILE"`
}

type fileHook struct {
	File      *os.File
	Formatter logrus.Formatter
}

var logger *logrus.Logger

func Init(ctx context.Context, cfg *Config) error {
	log := logrus.New()

	file, err := os.OpenFile(cfg.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open logs file: %w", err)
	}

	fileFormatter := &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	}

	consoleFormatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	}

	log.SetOutput(os.Stdout)
	log.SetFormatter(consoleFormatter)

	log.AddHook(&fileHook{
		File:      file,
		Formatter: fileFormatter,
	})

	logger = log

	return nil
}

func Debug(args ...any) {
	logger.Debug(args...)
}

func Info(args ...any) {
	logger.Info(args...)
}

func Error(args ...any) {
	logger.Error(args...)
}

func Fatal(args ...any) {
	logger.Fatal(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func (h *fileHook) Fire(entry *logrus.Entry) error {
	line, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.File.Write(line)
	return err
}

func (h *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
