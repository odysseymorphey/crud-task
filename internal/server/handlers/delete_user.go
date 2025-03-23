package handlers

import "github.com/gofiber/fiber/v3"

func (h *BaseHandler) DeleteUser(c fiber.Ctx) error {
	user_id := c.Params("user_id")
	if user_id == "" {
		return c.JSON(fiber.Map{
			"status": fiber.StatusForbidden,
			"error": "missing parameter: user_id",
		})
	}

	if err := h.repo.DeleteUser(user_id, c.Context()); err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}