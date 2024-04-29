package main

import (
	"docker-dev-go/config"
	"docker-dev-go/database"
	"docker-dev-go/router"
)

func main() {
	// configの初期化
	config.Init()

	// DBの初期化
	database.Init()
	defer database.Close()

	// routerの初期化
	router.Init()
}
