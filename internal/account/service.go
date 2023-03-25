package account

import (
	"context"
	"web/internal/entity"
)

type Service interface {
	DeleteTask(ctx context.Context, task *entity.Task) error
	CreateTask(ctx context.Context, task *entity.Task) (int, error)
	UpdateDescriptionTask(ctx context.Context, descriptionTask string, id int) error
	UpdateNameTask(ctx context.Context, nameTask string, id int) error
	SaveName(ctx context.Context, userID string, name string) error
	GetTasks(ctx context.Context, userID string) ([]string, []string, []string, error)
	Leave(ctx context.Context, userID string) error
	GetTask(ctx context.Context, id int) (string, string, error)
	GetName(ctx context.Context, userID string) (string, error)
	AddPhoto(ctx context.Context, userID string, imgByte []byte) error
	GetPhoto(ctx context.Context, userID string) ([]byte, error)
}
