package model

import (
	"docker-dev-go/database"
	"errors"
	"time"
)

// ユーザー情報の構造体
// usersテーブルのカラム名でjson出力するため、各項目に「`json:""`」を設定します。
type User struct {
	user_id    int       `json:"user_id"`
	name       string    `json:"name"`
	org_id     int       `json:"org_id"`
	creater    string    `json:"creater"`
	created_at time.Time `json:"created_at"`
	updater    string    `json:"updater"`
	updated_at time.Time `json:"updated_at"`
}

// DBからusersテーブルの全レコードを取得
func GetUsersAll() (*[]User, error) {
	db := database.GetDB()

	var users *[]User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("users not registerd")
	}

	return users, nil
}
