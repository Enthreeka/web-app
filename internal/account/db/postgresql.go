package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"web/internal/account"
	"web/internal/entity"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) account.Storage {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(ctx context.Context, account *entity.Account) error {
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

func (r *accountRepository) UpdateName(ctx context.Context, account *entity.Account, id int) error {

	query := `UPDATE account
				SET name = $1
				WHERE id = $2
				`

	//err := r.db.QueryRow(ctx, query, account.Name, account.Id).Scan(&account.Name, &account.Id)
	_, err := r.db.Exec(ctx, query, account.Name, id)

	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil

}

func (r *accountRepository) AddEmail(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) AddPhoto(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) AddTask(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) FindAll(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
