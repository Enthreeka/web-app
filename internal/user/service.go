package user

import (
	"context"
	"web/internal/entity"
)

type Service interface {
	SignUp(ctx context.Context, user *entity.User) (*entity.User, error)
	Leave(ctx context.Context, userID string) error
	LogIn(ctx context.Context, login string, password string) (entity.User, error)
	GenerateToken(ctx context.Context, userID string) (string, error)
}
