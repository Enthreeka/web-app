package account

import (
	"context"
	"web/internal/entity"
)

type Service interface {
	DeleteTask(ctx context.Context, task *entity.Task) error
	CreateTask(ctx context.Context, task *entity.Task) (int, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	GetTask(ctx context.Context, userID string) ([]string, []string, []string, error)
}
