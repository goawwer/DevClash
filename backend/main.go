package main

import (
	"context"

	"github.com/goawwer/devclash/config"
	"github.com/goawwer/devclash/internal/app"
	"github.com/goawwer/devclash/pkg/logger"
)

//	@title			Dev clash API
//	@version		1.0
//	@description	IT events platform
//
// @securityDefinitions.apikey  CookieAuth
// @in                          cookie
// @name                        access
// @description                 JWT access token stored in HttpOnly cookie named `access`

// @host		localhost:8080
// @BasePath	/
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
