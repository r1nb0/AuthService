package infra

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/r1nb0/UserService/internal/domain"
	"github.com/r1nb0/UserService/pkg/logging"
	"strings"
)

type userRepository struct {
	db     *sqlx.DB
	logger logging.Logger
}

func NewUserRepository(db *sqlx.DB, logger logging.Logger) domain.UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) Create(ctx context.Context, dto *domain.CreateUser) (int, error) {
	var id int
	query := "INSERT INTO users (first_name, last_name, nickname, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	stmt, err := r.db.PrepareContext(
		ctx,
		query,
	)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return 0, err
	}
	defer closeStmt(stmt, err)
	row := stmt.QueryRowContext(
		ctx, dto.FirstName, dto.LastName,
		dto.Nickname, dto.Email,
		dto.Password,
	)
	if err := row.Scan(&id); err != nil {
		r.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return 0, err
	}
	return id, nil
}

func (r *userRepository) GetByAuthData(ctx context.Context, dto *domain.AuthenticateUser) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE nickname = $1 and password = $2"
	stmt, err := r.db.PrepareContext(
		ctx,
		query,
	)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	defer closeStmt(stmt, err)
	row := stmt.QueryRowContext(ctx, dto.Nickname, dto.Password)
	if err = row.Scan(
		&user.ID, &user.FirstName,
		&user.LastName, &user.Nickname,
		&user.Email, &user.Password,
	); err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	users := make([]*domain.User, 0)
	query := "SELECT * FROM users"
	stmt, err := r.db.PrepareContext(
		ctx,
		query,
	)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	defer closeStmt(stmt, err)
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.ID, &user.FirstName,
			&user.LastName, &user.Nickname,
			&user.Email, &user.Password,
		); err != nil {
			r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE id = $1"
	stmt, err := r.db.PrepareContext(
		ctx,
		query,
	)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	defer closeStmt(stmt, err)
	row := stmt.QueryRowContext(ctx, id)
	if err := row.Scan(
		&user.ID, &user.FirstName,
		&user.LastName, &user.Nickname,
		&user.Email, &user.Password,
	); err != nil {
		r.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, id int, dto *domain.UpdateUserGeneralInfo) error {
	var (
		placeholders []string
		args         []interface{}
		cnt          = 1
	)
	if dto.FirstName != "" {
		placeholders = append(placeholders, fmt.Sprintf("first_name = $%d", cnt))
		args = append(args, dto.FirstName)
		cnt++
	}
	if dto.LastName != "" {
		placeholders = append(placeholders, fmt.Sprintf("last_name = $%d", cnt))
		args = append(args, dto.LastName)
		cnt++
	}
	if dto.Nickname != "" {
		placeholders = append(placeholders, fmt.Sprintf("nickname = $%d", cnt))
		args = append(args, dto.Nickname)
		cnt++
	}
	args = append(args, id)
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(placeholders, ","), cnt)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	if _, err := stmt.ExecContext(ctx, args...); err != nil {
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	return nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, id int, password string) error {
	query := "UPDATE users SET password = $1 WHERE id = $2"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	defer closeStmt(stmt, err)
	if _, err := stmt.ExecContext(ctx, password, id); err != nil {
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	return nil
}

func (r *userRepository) UpdateEmail(ctx context.Context, id int, email string) error {
	query := "UPDATE users SET email = $1 where id = $2"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	defer closeStmt(stmt, err)
	if _, err := stmt.ExecContext(ctx, email, id); err != nil {
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	return nil
}

func closeStmt(stmt *sql.Stmt, err error) {
	errStmt := stmt.Close()
	if err != nil {
		if errStmt != nil {
			errors.Join(err, errStmt)
		}
	}
	err = errStmt
}
