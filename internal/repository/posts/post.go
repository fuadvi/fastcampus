package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"strings"
)

func (r *Repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id, post_title, post_content, post_hash_tag, created_at, created_by, updated_at, updated_by) values (?,?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHashTags, model.CreatedAt, model.CreatedBy, model.UpdatedAt, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPost, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hash_tag FROM posts p join users u on p.user_id = u.id ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	var response posts.GetAllPost

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}

	data := make([]posts.Post, 0)
	for rows.Next() {
		var (
			model    posts.PostModel
			username string
		)

		err := rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashTags)

		if err != nil {
			return response, err
		}

		data = append(data, posts.Post{
			ID:           model.ID,
			UserID:       model.UserID,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashTags: strings.Split(model.PostHashTags, ","),
		})
	}

	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	return response, nil
}

func (r *Repository) GetPostByID(ctx context.Context, id int) (posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hash_tag, ua.is_liked 
				FROM posts p join users u on p.user_id = u.id 
				JOIN user_activities ua on p.id = ua.post_id
				WHERE p.id = ?`

	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)

	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashTags, &isLiked)

	if err != nil {
		return posts.Post{}, err
	}

	return posts.Post{
		ID:           model.ID,
		UserID:       model.UserID,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashTags: strings.Split(model.PostHashTags, ","),
		IsLiked:      isLiked,
	}, nil
}
