package posts

import (
	"context"
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (s *Service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {

	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("error get user activity")
		return err
	}

	if userActivity == nil {

		if !request.IsLiked {
			log.Error().Err(err).Msg("belum pernah melakukan like post")
			return errors.New("and belum pernah like sebelumnya")
		}

		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity")
		return err
	}

	return nil

}
