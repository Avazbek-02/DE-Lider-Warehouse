package usecase

import (
	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/usecase/repo"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/logger"
	"github.com/Avazbek-02/DE-Lider-Warehouse/pkg/postgres"
)

// UseCase -.
type UseCase struct {
	UserRepo       UserRepoI
	SessionRepo    SessionRepoI
	RoomsRepo      RoomsRepoI
	RoomReviewRepo RoomReviewRepoI
}

// New -.
func New(pg *postgres.Postgres, config *config.Config, logger *logger.Logger) *UseCase {
	return &UseCase{
		UserRepo:       repo.NewUserRepo(pg, config, logger),
		SessionRepo:    repo.NewSessionRepo(pg, config, logger),
		RoomsRepo:      repo.NewRoomsRepo(pg, config, logger),
		RoomReviewRepo: repo.NewRoomReviewRepo(pg, config, logger),
	}
}
