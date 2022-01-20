package main

import (
	"github.com/ziyadovea/todo-app/internal/app/apiserver"
	"github.com/ziyadovea/todo-app/internal/pkg/handler"
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/internal/pkg/service"
	"log"
)

func main() {

	repo := repository.NewRepository()
	services := service.NewService(repo)
	h := handler.NewHandler(services)
	
	server := apiserver.NewServer(":8080", h.InitRoutes())
	if err := server.Run(); err != nil {
		log.Fatalf("error occurered while running http server: %s", err.Error())
	}
}
