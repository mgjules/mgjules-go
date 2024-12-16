package http

import (
	"net/http"
)

func (s *Server) NotFoundHandler() http.HandlerFunc {
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
