package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projection.Get("index")
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "404 Not Found"})
			return
		}

		c.Data(http.StatusOK, "text/html", index)
	}
}
