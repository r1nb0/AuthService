package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/r1nb0/UserService/configs"
	"github.com/r1nb0/UserService/internal/domain"
	"github.com/r1nb0/UserService/internal/utils"
)

type userService struct {
	repo    domain.UserRepository
	jwtUtil *utils.JWTUtil
	cfg     *configs.Config
}

func NewUserService(repo domain.UserRepository, jwtUtil *utils.JWTUtil, cfg *configs.Config) UserUseCase {
	return &userService{
		repo:    repo,
		jwtUtil: jwtUtil,
		cfg:     cfg,
	}
}

func (s *userService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Password.Salt)))
}

func (s *userService) SignIn(ctx context.Context, dto *domain.AuthenticateUser) (string, error) {
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

func (s *userService) SignUp(ctx context.Context, dto *domain.CreateUser) (int, error) {
	dto.Password = s.generatePasswordHash(dto.Password)
	return s.repo.Create(ctx, dto)
}

func (s *userService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) Update(ctx context.Context, id int, dto *domain.UpdateUserGeneralInfo) error {
	if dto.IsValid() {
		return s.repo.Update(ctx, id, dto)
	}
	return errors.New("at least 1 field must be provided to update data")
}

func (s *userService) UpdatePassword(ctx context.Context, id int, password string) error {
	hashPassword := s.generatePasswordHash(password)
	return s.repo.UpdatePassword(ctx, id, hashPassword)
}

func (s *userService) UpdateEmail(ctx context.Context, id int, email string) error {
	return s.repo.UpdateEmail(ctx, id, email)
}

func (s *userService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}
