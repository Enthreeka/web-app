package account

import (
	"context"
	"web/internal/entity"
)

type Storage interface {
	CreateAccount(ctx context.Context, account *entity.Account) error
	UpdateName(ctx context.Context, account *entity.Account, id int) error
	AddEmail(ctx context.Context) error
	AddPhoto(ctx context.Context) error
	AddTask(ctx context.Context) error
	FindAll(ctx context.Context) error
	//DeleteName(ctx context.Context) error
	//DeleteEmail(ctx context.Context) error
	//DeletePhoto(ctx context.Context) error
	//DeleteTask(ctx context.Context) error
}
