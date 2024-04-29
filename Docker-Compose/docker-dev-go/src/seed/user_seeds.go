package seed

import (
	"docker-dev-go/model"
	"time"

	"gorm.io/gorm"
)

func UserSeeds(db *gorm.DB) error {

	// １件目のデータ
	user1 := model.User{
		UserId:    1,
		Name:      "HirohisaMizuno",
		OrgId:     2,
		Creater:   "Hiro1",
		CreatedAt: time.Now(),
		Updater:   "Kiri1",
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&user1).Error; err != nil {
		return err
	}

	// ２件目のデータ
	user2 := model.User{
		UserId:    2,
		Name:      "KirikoMizuno",
		OrgId:     1,
		Creater:   "Hiro2",
		CreatedAt: time.Now(),
		Updater:   "Kiri2",
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&user2).Error; err != nil {
		return err
	}

	// 3件目のデータ
	user3 := model.User{
		UserId:    3,
		Name:      "YuzukiMizuno",
		OrgId:     3,
		Creater:   "Hiro3",
		CreatedAt: time.Now(),
		Updater:   "Kiri3",
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&user3).Error; err != nil {
		return err
	}

	return nil
}
