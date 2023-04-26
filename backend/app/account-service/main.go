package main

import (
	"job-site/pkg/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DbURL = os.Getenv("DB_URL")
var ServicePort = os.Getenv("SERVICE_PORT")

func main() {

	Db = database.ConnectDB(DbURL)

	app := fiber.New(fiber.Config{
		AppName: "account-service",
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(cache.New())
	app.Use(logger.New())

	api := app.Group("/api/v1")
	api.Group("/accounts")

	// TODO: add routes

	app.Listen(":" + ServicePort)
}
