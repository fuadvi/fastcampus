package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *Service) GetPostByID(ctx context.Context, postID int) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error getting post by id")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error getting like count by post id")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error getting get post comments by post id")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: postDetail,
		LikeCount:  likeCount,
		Comments:   comments,
	}, nil
}
