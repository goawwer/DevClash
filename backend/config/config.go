package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/goawwer/devclash/internal/adapter/database"
	"github.com/goawwer/devclash/internal/adapter/storage"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/server"
)

type Config struct {
	Database database.Config
	Server   server.Config
	JWT      middleware.Config
	Storage  storage.Config
	Logger   logger.Config
}

func New() (*Config, error) {
	var cfg Config

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load config")
	}

	return &cfg, nil
}
