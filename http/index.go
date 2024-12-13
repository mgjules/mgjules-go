package http

import (
	"net/http"
)

func (s *Server) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index, found := s.projecter.Get("index")
		if !found {
			s.NotFoundHandler()(w, r)
			return
		}

		renderHTML(w, index)
	}
}
