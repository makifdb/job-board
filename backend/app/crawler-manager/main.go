package main

import (
	"job-site/pkg/crawlers"
	"job-site/pkg/database"
	"job-site/pkg/entities"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DbURL = os.Getenv("DB_URL")
var ServicePort = os.Getenv("SERVICE_PORT")

var CrawlerList = map[string]bool{
	"himalayas": true,
}

func main() {

	Db = database.ConnectDB(DbURL)
	// silent mode
	Db.Logger = Db.Logger.LogMode(0)

	go Runner()

	app := fiber.New(fiber.Config{
		AppName:               "crawler-manager",
		DisableStartupMessage: true,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(logger.New())

	api := app.Group("/api/v1")
	c := api.Group("/crawlers")

	c.Get("/", GetCrawlers)
	c.Get("/:name", GetCrawler)
	c.Patch("/:name/enabled", EnableCrawler)
	c.Patch("/:name/disabled", DisableCrawler)

	app.Listen(":" + ServicePort)
}

func GetCrawlers(c *fiber.Ctx) error {
	// return crawler list in object array name, url and status

	st := CrawlerList["himalayas"]

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data": []fiber.Map{
			{
				"name":   "himalayas",
				"url":    "https://himalayas.app/",
				"status": st,
			},
		},
	})
}

func GetCrawler(c *fiber.Ctx) error {

	// get crawler name from url param
	name := c.Params("name")

	// return crawler object name, url and status
	switch name {
	case "himalayas":
		st := CrawlerList[name]
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data": fiber.Map{
				"name":   "himalayas",
				"url":    "https://himalayas.app/",
				"status": st,
			},
		})
	default:
		return c.JSON(fiber.Map{
			"status":  404,
			"message": "not found",
			"data":    nil,
		})
	}
}

func EnableCrawler(c *fiber.Ctx) error {

	// get crawler name from url param
	name := c.Params("name")

	// return crawler object name, url and status
	switch name {
	case "himalayas":

		// update crawler status
		CrawlerList[name] = true

		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data": fiber.Map{
				"name":   "himalayas",
				"url":    "https://himalayas.app/",
				"status": true,
			},
		})
	default:
		return c.JSON(fiber.Map{
			"status":  404,
			"message": "not found",
			"data":    nil,
		})
	}
}

func DisableCrawler(c *fiber.Ctx) error {

	// get crawler name from url param
	name := c.Params("name")

	// return crawler object name, url and status
	switch name {
	case "himalayas":

		// update crawler status
		CrawlerList[name] = false

		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data": fiber.Map{
				"name":   "himalayas",
				"url":    "https://himalayas.app/",
				"status": false,
			},
		})
	default:
		return c.JSON(fiber.Map{
			"status":  404,
			"message": "not found",
			"data":    nil,
		})
	}
}

func Runner() {

	log.Println("Crawler Manager Runner is running")

	limit := 1
	offset := 0

	for {
		switch {
		case CrawlerList["himalayas"]:
			// run crawler
			res, err := crawlers.GetHimalayaJobList(limit, offset)
			if err != nil {
				log.Println(err)
			}

			// save to db
			for _, v := range res.Jobs {

				// check if job already exist
				var j entities.Job
				if err := Db.Where("title = ? AND company_name = ?", v.Title, v.CompanyName).First(&j).Error; err == nil {
					log.Println("Job already exist, skipping")
					continue
				}

				// create job object
				job := entities.Job{
					Base: entities.Base{
						PublicTags: v.Categories,
					},
					Title:       v.Title,
					Description: v.Description,
					CompanyName: v.CompanyName,
					CompanyLogo: v.CompanyLogo,
					Source:      "himalayas",
				}

				// save job to database
				if err := Db.Create(&job).Error; err != nil {
					log.Println(err)
				}

				offset += limit
				log.Println("Job saved to database")
			}
		}
		time.Sleep(10 * time.Second)
	}

}
