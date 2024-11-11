package memberships

import (
	"context"
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"github.com/fuadvi/fastcampus/pkg/jwt"
	tokenUntil "github.com/fuadvi/fastcampus/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func (s *Service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", "", errors.New("email or password is invalid")
	}

	log.Printf(s.cfg.Service.SecretJwt)
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshTokenByUserId(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUntil.GenerateRefreshToken()

	now := time.Now()
	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserId:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    now.Add(10 * 24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to insert refresh token")
		return "", "", err
	}

	return token, refreshToken, nil
}
