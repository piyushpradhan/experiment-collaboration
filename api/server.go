package api

import (
	"collaboration/storage"
	"collaboration/util"
	"encoding/json"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	allowedOrigins := []string{
		"https://experiment.piyushpradhan.space",
		"http://localhost",
		"https://showoff-frontend.vercel.app",
	}

	corsHandler := CORSMiddleware(allowedOrigins)

	http.HandleFunc("/user", corsHandler(http.HandlerFunc(s.handleGetUserById)).ServeHTTP)
	http.HandleFunc("/user/id", corsHandler(http.HandlerFunc(s.handleDeleteUserById)).ServeHTTP)

	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	user, _ := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	user, _ := s.store.Get(10)

	_ = util.Round2Dec(10.3933)

	json.NewEncoder(w).Encode(user)
}

