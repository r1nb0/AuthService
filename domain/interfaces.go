package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, dto *UserDTO) (int, error)
	GetByAuthData(ctx context.Context, dto *UserAuthDTO) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, id int, dto *UserDTO) error
}
