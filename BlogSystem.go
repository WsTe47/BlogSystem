package main

import (
	app "BlogSystem/internal/app/logic"
	"BlogSystem/internal/infra/data"
	"BlogSystem/internal/infra/impl"
	"BlogSystem/internal/infra/server"
	"fmt"
)

func main() {
	// 初始化数据库连接
	db, err := data.NewDbClient()
	if err != nil {
		fmt.Println(err)
	}

	// 初始化用户实现
	userImpl := impl.NewUserImpl(db)

	// 初始化用户应用
	userApp := app.NewUser(db, userImpl)

	// 初始化路由
	r := server.SetupRouter(userApp)

	// 启动HTTP服务器
	r.Run(":8080")
}
