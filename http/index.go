package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projecter.Get("index")
		if !found {
			s.NotFoundHandler()(c)
			return
		}

		s.respond(c, http.StatusOK, index)
	}
}
