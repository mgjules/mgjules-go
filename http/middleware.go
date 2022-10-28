package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		token := authHeader[len(BEARER_SCHEMA)+1:]
		if !s.auth.Validate(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Bearer token"})
			return
		}

		c.Next()
	}
}
