package main

import (
	"docker-dev-go/config"
	"docker-dev-go/database"
	"docker-dev-go/seed"
	"fmt"
)

func main() {
	// Configの初期化
	config.Init()

	// DBの初期化
	database.Init()
	defer database.Close()

	// DBの取得
	db := database.GetDB()

	// seedの実行
	err := seed.UserSeeds(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Execution user_seeds.go !!")
}
