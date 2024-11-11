package memberships

import (
	"context"
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"github.com/fuadvi/fastcampus/pkg/jwt"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Service) ValidateRefreshToken(ctx context.Context, userId int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshTokenByUserId(ctx, userId, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Error getting refresh token")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token expired")
	}

	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipRepo.GetUserById(ctx, userId)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		return "", err
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}

	return token, nil
}
