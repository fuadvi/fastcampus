package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to created post comment to repository")
		return err
	}

	return nil
}
