package user

import (
	"context"
	"web/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	CreateAccount(ctx context.Context, account *entity.Account) error
	GetUser(ctx context.Context, login string, password string) (*entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	UpdateToken(ctx context.Context, tokenID string, userID string) error
}
