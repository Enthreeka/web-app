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

func NewAccountRepository(db *pgxpool.Pool) account.Repository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateTask(ctx context.Context, task *entity.Task) error {
	query := `INSERT INTO tasks
				(account_id,name_task,description_task)
				VALUES
					($1,$2,$3)
				RETURNING id`

	if err := r.db.QueryRow(ctx, query, task.AccountId, task.NameTask, task.DescriptionTask).Scan(&task.Id); err != nil {
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
func (r *accountRepository) UpdateDescriptionTask(ctx context.Context, task *entity.Task) error {
	query := `UPDATE tasks
			SET description_task = $1
			WHERE account_id = $2`

	_, err := r.db.Exec(ctx, query, task.DescriptionTask, task.AccountId)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

func (r *accountRepository) UpdateNameTask(ctx context.Context, task *entity.Task) error {
	query := `UPDATE tasks
			SET name_task = $1
			WHERE account_id = $2`

	_, err := r.db.Exec(ctx, query, task.NameTask, task.AccountId)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

//TODO Change userID on accountID
func (r *accountRepository) GetTask(ctx context.Context, userID int) ([]string, []string, error) {

	query := `SELECT name_task , description_task
					FROM tasks
				WHERE account_id = $1`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tasks []entity.Task

	for rows.Next() {

		var task entity.Task
		err := rows.Scan(&task.NameTask, &task.DescriptionTask)
		if err != nil {
			return nil, nil, err
		}
		tasks = append(tasks, task)
	}
	var name []string
	var descriptions []string

	for _, task := range tasks {
		name = append(name, task.NameTask)
		descriptions = append(descriptions, task.DescriptionTask)
	}
	return name, descriptions, nil
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
