package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RefetchDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		go s.fetcher.Fetch()

		c.JSON(http.StatusOK, gin.H{"status": "data fetched successfully"})
	}
}
