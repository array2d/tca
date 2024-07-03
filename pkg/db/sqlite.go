package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func init() {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath != "" {
		var err error
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"dbPath": dbPath,
			}).Error("Failed to connect to sqlite: ", err)
			return
		}
		db = db.Debug()
	}
}
