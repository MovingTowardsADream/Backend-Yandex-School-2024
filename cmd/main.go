package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	yandex_lavka "yandex-lavka"
	"yandex-lavka/pkg/handler"
	"yandex-lavka/pkg/repository"
	"yandex-lavka/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed reading config")
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed reading db password")
	}

	db, err := repository.NewPostgresDB(repository.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		log.Fatalf("Failed initialization database")
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	server := new(yandex_lavka.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitHandler()); err != nil {
		log.Fatal("Error server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
