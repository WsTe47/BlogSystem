package app

import "BlogSystem/internal/infra/data"

type User struct {
	Db *data.DbClient
}

func (user *User) Login(userName, passWord string) {

}

func Register() {

}
