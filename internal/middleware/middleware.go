package middleware

import (
	"errors"
	"github.com/fuadvi/fastcampus/internal/configs"
	"github.com/fuadvi/fastcampus/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJwt
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userId)
		c.Set("username", username)
		c.Next()
	}
}
