package domain

import (
	"time"
)

type Article struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Bid       int `gorm:"primaryKey;autoIncrement"`
	AuthorID  int
	Title     string
	Content   string `gorm:"type:longtext"`
	IsDeleted bool
	ShowState int
}

// 表名为 "article"
func (Article) TableName() string {
	return "article"
}
