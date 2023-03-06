package user

import (
	"context"
	"web/internal/entity"
)

type Service interface {
	SignIn(ctx context.Context) error
	LogIn(ctx context.Context, login string, password string) (entity.User, error)
	GenerateToken(ctx context.Context, userID int) (string, error)
	//GetAll(ctx context.Context) ([]entity.User, error)
	//GetOne(ctx context.Context, login string, password string) (entity.User, error)
	//CreateUser(ctx context.Context, user *entity.User) error
}
