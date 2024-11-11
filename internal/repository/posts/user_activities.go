package posts

import (
	"context"
	"database/sql"
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/posts"
)

func (r *Repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {

	query := `SELECT id, post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by FROM user_activities where post_id = ? AND user_id = ?`
	var response posts.UserActivityModel

	row := r.db.QueryRowContext(ctx, query, model.PostID, model.UserID)

	err := row.Scan(&response.ID, &response.PostID, &response.UserID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &response, nil
}

func (r *Repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES (?,?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activities set is_liked = ?, updated_at = ?, updated_by = ? where post_id = ? AND user_id = ?`

	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.UpdatedAt, model.UpdatedBy, model.PostID, model.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CountLikeByPostID(ctx context.Context, postID int) (int, error) {

	query := `SELECT count(id) FROM user_activities where post_id = ? AND is_liked = true`
	var response int

	row := r.db.QueryRowContext(ctx, query, postID)

	err := row.Scan(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
