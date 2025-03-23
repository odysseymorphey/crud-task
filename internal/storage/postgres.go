package storage

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
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

func (db *PostgresDB) GetUser(ctx fiber.Ctx) error {
	return nil
}

func (db *PostgresDB) CreateUser(ctx fiber.Ctx) error {
	return nil
}

func (db *PostgresDB) DeleteUser(ctx fiber.Ctx) error {
	return nil
}

func (db *PostgresDB) UpdateUser(ctx fiber.Ctx) error {
	return nil
}