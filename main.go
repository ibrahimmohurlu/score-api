package main

import (
	score_module "score-api/api/score/module"
	"score-api/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	database.ConnectDb()

	app.Get("/", monitor.New(monitor.Config{Refresh: time.Second}))

	score_module.RegisterModule(app)

	app.Listen(":3000")
}
