package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "user:password@tcp(127.0.0.1:3306)/dbname?parseTime=True"

type DbClient struct {
	DB *gorm.DB
}

func NewDbClient() (*DbClient, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DbClient{}, err
	}

	return &DbClient{DB: db}, nil
}
