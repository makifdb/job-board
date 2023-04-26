package main

import (
	"job-site/pkg/database"
	"job-site/pkg/dtos"
	"job-site/pkg/entities"
	"log"
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
		AppName: "job-service",
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(cache.New())
	app.Use(logger.New())

	api := app.Group("/api/v1")
	j := api.Group("/jobs")

	j.Get("/:id", GetJob)
	j.Post("/", CreateJob)

	app.Listen(":" + ServicePort)
}

func GetJob(c *fiber.Ctx) error {
	// get id from url
	id := c.Params("id")

	var job entities.Job
	// get job from db
	if err := Db.Where("public_id = ?", id).First(&job).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Job not found", "data": nil})
	}

	resp := dtos.JobDTO{
		ID:          job.PublicID,
		Title:       job.Title,
		Description: job.Description,
		Location:    job.Location,
		CompanyName: job.CompanyName,
		CompanyLogo: job.CompanyLogo,
		Tags:        job.PublicTags,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Job found", "data": resp})
}

func CreateJob(c *fiber.Ctx) error {

	var input dtos.JobCreateDTO
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	job := entities.Job{
		Base: entities.Base{
			PublicTags: input.Tags,
		},
		Title:       input.Title,
		Description: input.Description,
		Location:    input.Location,
		CompanyName: input.CompanyName,
		CompanyLogo: input.CompanyLogo,
	}

	// save job to database
	if err := Db.Create(&job).Error; err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": "error", "message": "Job not created", "data": nil})
	}

	resp := dtos.JobDTO{
		ID:          job.PublicID,
		Title:       job.Title,
		Description: job.Description,
		Location:    job.Location,
		CompanyName: job.CompanyName,
		CompanyLogo: job.CompanyLogo,
		Tags:        job.PublicTags,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Job created", "data": resp})
}
