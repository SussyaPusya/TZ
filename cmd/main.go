package main

import (
	"github.com/SussyaPusya/TZ/internal/repository"
	"github.com/SussyaPusya/TZ/internal/transport/handlers"
	"github.com/SussyaPusya/TZ/internal/transport/router"
)

func main() {

	repo := repository.New()

	handl := handlers.New(repo)

	router := router.New(handl)

	router.Run()
}
