package controller

import (
	"AuthService/internal/constants"
	"AuthService/internal/domain"
	"AuthService/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type UserController struct {
	uc usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{
		uc: uc,
	}
}

func (c *UserController) GetAll(ctx fiber.Ctx) error {
	users, err := c.uc.GetAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func (c *UserController) GetByID(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	user, err := c.uc.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func (c *UserController) Update(ctx fiber.Ctx) error {
	authID := ctx.Locals(constants.UserIdKey).(int)
	var input domain.UserDTO
	if err := ctx.Bind().JSON(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	if err := c.uc.Update(ctx.Context(), authID, &input); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
