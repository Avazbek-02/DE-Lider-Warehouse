package usecase

import (
	"time"

	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/repository"
)

// UseCase is a structure that contains all use cases
type UseCase struct {
	Warehouse *WarehouseUseCase
}

// New creates a new UseCase instance
func New(
	cfg *config.Config,
	repo *repository.Repository,
) *UseCase {
	return &UseCase{
		Warehouse: NewWarehouseUseCase(
			repo.Warehouse,
			cfg.JWT.Secret,
			time.Duration(cfg.JWT.TTL)*time.Hour,
		),
	}
}
