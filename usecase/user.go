package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/r1nb0/UserService/configs"
	"github.com/r1nb0/UserService/domain"
	"github.com/r1nb0/UserService/pkg/utils"
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

func (s *userService) SignIn(ctx context.Context, dto *domain.UserAuthDTO) (string, error) {
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

func (s *userService) SignUp(ctx context.Context, dto *domain.UserDTO) (int, error) {
	dto.Password = s.generatePasswordHash(dto.Password)
	return s.repo.Create(ctx, dto)
}

func (s *userService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) Update(ctx context.Context, id int, dto *domain.UserDTO) error {
	return s.repo.Update(ctx, id, dto)
}

func (s *userService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}
