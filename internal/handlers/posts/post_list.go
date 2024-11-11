package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()

	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	pageIndex, err := strconv.Atoi(pageIndexStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("invalid page index").Error()})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("invalid page size").Error()})
		return
	}

	response, err := h.postSvc.GetAllPost(ctx, pageSize, pageIndex)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}
