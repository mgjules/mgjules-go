package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RefetchDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.projection.FetchData()

		c.JSON(http.StatusOK, gin.H{"status": "data fetched successfully"})
	}
}

func (s *Server) RebuildProjectionsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.projection.BuildProjections()

		c.JSON(http.StatusOK, gin.H{"status": "projections rebuilt successfully"})
	}
}
