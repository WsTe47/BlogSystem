package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "climber47:abcd4321@tcp(123.60.48.133:3306)/blogsystem?parseTime=True"

func NewDbClient() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
