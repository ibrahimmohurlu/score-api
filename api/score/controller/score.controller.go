package score_controller

import (
	"fmt"
	score_service "score-api/api/score/service"

	"github.com/gofiber/fiber/v2"
)

type ScoreController struct {
	scoreService score_service.ScoreService
}

func NewScoreController(scoreService score_service.ScoreService) *ScoreController {
	return &ScoreController{
		scoreService: scoreService,
	}
}
func (controller ScoreController) GetScores(c *fiber.Ctx) error {
	data := controller.scoreService.GetAllScores()
	res := fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	}

	return c.Status(200).JSON(res)
}
func (controller ScoreController) GetScoreByPlayerId(c *fiber.Ctx) error {
	param := struct {
		ID uint64 `params:"id"`
	}{}
	c.ParamsParser(&param)
	data := controller.scoreService.GetScoresByPlayerId(uint64(param.ID))
	res := fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	}
	return c.Status(200).JSON(res)
}
func (controller ScoreController) CreateScore(c *fiber.Ctx) error {
	bodyData := &score_service.CreateScoreDto{}
	err := c.BodyParser(&bodyData)

	if bodyData.PlayerId == 0 || bodyData.MatchId == 0 {
		return c.Status(400).JSON("body param missing")
	}
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(404)
	}
	fmt.Println(bodyData)
	data, err := controller.scoreService.CreateNewScore(bodyData)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(404)
	}
	return c.Status(201).JSON(data)
}
