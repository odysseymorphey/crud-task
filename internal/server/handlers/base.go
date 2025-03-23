package handlers

import "github.com/odysseymorphey/crud-task/internal/repository"

type BaseHandler struct {
	repo repository.Repository
}

func NewBaseHandler(r repository.Repository) *BaseHandler {
	return &BaseHandler{
		repo: r,
	}
}