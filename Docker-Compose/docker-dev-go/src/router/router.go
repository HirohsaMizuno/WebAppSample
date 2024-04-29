package router

import (
	"docker-dev-go/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	// routerの初期化
	router := gin.Default()

	// routerの設定
	router.GET("/", controller.Index)

	// APIの設定
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/config", controller.ConfigIndex)
	apiV1.GET("/users", controller.UsersIndex)

	// routerを起動
	router.Run(":3001")
}
