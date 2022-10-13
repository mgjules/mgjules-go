package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) BlogIndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projection.Get("blog", "index")
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "404 Not Found"})
			return
		}

		s.respond(c, http.StatusOK, index)
	}
}
