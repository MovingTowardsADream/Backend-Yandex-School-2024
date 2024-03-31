package main

import (
	"log"
	yandex_lavka "yandex-lavka"
	"yandex-lavka/pkg/handler"
	"yandex-lavka/pkg/repository"
	"yandex-lavka/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(yandex_lavka.Server)
	if err := server.Run("8080", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
