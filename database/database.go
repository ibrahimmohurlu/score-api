package database

import (
	"log"
	"math/rand"
	"score-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.PlayerScore{})

	DB = Dbinstance{
		Db: db,
	}
}

func SeedDB() {
	scores := [30]models.PlayerScore{}
	for i := range scores {
		scores[i] = models.PlayerScore{
			MatchId:     rand.Uint64(),
			PlayerId:    rand.Uint64(),
			PlayerScore: int32(rand.Float64() * 1000),
		}
	}

	DB.Db.Create(&scores)

}
