package memberships

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/middleware"
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userId int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/sign-in", h.SignIn)

	route.Use(middleware.AuthRefreshMiddleware())
	route.POST("/refresh", h.Refresh)
}
