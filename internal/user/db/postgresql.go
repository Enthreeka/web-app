package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"web/internal/entity"
	"web/internal/user"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) SetNullToken(ctx context.Context, userID string) error {
	query := `UPDATE users
				SET token = 0
					WHERE user_id = $1`

	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		log.Fatalf("failed to set null to token: %v", err)
		return fmt.Errorf("failed to set null to token: %v", err)
	}

	return nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `INSERT INTO users 
				 (login, password ) 
				VALUES
					($1, $2)
			RETURNING id`

	if err := r.db.QueryRow(ctx, query, user.Login, user.Password).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s,Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateAccount(ctx context.Context, account *entity.Account) error {
	query := `INSERT INTO account
				(user_id)
				VALUES
					($1)
				RETURNING id`

	if err := r.db.QueryRow(ctx, query, account.UserId).Scan(&account.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s,Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	query := `SELECT id, login, password 
				FROM users`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	users := make([]entity.User, 0)

	for rows.Next() {
		var user entity.User

		err = rows.Scan(&user.Id, &user.Login, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//TODO set time line for keeping life tokenID
func (r *userRepository) UpdateToken(ctx context.Context, tokenID string, userID string) error {
	query := `UPDATE users
				SET token = $1
				WHERE id = $2`

	_, err := r.db.Exec(ctx, query, tokenID, userID)
	if err != nil {
		log.Fatalf("failed to create token: %v", err)
		return fmt.Errorf("failed to create token: %v", err)
	}

	return nil
}

func (r *userRepository) GetUser(ctx context.Context, login string, password string) (*entity.User, error) {
	if login == "" || password == "" {
		log.Fatal("get empty login or password")
	}
	var user entity.User

	query := `SELECT id, login, password 
				FROM users 
					WHERE 
				login = $1`

	err := r.db.QueryRow(ctx, query, login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		log.Printf("failed to query user %v", err)
		return &entity.User{}, err
	}

	return &user, nil
}
