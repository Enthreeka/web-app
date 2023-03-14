package account

import (
	"context"
	"web/internal/entity"
)

type Repository interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	UpdateDescriptionTask(ctx context.Context, task *entity.Task) error
	UpdateNameTask(ctx context.Context, task *entity.Task) error
	GetTask(ctx context.Context, userID int) ([]string, []string, error)
	UpdateName(ctx context.Context, account *entity.Account, id int) error
	AddEmail(ctx context.Context) error
	AddPhoto(ctx context.Context) error
	FindAll(ctx context.Context) error
	//DeleteName(ctx context.Context) error
	//DeleteEmail(ctx context.Context) error
	//DeletePhoto(ctx context.Context) error
	//DeleteTask(ctx context.Context) error
}
