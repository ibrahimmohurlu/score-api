package database

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"score-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		//os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
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
