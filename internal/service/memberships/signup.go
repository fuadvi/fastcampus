package memberships

import (
	"context"
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Username:  req.Username,
		Password:  string(pass),
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Username,
		UpdatedBy: req.Username,
	}

	err = s.membershipRepo.CreateUser(ctx, model)

	if err != nil {
		return err
	}

	return nil
}
