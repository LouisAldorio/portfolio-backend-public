package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetAudios(status []*int) []*model.Audio {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2, 1}
	if status != nil {
		defaultStatus = []int{}
		for _, v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.Audio
	db.Table("audio").Where("on_duty IN (?)", defaultStatus).Find(&result)

	return result
}

func CreateAudio(input model.NewAudio) *model.Audio {
	db := config.ConnectDatabase()
	defer db.Close()

	newAudio := model.Audio{
		Title:  input.Title,
		Author: input.Author,
		Src:    input.Src,
		Cover:  input.Cover,
		OnDuty: 2,
	}

	db.Table("audio").Create(&newAudio)

	return &newAudio
}

func UpdateAudioByID(input model.EditedAudio) (*model.Audio, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedAudio := model.Audio{
		ID:     input.ID,
		Src:    input.Src,
		Title:  input.Title,
		Author: input.Author,
		Cover:  input.Cover,
		OnDuty: input.OnDuty,
	}

	err := db.Table("audio").Where("id = ?", input.ID).Update(&editedAudio).Error

	return &editedAudio, err
}

func DeleteAudioByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("audio").Where("id = ?",id).Delete(&model.Audio{}).Error

	return id,err
}
