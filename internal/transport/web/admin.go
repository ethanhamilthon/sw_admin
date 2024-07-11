package web

import (
	"log"
	"net/http"
	"templtest/views/pages"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.AllUsers()
	if err != nil {
		log.Println(err.Error())
	}
	pages.Index(users).Render(r.Context(), w)
}
