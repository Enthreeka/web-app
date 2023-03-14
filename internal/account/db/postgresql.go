package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"web/internal/account"
	"web/internal/entity"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) account.Repository {
	return &accountRepository{db: db}
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
func (r *accountRepository) CreateTask(ctx context.Context, account *entity.Account) error {
	query := `UPDATE account
			SET description_task = $1
			WHERE user_id = $2`

	_, err := r.db.Exec(ctx, query, account.DescriptionTask, account.UserId)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

func (r *accountRepository) CreateNameTask(ctx context.Context, account *entity.Account) error {
	query := `UPDATE account
			SET name_task = $1
			WHERE user_id = $2`

	_, err := r.db.Exec(ctx, query, account.NameTask, account.UserId)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

func (r *accountRepository) GetTask(ctx context.Context, userID int) ([]string, []string, error) {

	query := `SELECT name_task , description_task
					FROM account
				WHERE user_id = $1`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var accounts []entity.Account

	for rows.Next() {

		var account entity.Account
		err := rows.Scan(&account.NameTask, &account.DescriptionTask)
		if err != nil {
			return nil, nil, err
		}
		accounts = append(accounts, account)
	}
	var name []string
	var descriptions []string

	for _, account := range accounts {
		name = append(name, account.NameTask)
		descriptions = append(descriptions, account.DescriptionTask)
	}
	return name, descriptions, nil
}

type Acc struct {
	UserId          int
	NameTask        string
	DescriptionTask string
}

func (r *accountRepository) AddEmail(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) AddPhoto(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) FindAll(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
