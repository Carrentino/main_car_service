package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func Connect(dsn string) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal(err)
	}
	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().In(loc)
		},
	}

	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal(err)
	}
}
