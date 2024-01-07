package impl

import (
	"BlogSystem/internal/domain/usecase"
	"BlogSystem/internal/infra/data"
	"context"
)

type UserImpl struct {
	Db *data.DbClient
}

func NewUserImpl(
	Db *data.DbClient,
) *UserImpl {
	return &UserImpl{
		Db: Db,
	}
}

func (UI *UserImpl) QueryUser(ctx context.Context, username string) (int, error) {
	var id int
	err := UI.Db.DB.Where("username = ?", username).First(&id).Error
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (UI *UserImpl) SaveUserRegister(ctx context.Context, user usecase.UserInformation) (int, error) {

	err := UI.Db.DB.Create(&user).Error
	if err != nil {
		return -1, err
	}

	return int(user.ID), nil
}
