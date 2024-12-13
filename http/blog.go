package http

import (
	"net/http"
)

func (s *Server) BlogIndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index, found := s.projecter.Get("blog", "index")
		if !found {
			s.NotFoundHandler()(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(index)
	}
}

func (s *Server) BlogPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post, found := s.projecter.Get("blog", r.PathValue("slug"))
		if !found {
			s.NotFoundHandler()(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(post)
	}
}
