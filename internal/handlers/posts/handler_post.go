package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/middleware"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreateRequestPost) error
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, membershipSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)

}
