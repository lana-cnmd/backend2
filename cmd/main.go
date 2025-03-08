package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	server "github.com/lana-cnmd/backend2"
	"github.com/lana-cnmd/backend2/pkg/handlers"
	"github.com/lana-cnmd/backend2/pkg/repository"
	"github.com/lana-cnmd/backend2/pkg/repository/postgres"
	"github.com/lana-cnmd/backend2/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// @title Shop API Documentation
// @version 1.0
// @description This is the API documentation for the Shop application.
// @host localhost:8000
// @BasePath /api/v1
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env var: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	log.Println("Hi lapushka")

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	server := new(server.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("errror ocured in db conneciton close: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
