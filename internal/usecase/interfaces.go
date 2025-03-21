// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/entity"
)

type (
	// UserRepo -.
	UserRepoI interface {
		Create(ctx context.Context, req entity.User) (entity.User, error)
		GetSingle(ctx context.Context, req entity.UserSingleRequest) (entity.User, error)
		GetList(ctx context.Context, req entity.GetListFilter) (entity.UserList, error)
		Update(ctx context.Context, req entity.User) (entity.User, error)
		Delete(ctx context.Context, req entity.Id) error
	}

	// SessionRepo -.
	SessionRepoI interface {
		Create(ctx context.Context, req entity.Session) (entity.Session, error)
		GetSingle(ctx context.Context, req entity.Id) (entity.Session, error)
		GetList(ctx context.Context, req entity.GetListFilter) (entity.SessionList, error)
		Update(ctx context.Context, req entity.Session) (entity.Session, error)
		Delete(ctx context.Context, req entity.Id) error
		UpdateField(ctx context.Context, req entity.UpdateFieldRequest) (entity.RowsEffected, error)
	}

	RoomsRepoI interface {
		Create(ctx context.Context, req entity.Room) (entity.Room, error)
		GetSingle(ctx context.Context, req entity.Id) (entity.Room, error)
		GetList(ctx context.Context, req entity.GetListFilter) (entity.RoomList, error)
		Update(ctx context.Context, req entity.Room) (entity.Room, error)
		Delete(ctx context.Context, req entity.Id) error
	}

	RoomReviewRepoI interface {
		Create(ctx context.Context, req entity.RoomReview) (entity.RoomReview, error)
		GetSingle(ctx context.Context, req entity.Id) (entity.RoomReview, error)
		GetList(ctx context.Context, req entity.GetListFilter) (entity.RoomReviewList, error)
		Update(ctx context.Context, req entity.RoomReview) (entity.RoomReview, error)
		Delete(ctx context.Context, req entity.Id) error
	}
)
