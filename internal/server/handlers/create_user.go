package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/odysseymorphey/crud-task/internal/models"
)

func (h *BaseHandler) CreateUser(c fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error": err.Error(),
		})
	}

	if err := h.repo.CreateUser(&user, c.Context()); err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}