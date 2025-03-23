package repository

import (
	"github.com/gofiber/fiber/v3"
	"github.com/odysseymorphey/crud-task/internal/storage"
)

type Repository interface {
	GetUser(fiber.Ctx) error
	CreateUser(fiber.Ctx) error
	DeleteUser(fiber.Ctx) error
	UpdateUser(fiber.Ctx) error
}

type repository struct {
	storage *storage.PostgresDB
}

func NewRepository(storage *storage.PostgresDB) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetUser(ctx fiber.Ctx) error {
	return nil
}

func (r *repository) CreateUser(ctx fiber.Ctx) error {
	return nil
}

func (r *repository) DeleteUser(ctx fiber.Ctx) error {
	return nil
}

func (r *repository) UpdateUser(ctx fiber.Ctx) error {
	return nil
}