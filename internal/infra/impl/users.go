package impl

import (
	"BlogSystem/internal/domain/usecase"
	"context"
	"gorm.io/gorm"
)

var TableUser = "users"

type UserImpl struct {
	Db *gorm.DB
}

func NewUserImpl(
	Db *gorm.DB,
) *UserImpl {
	return &UserImpl{
		Db: Db,
	}
}

func (UI *UserImpl) QueryUser(ctx context.Context, username string) (int, error) {
	var userInformation usecase.UserInformation
	err := UI.Db.Table(TableUser).Where("name = ?", username).First(&userInformation).Error
	if err != nil {
		return -1, err
	}
	return int(userInformation.ID), nil
}
func (UI *UserImpl) SaveUserRegister(ctx context.Context, user usecase.UserInformation) (int, error) {

	err := UI.Db.Table(TableUser).Create(&user).Error
	if err != nil {
		return -1, err
	}

	return int(user.ID), nil
}
