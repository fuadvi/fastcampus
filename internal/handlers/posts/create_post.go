package posts

import (
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateRequestPost
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")
	err := h.postSvc.CreatePost(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
