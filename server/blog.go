package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) BlogIndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projection.Get("blog", "index")
		if !found {
			c.Redirect(http.StatusFound, "/not-found")
			return
		}

		s.respond(c, http.StatusOK, index)
	}
}

func (s *Server) BlogPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		index, found := s.projection.Get("blog", c.Param("slug"))
		if !found {
			c.Redirect(http.StatusFound, "/not-found")
			return
		}

		s.respond(c, http.StatusOK, index)
	}
}
