package models

import (
	"gorm.io/gorm"

	"time"
)

type Model struct {
	ID        uint64         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
type PlayerScore struct {
	Model
	// Id          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;column:id;default:uuid_generate_v4()"`
	MatchId     uint64 `json:"match_id"`
	PlayerId    uint64 `json:"player_id"`
	PlayerScore int32  `json:"player_score"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}
