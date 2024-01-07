package server

import (
	"BlogSystem/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

func NewHTTPServer(c *conf.Server, logger log.Logger) *gin.Engine {
	// 初始化Gin Engine
	r := gin.New()

	// 配置中间件
	r.Use(gin.Recovery())
	// r.Use(middleware1)
	// r.Use(middleware2)
	// ...

	// 配置路由
	// r.GET("/endpoint", handlerFunction)

	return r
}
