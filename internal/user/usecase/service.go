package usecase

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) SignUp(ctx context.Context, user *entity.User) error {
	hasedBytes := string(HashPassword(user.Login, user.Password))
	user.Password = hasedBytes

	err := s.repository.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user : %v", err)
	}

	validToken, err := s.GenerateToken(ctx, user.Id)
	if err != nil {
		log.Fatalf("failed to generate token %v", err)
		return fmt.Errorf("incorrect genereate token")
	}
	user.Token = validToken

	return err
}

func (s *Service) LogIn(ctx context.Context, login string, password string) (entity.User, error) {
	user, err := s.repository.GetUser(ctx, login, password)
	if err != nil {
		fmt.Println(user.Password)
		log.Fatalf("failed to get user %s", err)
		return entity.User{}, fmt.Errorf("failed to get user : %v", err)
	}
	if user == nil {
		log.Fatal("empty user")
		return entity.User{}, fmt.Errorf("user not found")
	}

	hashedSha256Password := sha256Hash(login, password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), hashedSha256Password)
	if err != nil {
		log.Printf("failed to compare hashed password with password %v", err)
		return entity.User{}, err
	}

	validToken, err := s.GenerateToken(ctx, user.Id)
	if err != nil {
		return entity.User{}, fmt.Errorf("incorrect genereate token")
	}
	user.Token = validToken

	return *user, nil
}

func (s *Service) GenerateToken(ctx context.Context, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret-token-gen"))
	if err != nil {
		log.Fatalf("failed to SignedString %v", err)
	}

	// Store the token ID in the database
	tokenID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("failed to create UUID %v ", err)
	}

	//Create token and implement in database
	if err := s.repository.UpdateToken(ctx, tokenID.String(), userID); err != nil {
		log.Fatalf("failed to store token: %v", err)
		return "", fmt.Errorf("failed to store token: %v", err)
	}
	return tokenString, nil
}

func sha256Hash(login, password string) []byte {
	passwordBytes := []byte(password + login)

	hash := sha256.New()
	_, err := hash.Write(passwordBytes)
	if err != nil {
		log.Fatalf("failed to write bytes password in hash %v", err)
		return nil
	}
	hashedSha256Password := hash.Sum(nil)

	return hashedSha256Password
}

//TODO salt and iterations
func HashPassword(login, password string) []byte {
	hashedSha256Password := sha256Hash(login, password)

	bcryptPassword, err := bcrypt.GenerateFromPassword(hashedSha256Password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to generate bcrypt %v", err)
		return nil
	}

	return bcryptPassword
}

func (s *Service) GetAll(ctx context.Context) ([]entity.User, error) {
	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user : %v", err)
	}
	return all, nil
}
