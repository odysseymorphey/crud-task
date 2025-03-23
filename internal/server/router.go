package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/odysseymorphey/crud-task/internal/repository"
	"github.com/odysseymorphey/crud-task/internal/server/handlers"
)

type Router struct {
	baseHandler *handlers.BaseHandler
}

func NewRouter(repo repository.Repository) *Router {
	return &Router{
		baseHandler: handlers.NewBaseHandler(repo),
	}
}

func (r *Router) RegisterRoutes(srv *fiber.App) {
	api := srv.Group("/api")

	{
		api.Get("/users", r.baseHandler.GetUsers)
		api.Get("/users/:user_id", r.baseHandler.GetUsers)
		api.Post("/users", r.baseHandler.CreateUser)
		api.Put("/users", r.baseHandler.UpdateUser)
		api.Delete("users/:user_id", r.baseHandler.DeleteUser)
	}
}