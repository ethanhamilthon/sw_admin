package main

import (
	"log"
	"templtest/config"
	"templtest/internal/services"
	"templtest/internal/storage/sqlite"
	"templtest/internal/transport/web"
)

func main() {
	config.Initial()
	repo, err := sqlite.New()
	if err != nil {
		log.Fatal("db connection failed")
	}

	service := services.New(repo)

	server := web.New(service)
	server.Serve()
}
