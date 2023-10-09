package sloger

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"golang.org/x/exp/slog"
	log "log/slog"
)

type appAuthLogger struct {
	cfg    *config.Config
	logger *log.Logger
}

func NewAppAuthLogger(cfg *config.Config) *appAuthLogger {
	return &appAuthLogger{
		cfg: cfg,
	}
}

var loggerLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func (a *appAuthLogger) getLoggerLevel(cfg *config.Config) slog.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return slog.LevelDebug
	}
	return level
}
