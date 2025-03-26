package handler

import (
	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/usecase"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/logger"
	rediscache "github.com/golanguzb70/redis-cache"
)

// Handler is a structure that contains all the dependencies for the HTTP handlers
type Handler struct {
	logger  *logger.Logger
	cfg     *config.Config
	useCase *usecase.UseCase
	redis   rediscache.RedisCache
}

// NewHandler creates a new Handler instance
func NewHandler(
	logger *logger.Logger,
	cfg *config.Config,
	useCase *usecase.UseCase,
	redis rediscache.RedisCache,
) *Handler {
	return &Handler{
		logger:  logger,
		cfg:     cfg,
		useCase: useCase,
		redis:   redis,
	}
}
