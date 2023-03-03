package usecase

import (
	"context"
	"fmt"
	"log"
	"web/internal/entity"
	"web/internal/user"
)

type Service struct {
	repository user.Repository
}

func NewService(repository user.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) SignIn(ctx context.Context) error {

	return nil
}

func (s *Service) LogIn(ctx context.Context, login string, password string) (entity.User, error) {

	user, err := s.repository.GetUser(ctx, login, password)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create user : %v", err)
	}

	if user.Login != login || user.Password != password {
		log.Fatalf("input data incorrect %v", err)
	}

	//validToken , err :=

	return user, nil
}

func (s *Service) GetAll(ctx context.Context) ([]entity.User, error) {
	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user : %v", err)
	}
	return all, nil
}

func (s *Service) CreateUser(ctx context.Context, user *entity.User) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user : %v", err)
	}
	return err
}

func (s *Service) GetOne(ctx context.Context, login string, password string) (entity.User, error) {
	one, err := s.repository.GetUser(ctx, login, password)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create user : %v", err)
	}
	return one, nil
}
