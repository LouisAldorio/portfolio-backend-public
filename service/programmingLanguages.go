package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetAllProgrammingLanguages(status []*int) []*model.ProgrammingLanguage {

	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2, 1}
	if status != nil {
		defaultStatus = []int{}
		for _, v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.ProgrammingLanguage
	db.Table("programming_languages").Where("on_duty IN (?)", defaultStatus).Find(&result)

	return result
}

func CreateProgrammingLanguage(input model.NewProgrammingLanguage) *model.ProgrammingLanguage {
	db := config.ConnectDatabase()
	defer db.Close()

	newProgrammingLanguage := model.ProgrammingLanguage{
		Name:   input.Name,
		Value:  input.Value,
		Color:  input.Color,
		Total:  input.Total,
		OnDuty: 2,
	}

	db.Table("programming_languages").Create(&newProgrammingLanguage)

	return &newProgrammingLanguage
}

func UpdateProgrammingLanguageByID(input model.EditedProgrammingLanguage) (*model.ProgrammingLanguage, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedProgrammingLanguage := model.ProgrammingLanguage{
		ID:     input.ID,
		Name:   input.Name,
		Value:  input.Value,
		Color:  input.Color,
		Total:  input.Total,
		OnDuty: input.OnDuty,
	}

	err := db.Table("programming_languages").Where("id = ?", input.ID).Update(&editedProgrammingLanguage).Error

	return &editedProgrammingLanguage, err
}

func DeleteProgrammingLanguageByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("programming_languages").Where("id = ?",id).Delete(&model.ProgrammingLanguage{}).Error

	return id,err
}
