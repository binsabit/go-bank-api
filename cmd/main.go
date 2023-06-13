package main

import "github.com/binsabit/go-bank-api/internal/api"

func main() {
	server := api.NewServer(":3000")
	server.Run()
}
