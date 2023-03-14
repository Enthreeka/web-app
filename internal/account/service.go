package account

import (
	"context"
	"web/internal/entity"
)

type Service interface {
	AddTask(ctx context.Context, account *entity.Account) error
	GetTask(ctx context.Context, userID int) ([]string, []string, error)
}
