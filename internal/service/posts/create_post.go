package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

func (s *Service) CreatePost(ctx context.Context, userID int64, req posts.CreateRequestPost) error {
	postHashTags := strings.Join(req.PostHashTags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashTags: postHashTags,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.FormatInt(userID, 10),
		UpdatedBy:    strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error creating post")
		return err
	}

	return nil
}
