package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itkln/ease-backend/app/routes"
	"github.com/itkln/ease-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func SetupConfig() *config.Config {
	return config.InitConfig()
}

func SetupServer(port int) {
	serv := fiber.New()
	routes.SetupRoutes(serv)
	serv.Listen(":" + strconv.Itoa(port))
}

func SetupDb(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("The connection to db failed")
	}
	return db
}
