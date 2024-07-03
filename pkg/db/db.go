package db

import (
	"errors"
	"gorm.io/gorm"
)

var db *gorm.DB

func Getdb() (db *gorm.DB, err error) {
	if db != nil {
		return db, nil
	}
	return nil, errors.New("need DB!")
}
