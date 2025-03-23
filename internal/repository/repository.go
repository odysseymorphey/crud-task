package repository

import (
	"context"
	"github.com/odysseymorphey/crud-task/internal/models"
	"github.com/odysseymorphey/crud-task/internal/storage"
)

type Repository interface {
	GetUser(user_id string, ctx context.Context) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	CreateUser(user *models.User, ctx context.Context) error
	DeleteUser(user_id string, ctx context.Context) error
	UpdateUser(user models.User, ctx context.Context) error
	Close() error
}

type repository struct {
	storage *storage.PostgresDB
}

func NewRepository(storage *storage.PostgresDB) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetUser(user_id string, ctx context.Context) (*models.User, error) {
	return r.storage.GetUser(user_id, ctx)
}

func (r *repository) GetUsers(ctx context.Context) ([]models.User, error) {
	return r.storage.GetUsers(ctx)
}

func (r *repository) CreateUser(user *models.User, ctx context.Context) error {
	return r.storage.CreateUser(user, ctx)
}

func (r *repository) DeleteUser(user_id string, ctx context.Context) error {
	return r.storage.DeleteUser(user_id, ctx)
}

func (r *repository) UpdateUser(user models.User, ctx context.Context) error {
	return r.storage.UpdateUser(user, ctx)
}

func (r *repository) Close() error {
	return r.storage.Close()
}
