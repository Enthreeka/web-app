package account

import (
	"context"
	"web/internal/entity"
)

type Service interface {
	CreateTask(ctx context.Context, task *entity.Task) error
	UpdateTask(ctx context.Context, task *entity.Task) error
	GetTask(ctx context.Context, userID int) ([]string, []string, error)
}
