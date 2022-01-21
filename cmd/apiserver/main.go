package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ziyadovea/todo-app/internal/app/apiserver"
	"github.com/ziyadovea/todo-app/internal/pkg/handler"
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/internal/pkg/service"
	"os"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := apiserver.InitConfig(); err != nil {
		logrus.Fatalf("error occurered while reading configuration file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occurered while reading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBName:   viper.GetString("db.db_name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error occurered while connecting to the postgres database: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	server := apiserver.NewServer(viper.GetString("port"), h.InitRoutes())
	if err := server.Run(); err != nil {
		logrus.Fatalf("error occurered while running http server: %s", err.Error())
	}
}
