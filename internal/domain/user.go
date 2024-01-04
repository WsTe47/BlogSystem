package domain

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(12);unique;not null"`
	Phone     string `gorm:"type:varchar(11);unique;not null"`
	Password  string `gorm:"type:varchar(12);not null"`
	Role      int
	State     int
	Likes     []int `gorm:"type:text"` // 需要序列化处理，具体开发时再定
}

// 表名为 "users"
func (User) TableName() string {
	return "users"
}
