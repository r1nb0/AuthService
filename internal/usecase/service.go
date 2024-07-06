package usecase

import (
	"AuthService/internal/config"
	"AuthService/internal/domain"
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type authService struct {
	repo domain.UserRepository
	cfg  *config.Config
}

type jwtClaims struct {
	Nickname string `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthService(repo domain.UserRepository, cfg *config.Config) UseCase {
	return &authService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *authService) generateAuthToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		Nickname: user.Nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	tokenStr, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (s *authService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Password.Salt)))
}

func (s *authService) SignIn(ctx context.Context, dto *domain.UserAuthDTO) (string, error) {
	dto.Password = s.generatePasswordHash(dto.Password)
	user, err := s.repo.GetUser(ctx, dto)
	if err != nil {
		return "", err
	}
	token, err := s.generateAuthToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authService) SignUp(ctx context.Context, dto *domain.UserDTO) (int, error) {
	dto.Password = s.generatePasswordHash(dto.Password)
	return s.repo.CreateUser(ctx, dto)
}
