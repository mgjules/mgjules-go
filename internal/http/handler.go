package http

import (
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		b, found := s.projecter.Get(path)
		if !found {
			b, found = s.projecter.Get(filepath.Join(path, "index"))
			if !found {
				http.ServeFileFS(w, r, s.static, path)
				return
			}
		}

		renderHTML(w, b)
	}
}

func (s *Server) notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notFound, found := s.projecter.Get("404")
		if !found {
			// If we get there, we are in big trouble lol.
			renderJSON(w, http.StatusInternalServerError, map[string]string{"error": "oh no! we don't have a 404 page."})
			return
		}

		renderHTML(w, notFound)
	}
}
