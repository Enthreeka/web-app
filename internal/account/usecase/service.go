package usecase

import (
	"context"
	"fmt"
	"log"
	"web/internal/account"
	"web/internal/entity"
)

type Service struct {
	repository account.Repository
}

func NewAccountService(repository account.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) AddTask(ctx context.Context, account *entity.Account) error {

	err := s.repository.CreateNameTask(ctx, account)
	if err != nil {
		fmt.Printf("failed to add name task %v", err)
		return err
	}

	err = s.repository.CreateTask(ctx, account)
	if err != nil {
		fmt.Printf("failed to add task %v", err)
		return err
	}

	return nil
}

func (s *Service) GetTask(ctx context.Context, userID int) ([]string, []string, error) {

	name, description, err := s.repository.GetTask(ctx, userID)
	if err != nil {
		log.Printf("failed to get name and description from db %v", err)
		return nil, nil, err
	}

	return name, description, err
}
