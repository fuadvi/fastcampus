package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/configs"
	"github.com/fuadvi/fastcampus/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
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
