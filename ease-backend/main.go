package main

import (
	"github.com/itkln/ease-backend/app"
	"github.com/itkln/ease-backend/app/controllers"
	"github.com/itkln/ease-backend/app/handlers"
	"github.com/itkln/ease-backend/app/repository"
	"log"
)

func main() {
	config := app.SetupConfig()
	log.Println(config.AppConfig)
	app.SetupServer(config.AppConfig.Port)
	db := app.SetupDb(config.AppConfig.Dsn)
	userRepo := repository.NewUserRepository(db)
	authCtrl := controllers.NewAuthServiceImpl(userRepo)
	handlers.NewHandler(authCtrl)
}
