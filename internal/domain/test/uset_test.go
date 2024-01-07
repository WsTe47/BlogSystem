package test

import (
	app "BlogSystem/internal/app/logic"
	"BlogSystem/internal/infra/data"
	"BlogSystem/internal/infra/impl"
	"context"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	db, err := data.NewDbClient()
	if err != nil {
		fmt.Println(err)
	}

	// 初始化用户实现
	userImpl := impl.NewUserImpl(db)

	// 初始化用户应用
	userApp := app.NewUser(db, userImpl)
	ctx := context.Background()
	id, err := userApp.Register(ctx, "username123", "password123", "13363631234")
	fmt.Println(id, err)
}
