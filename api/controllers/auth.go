package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/UserService/domain"
	"github.com/r1nb0/UserService/usecase"
	"net/http"
)

type AuthController struct {
	uc usecase.AuthUseCase
}

func NewAuthController(uc usecase.AuthUseCase) *AuthController {
	return &AuthController{
		uc: uc,
	}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var input domain.UserAuthDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	token, err := c.uc.SignIn(ctx, &input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var input domain.UserDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	id, err := c.uc.SignUp(ctx, &input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      id,
	})
}
