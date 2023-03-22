package account

import (
	"context"
	"web/internal/entity"
)

type Repository interface {
	DeleteTask(ctx context.Context, taskID int) error
	CreateTask(ctx context.Context, task *entity.Task) (int, error)
	UpdateDescriptionTask(ctx context.Context, descriptionTask string, id int) error
	UpdateNameTask(ctx context.Context, nameTask string, id int) error
	SetNullToken(ctx context.Context, userID string) error
	GetTasks(ctx context.Context, userID string) ([]string, []string, []string, error)
	UpdateTask(ctx context.Context, id string) error
	GetTask(ctx context.Context, id int) (string, string, error)
	UpdateNameUser(ctx context.Context, userID string, name string) error
	GetName(ctx context.Context, userID string) (string, error)
	AddByneriPhoto(ctx context.Context, userID string, imgByte []byte) error
}
