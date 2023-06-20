package api

import "github.com/gorilla/mux"

func (s APIServer) router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount)).Methods("GET")
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleCreateAccount)).Methods("POST")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleDeleteAccount)).Methods("DELETE")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleTransferAccount)).Methods("PUT")

	return router
}
