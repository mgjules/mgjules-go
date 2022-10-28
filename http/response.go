package http

import "github.com/gin-gonic/gin"

func (s *Server) respond(c *gin.Context, code int, data []byte) {
	c.Header("Content-Encoding", "gzip")
	c.Data(code, "text/html", data)
}
