package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetAllSkills(status []*int) []*model.Skill{

	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2,1}
	if status != nil {
		defaultStatus = []int{}
		for _,v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.Skill
	db.Table("skills").Where("on_duty IN (?)",defaultStatus).Find(&result)

	return result
}


func CreateSkill(input model.NewSkill) *model.Skill {
	db := config.ConnectDatabase()
	defer db.Close()

	newSkill := model.Skill{
		Logo: input.Logo,
		Name: input.Name,
		State: input.State,
		OnDuty: 2,
	}

	db.Table("skills").Create(&newSkill)

	return &newSkill
}


func UpdateSkillByID(input model.EditedSkill) (*model.Skill,error){
	db := config.ConnectDatabase()
	defer db.Close()

	editedSkill := model.Skill{
		ID: input.ID,
		Logo: input.Logo,
		Name: input.Name,
		State: input.State,
		OnDuty: input.OnDuty,
	}

	err := db.Table("skills").Where("id = ?",input.ID).Update(&editedSkill).Error

	return &editedSkill,err
}

func DeleteSkillByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("skills").Where("id = ?",id).Delete(&model.Skill{}).Error

	return id,err
}