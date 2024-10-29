package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/model/posts"
)

func (r *Repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id, post_title, post_content, post_hash_tag, created_at, created_by, updated_at, updated_by) values (?,?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHashTags, model.CreatedAt, model.CreatedBy, model.UpdatedAt, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}
