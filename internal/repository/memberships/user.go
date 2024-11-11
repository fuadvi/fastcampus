package memberships

import (
	"context"
	"database/sql"
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/memberships"
)

func (r *Repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := "SELECT id, email, password, username, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ?"
	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username,
		&response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *Repository) GetUserById(ctx context.Context, userId int64) (*memberships.UserModel, error) {
	query := "SELECT id, email, password, username, created_at, updated_at, created_by, updated_by FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, userId)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username,
		&response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *Repository) CreateUser(ctx context.Context, user memberships.UserModel) error {
	query := "INSERT INTO users (email, password, created_at, updated_at, created_by, updated_by, username) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, &user.Email, user.Password, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy, &user.Username)

	if err != nil {
		return err
	}

	return nil
}
