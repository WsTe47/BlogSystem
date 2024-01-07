package usecase

import (
	"time"
)

type Comment struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int `gorm:"primaryKey;autoIncrement"`
	ArticleID int
	AuthorID  int
	Content   string `gorm:"type:text"`
	IsDeleted bool
}
type Like struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int `gorm:"primaryKey;autoIncrement"`
	ArticleID int
	UserID    int
}

type Collect struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int `gorm:"primaryKey;autoIncrement"`
	ArticleID int
	UserID    int
}

// 表名为 "collect"
func (Collect) TableName() string {
	return "collect"
}

// 表名为 "likes"
func (Like) TableName() string {
	return "likes"
}

// 表名为 "comments"
func (Comment) TableName() string {
	return "comments"
}
