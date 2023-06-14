package main

import (
	"log"

	"github.com/binsabit/go-bank-api/internal/api"
	"github.com/binsabit/go-bank-api/internal/data/postgres"
)

func main() {
	storage, err := postgres.NewPostgreStore()
	if err != nil {
		log.Fatal(err)
	}
	if err = storage.Init(); err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(":3000", storage)

	server.Run()
}
