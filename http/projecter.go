package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RebuildProjectionsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		go s.projecter.Build()

		c.JSON(http.StatusOK, gin.H{"status": "projections rebuilt successfully"})
	}
}
