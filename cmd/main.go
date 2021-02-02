package main

import (
	"log"
	"os"

	"github.com/Avepa/shortener/pkg/handler"
	"github.com/Avepa/shortener/pkg/repository"
	"github.com/Avepa/shortener/pkg/repository/mongodb"
	"github.com/Avepa/shortener/pkg/server"
	"github.com/Avepa/shortener/pkg/service"
)

func main() {
	DBcfg := &mongodb.Config{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBName:   os.Getenv("DATABASE_DBName"),
	}

	_, db, err := mongodb.NewMongoDB(DBcfg)
	if err != nil {
		log.Println(err)
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewSrvice(repos)
	handlers := handler.NewHandler(services)
	err = server.RunHTTPServer(
		os.Getenv("HTTTPSERVER_PORT"),
		handlers.Routes(),
	)

	if err != nil {
		log.Println(err)
		return
	}
}
