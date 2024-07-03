package db

import (
	"errors"
	"gorm.io/gorm"
)

var db *gorm.DB

func Getdb() (d *gorm.DB, err error) {
	_mysql()
	_sqlite()
	_postgres()
	if db != nil {
		return db, nil
	}
	return nil, errors.New("need DB!")
}
