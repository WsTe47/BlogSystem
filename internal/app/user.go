package app

import (
	"BlogSystem/internal/domain/usecase"
	"BlogSystem/internal/infra/data"
	"BlogSystem/internal/infra/impl"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
)

var (
	phonePattern    = `^\d{11}$`
	usernamePattern = `^[a-zA-Z0-9]{6,12}$`
	passwordPattern = `^[a-zA-Z0-9]{6,12}$`
)

type User struct {
	Db *data.DbClient
	UI *impl.UserImpl
}

func NewUser(
	Db *data.DbClient,
	UI *impl.UserImpl,
) *User {
	return &User{
		Db: Db,
		UI: UI,
	}
}

func (user *User) Login(ctx context.Context, db *data.DbClient, userName, passWord string) {

}

// Register 函数用于注册新用户
func (user *User) Register(ctx context.Context, username, password, phone string) (int, error) {
	// 校验手机号
	matched, _ := regexp.MatchString(phonePattern, phone)
	if !matched {
		return -1, fmt.Errorf("手机号格式不正确")
	}

	// 校验用户名
	matched, _ = regexp.MatchString(usernamePattern, username)
	if !matched {
		return -1, fmt.Errorf("用户名格式不正确")
	}

	// 校验密码
	matched, _ = regexp.MatchString(passwordPattern, password)
	if !matched {
		return -1, fmt.Errorf("密码格式不正确")
	}

	// 查询用户
	userID, err := user.UI.QueryUser(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建用户
			userInformation := usecase.UserInformation{
				Name:     username,
				Phone:    phone,
				Password: password,
				Role:     1,
				State:    1,
				// Likes:  可以根据实际需求进行处理
			}

			userID, err = user.UI.SaveUserRegister(ctx, userInformation)
			if err != nil {
				return -1, err
			}

		} else {
			return -1, err
		}
	} else {
		// 用户已存在
		return userID, fmt.Errorf("用户已存在")
	}

	return userID, nil
}
