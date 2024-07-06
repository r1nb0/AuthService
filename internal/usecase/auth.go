package usecase

import (
	"AuthService/internal/config"
	"AuthService/internal/domain"
	"AuthService/pkg/utils"
	"context"
	"crypto/sha1"
	"fmt"
)

type authService struct {
	repo    domain.UserRepository
	jwtUtil *utils.JWTUtil
	cfg     *config.Config
}

func NewAuthService(repo domain.UserRepository, jwtUtil *utils.JWTUtil, cfg *config.Config) AuthUseCase {
	return &authService{
		repo:    repo,
		jwtUtil: jwtUtil,
		cfg:     cfg,
	}
}

func (s *authService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Password.Salt)))
}

func (s *authService) SignIn(ctx context.Context, dto *domain.UserAuthDTO) (string, error) {
	dto.Password = s.generatePasswordHash(dto.Password)
	user, err := s.repo.GetByAuthData(ctx, dto)
	if err != nil {
		return "", err
	}
	token, err := s.jwtUtil.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authService) SignUp(ctx context.Context, dto *domain.UserDTO) (int, error) {
	dto.Password = s.generatePasswordHash(dto.Password)
	return s.repo.Create(ctx, dto)
}
