package handlers

import "github.com/gofiber/fiber/v3"

func (h *BaseHandler) GetUsers(c fiber.Ctx) error {
	users, err := h.repo.GetUsers(c.Context())
	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *BaseHandler) GetUser(c fiber.Ctx) error {
	user_id := c.Params("user_id")
	if user_id == "" {
		return c.JSON(fiber.Map{
			"status": fiber.StatusForbidden,
			"error": "missing parameter: user_id",
		})
	}

	user, err := h.repo.GetUser(user_id, c.Context())
	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	c.Response().SetStatusCode(fiber.StatusOK)

	return c.JSON(user)
}