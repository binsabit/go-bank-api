package api

import (
	"log"
	"net/http"

	"github.com/binsabit/go-bank-api/internal/data"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      data.Storage
}

func NewServer(listenAddr string, store data.Storage) *APIServer {
	return &APIServer{listenAddr: listenAddr, store: store}
}

func (s *APIServer) Run() {

	log.Println("server is running on port ", s.listenAddr)

	srv := &http.Server{
		Addr:    s.listenAddr,
		Handler: s.router(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
