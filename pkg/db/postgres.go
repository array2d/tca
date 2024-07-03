package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)

func _postgres() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn != "" {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.WithFields(log.Fields{
				"dsn": dsn,
			}).Error("Failed to connect to postgres: ", err)
			return
		}
		db = db.Debug()
	}
}
