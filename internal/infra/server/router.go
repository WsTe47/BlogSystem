package server

import (
	"BlogSystem/internal/api"
	"BlogSystem/internal/app/logic"
	"github.com/gin-gonic/gin"
)

func SetupRouter(user *logic.UserOperate) *gin.Engine {
	r := gin.Default()

	// 注册用户注册的路由和处理函数
	r.POST("/register", func(c *gin.Context) {
		api.UserRegisterHandler(c, user)
	})

	return r
}
