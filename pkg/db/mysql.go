package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
)

func _mysql() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn != "" {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.WithFields(log.Fields{
				"dsn": dsn,
			}).Error("Failed to connect to mysql: ", err)
			return
		}
		db = db.Debug()
		log.Infoln("connect to mysql")
	}
	return
}
