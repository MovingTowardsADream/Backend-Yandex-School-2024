package main

import (
	"log"
	yandex_lavka "yandex-lavka"
	"yandex-lavka/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(yandex_lavka.Server)
	if err := server.Run("8080", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
