package main

import (
	"github.com/spf13/viper"
	"log"
	"main_car_service"
	"main_car_service/database"
	"main_car_service/internal/handler"
	"main_car_service/internal/models"
	"main_car_service/internal/repository"
	"main_car_service/internal/service"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading config config, %s", err.Error())
	}

	dsn := os.Getenv("DSN")
	database.Connect(dsn)

	if err := models.Migrate(database.DB); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	repos := repository.NewRepository(database.DB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(main_car_service.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	return viper.ReadInConfig()
}
