package memberships

import (
	"github.com/fuadvi/fastcampus/internal/model/memberships"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accassToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := memberships.LoginResponse{
		AccessToken: accassToken,
	}

	c.JSON(http.StatusCreated, response)
}
