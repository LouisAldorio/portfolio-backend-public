package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetBuiltProjects(status []*int) []*model.BuiltProject {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2,1}
	if status != nil {
		defaultStatus = []int{}
		for _,v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.BuiltProject
	db.Table("built_projects").Where("on_duty IN (?)",defaultStatus).Find(&result)

	return result
}

//mutation
func CreateBuiltProject(input model.NewBuiltProject) *model.BuiltProject {
	db := config.ConnectDatabase()
	defer db.Close()

	newBuiltProject := model.BuiltProject{
		Thumbnail:      input.Thumbnail,
		Name:           input.Name,
		Overview:       input.Overview,
		GithubLink:     input.GithubLink,
		DeploymentLink: input.DeploymentLink,
		OnDuty:         2,
	}

	db.Table("built_projects").Create(&newBuiltProject)

	return &newBuiltProject
}

func UpdateBuiltProjectByID(input model.EditedBuiltProject) (*model.BuiltProject, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedBuiltProject := model.BuiltProject{
		ID:             input.ID,
		Thumbnail:      input.Thumbnail,
		Name:           input.Name,
		Overview:       input.Overview,
		GithubLink:     input.GithubLink,
		DeploymentLink: input.DeploymentLink,
		OnDuty:         input.OnDuty,
	}

	err := db.Table("built_projects").Where("id = ?", input.ID).Update(&editedBuiltProject).Error

	return &editedBuiltProject, err
}

func DeleteBuiltProjectByID(id int)(int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("built_projects").Where("id = ?",id).Delete(&model.BuiltProject{}).Error

	return id,err
}
