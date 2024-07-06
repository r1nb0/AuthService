package infra

import (
	"AuthService/internal/domain"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.UserDTO) (int, error) {
	var id int
	stmt, err := r.db.PrepareContext(
		ctx,
		"INSERT INTO users (first_name, last_name, nickname, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
	)
	if err != nil {
		return 0, err
	}
	defer closeStatement(stmt, err)
	row := stmt.QueryRowContext(
		ctx, user.FirstName, user.LastName,
		user.Nickname, user.Email,
		user.Password,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *userRepository) GetUser(ctx context.Context, dto *domain.UserAuthDTO) (*domain.User, error) {
	var user domain.User
	stmt, err := r.db.PrepareContext(
		ctx,
		"SELECT * FROM users WHERE nickname = $1 and password = $2",
	)
	if err != nil {
		return nil, err
	}
	defer closeStatement(stmt, err)
	row := stmt.QueryRowContext(ctx, dto.Nickname, dto.Password)
	if err = row.Scan(
		&user.ID, &user.FirstName,
		&user.LastName, &user.Nickname,
		&user.Email, &user.Password,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func closeStatement(stmt *sql.Stmt, err error) {
	errStmt := stmt.Close()
	if err != nil {
		if errStmt != nil {
			errors.Join(err, errStmt)
		}
	}
	err = errStmt
}
