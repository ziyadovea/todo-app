package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ziyadovea/todo-app/internal/app/apiserver"
	"github.com/ziyadovea/todo-app/internal/app/apiserver/handler"
	repository2 "github.com/ziyadovea/todo-app/internal/app/repository"
	"github.com/ziyadovea/todo-app/internal/app/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := apiserver.InitConfig(); err != nil {
		logrus.Fatalf("error occured while reading configuration file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occured while reading .env file: %s", err.Error())
	}

	db, err := repository2.NewPostgresDB(&repository2.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBName:   viper.GetString("db.db_name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error occured while connecting to the postgres database: %s", err.Error())
	}

	repo := repository2.NewRepository(db)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	server := apiserver.NewServer(viper.GetString("port"), h.InitRoutes())
	go func() {
		if err := server.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Todo app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on database connection closing: %s", err.Error())
	}

	logrus.Print("Todo app finished")
}
