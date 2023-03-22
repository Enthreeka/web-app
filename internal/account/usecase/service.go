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

func (s *Service) GetName(ctx context.Context, userID string) (string, error) {

	name, err := s.repository.GetName(ctx, userID)
	if name == "" {
		name = ""
		return name, nil
	}
	if err != nil {
		log.Fatalf("failed to get name in service %v", err)
		return "", nil
	}

	return name, nil
}

func (s *Service) SaveName(ctx context.Context, userID string, name string) error {

	err := s.repository.UpdateNameUser(ctx, userID, name)
	if err != nil {
		log.Printf("failed to update name user in service %v", err)
		return err
	}

	return nil
}

func (s *Service) GetTask(ctx context.Context, id int) (string, string, error) {
	nameTask, descriptionTask, err := s.repository.GetTask(ctx, id)
	if err != nil {
		log.Fatalf("failed to get task in service %v", err)
		return "", "", nil
	}

	return nameTask, descriptionTask, nil
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

func (s *Service) UpdateDescriptionTask(ctx context.Context, descriptionTask string, id int) error {
	err := s.repository.UpdateDescriptionTask(ctx, descriptionTask, id)
	if err != nil {
		log.Printf("failed to update description in service %v", err)
		return err
	}

	return nil
}
func (s *Service) UpdateNameTask(ctx context.Context, nameTask string, id int) error {
	err := s.repository.UpdateNameTask(ctx, nameTask, id)
	if err != nil {
		log.Printf("failed to update name in service %v", err)
		return err
	}

	return nil
}

func (s *Service) GetTasks(ctx context.Context, userID string) ([]string, []string, []string, error) {
	id, name, description, err := s.repository.GetTasks(ctx, userID)
	if err != nil {
		log.Printf("failed to get name and description from db %v", err)
		return nil, nil, nil, err
	}

	return id, name, description, err
}
