package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("postID param tidak valid").Error()})
		return
	}

	data, err := h.postSvc.GetPostByID(ctx, int(postID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
