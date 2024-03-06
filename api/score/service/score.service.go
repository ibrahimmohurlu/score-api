package score_service

import (
	"fmt"
	"score-api/database"
	"score-api/models"
)

type ScoreService struct {
	db database.Dbinstance
}

func NewScoreService(db database.Dbinstance) *ScoreService {
	return &ScoreService{
		db: db,
	}
}

func (s ScoreService) GetAllScores() []models.PlayerScore {
	scores := []models.PlayerScore{}
	result := database.DB.Db.Find(&scores)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return scores
}

func (s ScoreService) GetScoresByPlayerId(playerId uint64) []models.PlayerScore {
	scores := []models.PlayerScore{}
	result := database.DB.Db.Where("player_id= ?", playerId).Find(&scores)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return scores
}

type CreateScoreDto struct {
	MatchId      uint64 `json:"match_id"`
	PlayerId     uint64 `json:"player_id"`
	KillCount    uint8  `json:"kill_count"`
	SurvivalTime uint16 `json:"survival_time"`
}

func (s ScoreService) CreateNewScore(createScoreDto *CreateScoreDto) (*models.PlayerScore, error) {
	score := int32(createScoreDto.KillCount*3) + int32(createScoreDto.SurvivalTime*2)
	playerScore := models.PlayerScore{
		MatchId:     createScoreDto.MatchId,
		PlayerId:    createScoreDto.PlayerId,
		PlayerScore: score,
	}
	result := s.db.Db.Create(&playerScore)
	if result.Error != nil {
		return nil, result.Error
	}

	return &playerScore, nil
}
