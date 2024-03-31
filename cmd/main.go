package main

import (
	_ "github.com/lib/pq"
	"log"
	yandex_lavka "yandex-lavka"
	"yandex-lavka/pkg/handler"
	"yandex-lavka/pkg/repository"
	"yandex-lavka/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.ConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Failed initialization database")
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(yandex_lavka.Server)
	if err := server.Run("8080", handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}
