package posts

import (
	"errors"
	"github.com/fuadvi/fastcampus/internal/model/posts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("postID param tidak valid").Error()})
		return
	}

	userID := c.GetInt64("userID")
	err = h.postSvc.UpsertUserActivity(ctx, postID, userID, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
