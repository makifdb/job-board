package main

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

var (
	ProxyPort     = os.Getenv("PROXY_PORT")
	JobApiURL     = os.Getenv("JOB_API_URL")
	AccountApiURL = os.Getenv("ACCOUNT_API_URL")
)

func main() {

	log.Println("Service Gateway is running on port " + ProxyPort)

	app := fiber.New(fiber.Config{
		AppName:               "service-gateway",
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	app.All("/api/v1/*", func(c *fiber.Ctx) error {
		fullPath := strings.Split(string(c.Request().URI().RequestURI()), "/api/v1")[1]
		parentPath := strings.Split(fullPath, "/")[1]

		switch parentPath {
		case "accounts":
			url := strings.Join([]string{AccountApiURL, "/api/v1", fullPath}, "")

			if err := proxy.Do(c, url); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
		case "jobs":
			url := strings.Join([]string{JobApiURL, "/api/v1", fullPath}, "")

			if err := proxy.Do(c, url); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
		default:
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		}

		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	app.Listen(net.JoinHostPort("", ProxyPort))
}
