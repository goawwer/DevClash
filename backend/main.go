package main

import (
	"context"

	"github.com/goawwer/devclash/config"
	"github.com/goawwer/devclash/internal/app"
	"github.com/goawwer/devclash/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		panic("failed to load config")
	}

	logger.Init(ctx, &cfg.Logger)

	app.Start(ctx, cfg)
}
