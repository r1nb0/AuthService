package middleware

import (
	"AuthService/internal/constants"
	"AuthService/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"strings"
)

type AuthMiddleware struct {
	jwtUtil *utils.JWTUtil
}

func NewAuthMiddleware(jwtUtil *utils.JWTUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) Authentication(ctx fiber.Ctx) error {
	auth := ctx.Get(constants.AuthorizationHeaderKey)
	if auth == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "JWT token required",
		})
	}
	headerParts := strings.Split(auth, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "JWT token required",
		})
	}
	claimsMap, err := m.jwtUtil.GetClaims(headerParts[1])
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	ctx.Locals(constants.UserIdKey, claimsMap[constants.UserIdKey])
	ctx.Locals(constants.EmailKey, claimsMap[constants.EmailKey])
	ctx.Locals(constants.NicknameKey, claimsMap[constants.NicknameKey])
	ctx.Locals(constants.ExpireTimeKey, claimsMap[constants.ExpireTimeKey])
	return ctx.Next()
}
