package server

import (
	"ects/config"
	"ects/internal/routes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var Conf *config.Config

func Start() *iris.Application {
	server := iris.New()

	server.Use(iris.Gzip)
	server.Use(recover.New())
	server.Use(logger.New())

	// 注册路由
	routes.Register(server)

	return server
}