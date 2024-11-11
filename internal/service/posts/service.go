package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/configs"
	"github.com/fuadvi/fastcampus/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPost, error)
	GetPostByID(ctx context.Context, id int) (posts.Post, error)

	CountLikeByPostID(ctx context.Context, postID int) (int, error)
	GetCommentByPostID(ctx context.Context, postID int) ([]posts.Comment, error)
}

type Service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *Service {
	return &Service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
