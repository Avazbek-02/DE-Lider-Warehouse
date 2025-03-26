package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	v1 "github.com/Avazbek-02/DE-Lider-Warehouse/internal/controller/http/v1"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/repository"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/usecase"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/httpserver"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/logger"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/postgres"
	rediscache "github.com/golanguzb70/redis-cache"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Initialize database connection
	db, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer db.Close()

	// Initialize Redis cache
	redis, err := rediscache.New(&rediscache.Config{
		RedisHost: cfg.Redis.RedisHost,
		RedisPort: cfg.Redis.RedisPort,
	})
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rediscache.New: %w", err))
	}

	// Initialize repositories
	repos := repository.New(db.Pool)

	// Initialize use cases
	uc := usecase.New(cfg, repos)

	// Initialize HTTP server and routes
	handler := gin.New()
	v1.NewRouter(handler, l, cfg, uc, redis)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	l.Info(fmt.Sprintf("Warehouse API server started on port: %s", cfg.HTTP.Port))

	// Waiting for interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
