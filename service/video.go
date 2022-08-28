package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetVideos(status []*int) []*model.Video {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2,1}
	if status != nil {
		defaultStatus = []int{}
		for _,v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.Video
	db.Table("video").Where("on_duty IN (?)",defaultStatus).Find(&result)

	return result
}

func CreateVideo(input model.NewVideo) *model.Video {
	db := config.ConnectDatabase()
	defer db.Close()

	newVideo := model.Video{
		Src:    input.Src,
		OnDuty: 2,
	}

	db.Table("video").Create(&newVideo)

	return &newVideo
}

func UpdateVideoByID(input model.EditedVideo) (*model.Video, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedVideo := model.Video{
		ID:     input.ID,
		Src:    input.Src,
		OnDuty: input.OnDuty,
	}

	err := db.Table("video").Where("id = ?", input.ID).Update(&editedVideo).Error

	return &editedVideo, err
}

func DeleteVideoByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("video").Where("id = ?",id).Delete(&model.Video{}).Error

	return id,err
}