package memberships

import (
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")
	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := memberships.RefreshTokenResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusCreated, response)
}
