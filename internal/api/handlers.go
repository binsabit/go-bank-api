package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/binsabit/go-bank-api/internal/data"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	case "PUT":
		return s.handleTransferAccount(w, r)
	}
	return fmt.Errorf("Method Not Allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("creates")
	var reqAccount data.CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&reqAccount)
	if err != nil {
		return err
	}

	acc := data.NewAccount(reqAccount.FirstName, reqAccount.LastName)
	err = s.store.CreateAccount(acc)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
