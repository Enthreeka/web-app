package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strconv"
	"web/internal/account"
	"web/internal/entity"
)

type accountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) account.Repository {
	return &accountRepository{db: db}
}

var (
	task entity.Task
	acc  entity.Account
)

func (r *accountRepository) GetByneriPhoto(ctx context.Context, userID string) ([]byte, error) {
	query := `SELECT photo
					FROM account	
						WHERE user_id = $1`

	err := r.db.QueryRow(ctx, query, userID).Scan(&acc.Photo)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where)
			fmt.Println(newErr)
			return nil, newErr
		}
		return nil, err
	}
	return acc.Photo, nil
}

func (r *accountRepository) AddByneriPhoto(ctx context.Context, userID string, imgByte []byte) error {
	query := `UPDATE account
				SET photo = $1
					WHERE user_id = $2`

	_, err := r.db.Exec(ctx, query, imgByte, userID)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

func (r *accountRepository) GetName(ctx context.Context, userID string) (string, error) {
	query := `SELECT name 	
		          FROM account 
                  WHERE user_id = $1`

	err := r.db.QueryRow(ctx, query, userID).Scan(&acc.Name)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where)
			fmt.Println(newErr)
			return "", newErr
		}
		return "", err
	}
	return acc.Name, nil
}

func (r *accountRepository) UpdateNameUser(ctx context.Context, userID string, name string) error {
	query := `UPDATE account 
				SET name = $1
				WHERE user_id = $2`

	_, err := r.db.Exec(ctx, query, name, userID)
	if err != nil {
		log.Fatalf("failed to update name: %v", err)
		return fmt.Errorf("failed to set null to token: %v", err)
	}
	return nil
}

func (r *accountRepository) GetTask(ctx context.Context, id int) (string, string, error) {
	query := `SELECT name_task, description_task 	
		          FROM tasks 
                  WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(&task.NameTask, &task.DescriptionTask)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where)
			fmt.Println(newErr)
			return "", "", newErr
		}
		return "", "", err
	}
	return task.NameTask, task.DescriptionTask, nil
}

func (r *accountRepository) UpdateTask(ctx context.Context, id string) error {
	query := `UPDATE tasks
		SET name_task = $1 
			WHERE id = $2`

	_, err := r.db.Exec(ctx, query, task.NameTask, id)
	if err != nil {
		log.Fatalf("failed to update task: %v", err)
		return fmt.Errorf("failed to set null to token: %v", err)
	}
	return nil
}

func (r *accountRepository) SetNullToken(ctx context.Context, userID string) error {
	query := `UPDATE users
				SET token = 0
					WHERE id = $1`

	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		log.Fatalf("failed to set null to token: %v", err)
		return fmt.Errorf("failed to set null to token: %v", err)
	}

	return nil
}

func (r *accountRepository) DeleteTask(ctx context.Context, taskID int) error {
	query := `DELETE FROM tasks
	            where 
				id= $1`

	if err := r.db.QueryRow(ctx, query, taskID).Scan(&task.Id); err != nil {
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

func (r *accountRepository) CreateTask(ctx context.Context, task *entity.Task) (int, error) {
	query := `INSERT INTO tasks
				(user_id,name_task,description_task)
				VALUES
					($1,$2,$3)
				RETURNING id`

	if err := r.db.QueryRow(ctx, query, task.AccountId, task.NameTask, task.DescriptionTask).Scan(&task.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s,Detail: %s, Where: %s", pgErr.Error(), pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return 0, nil
		}
		return 0, err
	}
	return task.Id, nil

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
func (r *accountRepository) UpdateDescriptionTask(ctx context.Context, descriptionTask string, id int) error {
	query := `UPDATE tasks
			SET description_task = $1
			WHERE id = $2`

	_, err := r.db.Exec(ctx, query, descriptionTask, id)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

func (r *accountRepository) UpdateNameTask(ctx context.Context, nameTask string, id int) error {
	query := `UPDATE tasks
			SET name_task = $1
			WHERE id = $2`

	_, err := r.db.Exec(ctx, query, nameTask, id)
	if err != nil {
		fmt.Printf("ERROR - %v", err)
	}
	return nil
}

func (r *accountRepository) GetTasks(ctx context.Context, userID string) ([]string, []string, []string, error) {
	query := `SELECT id ,name_task , description_task
					FROM tasks
				WHERE user_id = $1`

	rows, err := r.db.Query(ctx, query, userID)

	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()

	var tasks []entity.Task

	for rows.Next() {

		var task entity.Task
		err := rows.Scan(&task.Id, &task.NameTask, &task.DescriptionTask)
		if err != nil {
			return nil, nil, nil, err
		}
		tasks = append(tasks, task)
	}
	var id []string
	var name []string
	var descriptions []string

	for _, task := range tasks {
		id = append(id, strconv.Itoa(task.Id))
		name = append(name, task.NameTask)
		descriptions = append(descriptions, task.DescriptionTask)
	}
	return id, name, descriptions, nil
}
