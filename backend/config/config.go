package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/goawwer/devclash/internal/adapter/database"
	"github.com/goawwer/devclash/internal/adapter/redis"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/mail"
	"github.com/goawwer/devclash/pkg/s3"
	"github.com/goawwer/devclash/pkg/server"
)

type Config struct {
	Database database.Config
	S3       s3.Config
	Server   server.Config
	Mail     mail.Config
	JWT      middleware.Config
	Storage  redis.Config
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
