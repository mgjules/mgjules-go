package http

import (
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) RootHandler() http.HandlerFunc {
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
