package usecase

import (
	"context"
	"time"
)

type UserInformation struct {
	ID        uint      `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:createdat"`
	UpdatedAt time.Time `gorm:"column:updatedat"`
	Name      string    `gorm:"column:name"`
	Phone     string    `gorm:"column:phone"`
	Password  string    `gorm:"column:password"`
	Role      int       `gorm:"column:role"`
	State     int       `gorm:"column:state"`
	Likes     string    `gorm:"column:likes"` //这里该存切片，但现在写会有问题。后续处理。 需要序列化处理，具体开发时再定
}

// 表名为 "users"
func (UserInformation) TableName() string {
	return "users"
}

type User interface {
	QueryUser(ctx context.Context, username string) (int, error)
	SaveUserRegister(ctx context.Context, user UserInformation) (int, error)
}
