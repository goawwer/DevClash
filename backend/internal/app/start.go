package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goawwer/devclash/config"
	"github.com/goawwer/devclash/internal/adapter/database"
	"github.com/goawwer/devclash/internal/adapter/database/repository"
	"github.com/goawwer/devclash/internal/adapter/storage"
	"github.com/goawwer/devclash/internal/controller"
	"github.com/goawwer/devclash/internal/usecase"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/server"
)

func Start(ctx context.Context, cfg *config.Config) {
	logger.Info("starting application")

	if err := database.Init(ctx, &cfg.Database); err != nil {
		logger.Error("failed to initialize database in start function: ", err)
	}

	defer func() {
		logger.Info("closing database connection")
		if err := database.Close(); err != nil {
			logger.Error("cannot close database connection: ", err)
		}
	}()

	logger.Info("initialize redis")
	if err := storage.Init(ctx, &cfg.Storage); err != nil {
		logger.Error("cannot initialize redis: ", err)
	}

	logger.Info("create instance of whole application repository")
	repository := repository.NewRepository(database.Get())

	logger.Info("create instance of whole application usecase")
	usecase := usecase.NewUsecase(repository)

	logger.Info("initialize middleware config")
	middleware.InitAuthConfig(cfg.JWT.Secret, repository)

	logger.Info("implementing router")
	r := controller.Router(usecase)

	logger.Info("create http server instance")
	httpServer := server.New(r, &cfg.Server)

	logger.Info("starting server on port: ", cfg.Server.Port)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("failed to start server: ", err)
		}
	}()

	logger.Info("implement worker for deleting extra values from database")
	go func() {
		ticker := time.NewTicker(time.Minute * 20)
		for range ticker.C {
			if err := repository.CleanupExpiredRefreshTokens(ctx); err != nil {
				logger.Error("failed to delete value: ", err)
			}
		}
	}()

	logger.Info("application successfully started on address: ", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Info("shutdown by parent context")
	case <-sig:
		logger.Info("graceful by termination signal, init graceful shutdown")
	}

	logger.Info("creating context for graceful shutdown")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			logger.Info("graceful shutdown timed out")
		} else {
			logger.Error("greceful shutdown err: ", err)
		}
	} else {
		logger.Info("graceful shutdown complete")
	}

}
