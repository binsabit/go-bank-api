package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/binsabit/go-bank-api/internal/data"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	accs, err := s.store.GetAccount()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, accs)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(GetFromParams("id", r))
	if err != nil {
		return err
	}
	acc, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println("creates")
	var reqAccount data.CreateAccountRequest

	err := ReadJSON(r, &reqAccount)

	if err != nil {
		return err
	}
	fmt.Printf("%+v", reqAccount)

	acc := data.NewAccount(reqAccount.FirstName, reqAccount.LastName)
	err = s.store.CreateAccount(acc)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(GetFromParams("id", r))
	if err != nil {
		return err
	}
	err = s.store.DeleteAccount(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
