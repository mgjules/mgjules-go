package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		notFound, found := s.projection.Get("404")
		if !found {
			// If we get there, we are in big trouble lol
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the 404 page itself is not found. Awkward :s"})
			return
		}

		s.respond(c, http.StatusNotFound, notFound)
	}
}
