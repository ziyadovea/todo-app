package main

import (
	"github.com/spf13/viper"
	"github.com/ziyadovea/todo-app/internal/app/apiserver"
	"github.com/ziyadovea/todo-app/internal/pkg/handler"
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/internal/pkg/service"
	"log"
)

func main() {

	if err := apiserver.InitConfig(); err != nil {
		log.Fatalf("error occurered while reading configuration file: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	server := apiserver.NewServer(viper.GetString("port"), h.InitRoutes())
	if err := server.Run(); err != nil {
		log.Fatalf("error occurered while running http server: %s", err.Error())
	}
}
