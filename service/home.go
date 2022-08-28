package service

import (
	"myapp/config"
	"myapp/graph/model"
	"time"
)

func GetHomeInfo() *model.Home {
	db := config.ConnectDatabase()
	defer db.Close()

	var result []*model.Home
	db.Table("home").Find(&result)

	time.Sleep(time.Second * 2)

	if len(result) > 0 {
		return result[0]
	}

	return nil
}

func UpdateHomeInfo(input model.EditedHome) (*model.Home, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	newSocialMediaInfo := model.Home{
		ID:          input.ID,
		Logo:        input.Logo,
		Profession:  input.Profession,
		Title:       input.Title,
		Description: input.Description,
	}

	err := db.Table("home").Where("id = ?", input.ID).Update(&newSocialMediaInfo).Error

	return &newSocialMediaInfo, err
}
