package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v3"
	"github.com/r1nb0/UserService/constants"
	"github.com/r1nb0/UserService/internal/utils"
	"net/http"
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

func (m *AuthMiddleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader(constants.AuthorizationHeaderKey)
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{
				"success": false,
				"error":   "JWT token required",
			})
			return
		}
		headerParts := strings.Split(auth, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{
				"success": false,
				"error":   "JWT token required",
			})
			return
		}
		claimsMap, err := m.jwtUtil.GetClaims(headerParts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{
				"success": false,
				"error":   "JWT token required",
			})
			return
		}
		ctx.Set(constants.UserIdKey, claimsMap[constants.UserIdKey])
		ctx.Set(constants.EmailKey, claimsMap[constants.EmailKey])
		ctx.Set(constants.NicknameKey, claimsMap[constants.NicknameKey])
		ctx.Set(constants.ExpireTimeKey, claimsMap[constants.ExpireTimeKey])
		ctx.Next()
	}
}

// Authorization
// TODO impl
func (m *AuthMiddleware) Authorization(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
