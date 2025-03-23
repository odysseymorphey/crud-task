package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/odysseymorphey/crud-task/internal/models"
)

type PostgresDB struct {
	storage *sql.DB
}

func NewPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	const op = "storage.NewConnection"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Printf("connStr: %s\n", connStr)

	var err error
	storage, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open database: %v", op, err)
	}

	if err := storage.Ping(); err != nil {
		return nil, fmt.Errorf("%s: failed to ping database: %v", op, err)
	}

	fmt.Printf("Open database: %s\n", op)

	return &PostgresDB{
		storage: storage,
	}, nil
}

func (db *PostgresDB) GetUser(user_id string, ctx context.Context) (*models.User, error) {
	const op = "storage.GetUser"
	
	query := `SELECT id, name FROM users WHERE id=$1`
	row := db.storage.QueryRowContext(ctx, query, user_id)

	var user models.User

	err := row.Scan(&user.Id, &user.Name, &user.SecondName)
	if err != nil {
		fmt.Printf("%s: failed to scan row: %v", op, err)
		return nil, err
	}

	return &user, nil
}

func (db *PostgresDB) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "storage.GetUsers"
	
	query := `SELECT * FROM users`
	rows, err := db.storage.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.SecondName); err != nil {
			fmt.Printf("%s: failed to scan row: %v", op, err)
			return nil, err
		}

		users = append(users, user)
	}



	return users, nil
}

func (db *PostgresDB) CreateUser(user *models.User, ctx context.Context) error {
	const op = "storage.CreateUser"

	query := `INSERT INTO users (name, second_name) VALUES ($1, $2)`

	_, err := db.storage.ExecContext(ctx, query, user.Name, user.SecondName)
	if err != nil {
		return fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return nil
}

func (db *PostgresDB) DeleteUser(user_id string, ctx context.Context) error {
	const op = "storage.DeleteUser"

	query := `DELETE FROM users WHERE user_id = $1`

	res, err := db.storage.ExecContext(ctx, query, user_id)
	if err != nil {
		return fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: user not found", op)
	}

	return nil
}

func (db *PostgresDB) UpdateUser(user models.User, ctx context.Context) error {
	const op = "storage.UpdateUser"

	query := `UPDATE users SET name = $1, second_name = $2 WHERE id = $3`

	res, err := db.storage.ExecContext(ctx, query, user.Name, user.SecondName, user.Id)
	if err != nil {
		return fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: user not found", op)
	}

	return nil
}

func (db *PostgresDB) Close() error {
	const op = "storage.Close"

	if err := db.storage.Close(); err != nil {
		return fmt.Errorf("%s: failed to close database: %v", op, err)
	}

	return nil
}