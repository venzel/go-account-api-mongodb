package handlers

import (
	"encoding/json"
	"mongoapi/internal/domain/accounts"

	"github.com/gofiber/fiber/v2"
)

type AccountsHandler struct {
	UseCases accounts.UseCaseInterface
}

func (h *AccountsHandler) Create(ctx *fiber.Ctx) error {
	body := []byte(ctx.Body())

	var dto accounts.InputAccountDto

	json.Unmarshal(body, &dto)

	account, err := h.UseCases.Create(&dto)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Account created successfully",
		"data":    account,
	})
}
