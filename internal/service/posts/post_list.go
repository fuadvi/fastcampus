package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *Service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPost, error) {
	limit := pageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 0 {
		pageSize = 1
	}
	offset := pageSize * (pageIndex - 1)

	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all post from database")
		return response, err
	}

	return response, nil
}
