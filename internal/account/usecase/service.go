package usecase

import (
	"context"
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

func (s *Service) Leave(ctx context.Context, userID string) error {

	err := s.repository.SetNullToken(ctx, userID)
	if err != nil {
		log.Fatalf("failed to set null in service %v", err)
		return err
	}

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, task *entity.Task) error {

	err := s.repository.DeleteTask(ctx, task.Id)
	if err != nil {
		log.Printf("failed to delete task in service %v", err)
		return err
	}

	return nil
}
func (s *Service) CreateTask(ctx context.Context, task *entity.Task) (int, error) {

	id, err := s.repository.CreateTask(ctx, task)
	if err != nil {
		log.Printf("failed to create task %v", err)
		return 0, err
	}

	return id, nil
}

func (s *Service) UpdateTask(ctx context.Context, task *entity.Task) error {

	err := s.repository.UpdateNameTask(ctx, task)
	if err != nil {
		log.Printf("failed to add name task %v", err)
		return err
	}

	err = s.repository.UpdateDescriptionTask(ctx, task)
	if err != nil {
		log.Printf("failed to add task %v", err)
		return err
	}

	return nil
}

func (s *Service) GetTask(ctx context.Context, userID string) ([]string, []string, []string, error) {

	id, name, description, err := s.repository.GetTask(ctx, userID)
	if err != nil {
		log.Printf("failed to get name and description from db %v", err)
		return nil, nil, nil, err
	}

	return id, name, description, err
}
