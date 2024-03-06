package score_module

import (
	score_controller "score-api/api/score/controller"
	score_service "score-api/api/score/service"
	"score-api/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterModule(app *fiber.App) {
	service := score_service.NewScoreService(database.DB)
	controller := score_controller.NewScoreController(*service)
	defPath := "api/v1"
	app.Get(defPath+"/scores", controller.GetScores)
	app.Post(defPath+"/scores", controller.CreateScore)
	app.Get(defPath+"/scores/:id", controller.GetScoreByPlayerId)
}
