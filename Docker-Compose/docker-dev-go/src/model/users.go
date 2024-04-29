package model

import (
	"docker-dev-go/database"
	"errors"
	"time"
)

// ユーザー情報の構造体
// usersテーブルのカラム名でjson出力するため、各項目に「`json:""`」を設定します。
type User struct {
	UserId    int       `json:"user_id"`
	Name      string    `json:"name"`
	OrgId     int       `json:"org_id"`
	Creater   string    `json:"creater"`
	CreatedAt time.Time `json:"created_at"`
	Updater   string    `json:"updater"`
	UpdatedAt time.Time `json:"updated_at"`
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
