package memberships

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/configs"
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"time"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user memberships.UserModel) error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshTokenByUserId(ctx context.Context, userId int64, now time.Time) (*memberships.RefreshTokenModel, error)
	GetUserById(ctx context.Context, userId int64) (*memberships.UserModel, error)
}

type Service struct {
	cfg            *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *Service {
	return &Service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
