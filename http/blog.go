package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) BlogIndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projecter.Get("blog", "index")
		if !found {
			s.NotFoundHandler()(c)
			return
		}

		s.respond(c, http.StatusOK, index)
	}
}

func (s *Server) BlogPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projecter.Get("blog", c.Param("slug"))
		if !found {
			s.NotFoundHandler()(c)
			return
		}

		s.respond(c, http.StatusOK, index)
	}
}
