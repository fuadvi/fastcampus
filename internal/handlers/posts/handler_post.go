package posts

import (
	"context"
	"github.com/fuadvi/fastcampus/internal/middleware"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreateRequestPost) error
	CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPost, error)
	GetPostByID(ctx context.Context, postID int) (*posts.GetPostResponse, error)
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
	route.POST("/comments/:postID", h.CreateComment)
	route.POST("/user-activity/:postID", h.UpsertUserActivity)
	route.GET("/", h.GetAllPost)
	route.GET("/:postID", h.GetPostByID)

}
