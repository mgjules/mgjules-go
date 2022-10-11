package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) CVHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cv, found := s.projection.Get("cv", strings.ToLower(c.Param("section")))
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "404 Not Found"})
			return
		}

		s.respond(c, http.StatusOK, cv)
	}
}
