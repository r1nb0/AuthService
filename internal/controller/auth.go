package controller

import (
	"AuthService/internal/domain"
	"AuthService/internal/usecase"
	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	uc usecase.UseCase
}

func NewAuthController(uc usecase.UseCase) *AuthController {
	return &AuthController{
		uc: uc,
	}
}

func (c *AuthController) SignIn(ctx fiber.Ctx) error {
	var input domain.UserAuthDTO
	if err := ctx.Bind().JSON(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	token, err := c.uc.SignIn(ctx.Context(), &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

func (c *AuthController) SignUp(ctx fiber.Ctx) error {
	var input domain.UserDTO
	if err := ctx.Bind().JSON(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	id, err := c.uc.SignUp(ctx.Context(), &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"id":      id,
	})
}
