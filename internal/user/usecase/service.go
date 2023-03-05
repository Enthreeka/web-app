package usecase

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
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

	if user == nil {
		return entity.User{}, fmt.Errorf("user not found")
	}
	if user.Password != password {
		return entity.User{}, fmt.Errorf("incorrect password")
	}

	validToken, err := s.GenerateToken(user.Id)
	if err != nil {
		return entity.User{}, fmt.Errorf("incorrect genereate token")
	}

	fmt.Println(validToken)
	return *user, nil
}

func (s *Service) GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatalf("failed to SignedString %v", err)
	}

	// Store the token ID in the database
	//if err := s.repository.CreateToken(tokenID.String(), userID); err != nil { TODO Create func in db/postgresql CreateToken
	//	return "", fmt.Errorf("failed to store token: %v", err)
	//}
	return tokenString, nil
}

func (s *Service) GetAll(ctx context.Context) ([]entity.User, error) {
	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user : %v", err)
	}
	return all, nil
}

//
//func (s *Service) CreateUser(ctx context.Context, user *entity.User) error {
//	err := s.repository.Create(ctx, user)
//	if err != nil {
//		return fmt.Errorf("failed to create user : %v", err)
//	}
//	return err
//}
//
//func (s *Service) GetOne(ctx context.Context, login string, password string) (entity.User, error) {
//	one, err := s.repository.GetUser(ctx, login, password)
//	if err != nil {
//		return entity.User{}, fmt.Errorf("failed to create user : %v", err)
//	}
//	return one, nil
//}
