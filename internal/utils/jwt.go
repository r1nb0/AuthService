package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/r1nb0/UserService/internal/config"
	"github.com/r1nb0/UserService/internal/constants"
	"github.com/r1nb0/UserService/internal/domain"
	"time"
)

type JWTUtil struct {
	cfg *config.Config
}

func NewJWTUtil(cfg *config.Config) *JWTUtil {
	return &JWTUtil{
		cfg: cfg,
	}
}

func (u *JWTUtil) GenerateToken(user *domain.User) (string, error) {
	claimsMap := jwt.MapClaims{}
	claimsMap[constants.UserIdKey] = user.ID
	claimsMap[constants.NicknameKey] = user.Nickname
	claimsMap[constants.EmailKey] = user.Email
	claimsMap[constants.ExpireTimeKey] = time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)
	tokenStr, err := token.SignedString([]byte(u.cfg.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (u *JWTUtil) GetClaims(tokenStr string) (map[string]interface{}, error) {
	verifyToken, err := u.verifyToken(tokenStr)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (u *JWTUtil) verifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
