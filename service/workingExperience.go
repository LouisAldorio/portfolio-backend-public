package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetWorkingExperience(status []*int) []*model.WorkingExperience {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2,1}
	if status != nil {
		defaultStatus = []int{}
		for _,v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.WorkingExperience
	db.Table("working_experience").Where("on_duty IN (?)",defaultStatus).Find(&result)

	return result
}

//mutation
func CreateWorkingExperience(input model.NewWorkingExperience) *model.WorkingExperience {
	db := config.ConnectDatabase()
	defer db.Close()

	newWorkingExperience := model.WorkingExperience{
		CompanyLogo: input.CompanyLogo,
		CompanyName: input.CompanyName,
		Address:     input.Address,
		StartDate:   input.StartDate,
		Image:       input.Image,
		EndDate:     input.EndDate,
		Position:    input.Position,
		Overview:    input.Overview,
		Description: input.Description,
		OnDuty:      2,
	}

	db.Table("working_experience").Create(&newWorkingExperience)

	return &newWorkingExperience
}

func UpdateWorkingExperienceByID(input model.EditedWorkingExperience) (*model.WorkingExperience, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedWorkingExperience := model.WorkingExperience{
		ID:          input.ID,
		CompanyLogo: input.CompanyLogo,
		CompanyName: input.CompanyName,
		Address:     input.Address,
		StartDate:   input.StartDate,
		Image:       input.Image,
		EndDate:     input.EndDate,
		Position:    input.Position,
		Overview:    input.Overview,
		Description: input.Description,
		OnDuty:      input.OnDuty,
	}

	err := db.Table("working_experience").Where("id = ?",input.ID).Update(&editedWorkingExperience).Error
	return &editedWorkingExperience, err
}

func DeleteWorkingExperienceByID(id int)(int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("working_experience").Where("id = ?",id).Delete(&model.WorkingExperience{}).Error

	return id,err
}
