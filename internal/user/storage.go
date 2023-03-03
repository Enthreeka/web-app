package user

import (
	"context"
	"web/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, login string, password string) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
}
