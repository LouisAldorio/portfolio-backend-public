package service

import (
	"myapp/config"
	"myapp/graph/model"
)

//dataloader function
func GetEducationByWhereInCategoryID(params []*model.EducationLoaderParams) ([][]*model.Education, []error) {
	db := config.ConnectDatabase()
	defer db.Close()

	categoryIDs := []int{}
	status := params[0].Status

	for _, v := range params {
		categoryIDs = append(categoryIDs, v.ID)
	}

	var education []*model.Education
	db.Table("education").Where("education_category_id IN (?) AND on_duty IN (?)", categoryIDs, status).Find(&education)

	educationByCategoryId := map[int][]*model.Education{}
	for _, v := range education {
		educationByCategoryId[v.EducationCategoryID] = append(educationByCategoryId[v.EducationCategoryID], v)
	}

	educations := make([][]*model.Education, len(categoryIDs))
	for i, categoryId := range categoryIDs {
		educations[i] = educationByCategoryId[categoryId]
	}

	return educations, nil
}

//service function
func GetEducationCategory(status []*int) []*model.EducationCategory {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2, 1}
	if status != nil {
		defaultStatus = []int{}
		for _, v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var category []*model.EducationCategory
	err := db.Table("education_category").Where("on_duty IN (?)", defaultStatus).Find(&category).Error
	if err != nil {
		panic(err)
	}

	return category
}

func UpdateEducationCategoryByID(input model.EditedEducationCategory) (*model.EducationCategory,error){
	db := config.ConnectDatabase()
	defer db.Close()

	editedImageCategory := model.EducationCategory{
		ID: input.ID,
		Title: input.Title,
		Logo: input.Logo,
		OnDuty: input.OnDuty,
	}

	err := db.Table("education_category").Where("id = ?",input.ID).Update(&editedImageCategory).Error

	return &editedImageCategory,err
}

func GetAllEducation() []*model.Education {
	db := config.ConnectDatabase()
	defer db.Close()

	var education []*model.Education
	db.Table("education").Find(&education)

	return education
}

func GetEducationCategoryByID(educationCategoryID int) *model.EducationCategory {
	db := config.ConnectDatabase()
	defer db.Close()

	var category []*model.EducationCategory
	err := db.Table("education_category").Where("id = ?", educationCategoryID).Find(&category).Error
	if err != nil {
		panic(err)
	}

	if len(category) == 0 {
		panic("Record Not Found!")
	}

	return category[0]
}

//mutation
func CreateEducation(input model.NewEducation) *model.Education {
	db := config.ConnectDatabase()
	defer db.Close()

	newEducation := model.Education{
		Name:                input.Name,
		Link:                input.Link,
		Logo:                input.Logo,
		Image:               input.Image,
		StartYear:           input.StartYear,
		EndYear:             input.EndYear,
		Address:             input.Address,
		EducationCategoryID: input.EducationCategoryID,
		OnDuty:              2,
	}

	db.Table("education").Create(&newEducation)

	return &newEducation
}

func UpdateEducationByID(input model.EditedEducation) (*model.Education, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedEducation := model.Education{
		ID:                  input.ID,
		Name:                input.Name,
		Link:                input.Link,
		Logo:                input.Logo,
		Image:               input.Image,
		StartYear:           input.StartYear,
		EndYear:             input.EndYear,
		Address:             input.Address,
		EducationCategoryID: input.EducationCategoryID,
		OnDuty:              input.OnDuty,
	}

	err := db.Table("education").Where("id = ?", input.ID).Update(&editedEducation).Error

	return &editedEducation, err
}

func DeleteEducationByID(id int)(int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("education").Where("id = ?",id).Delete(&model.Education{}).Error

	return id,err
}

//education category
func CreateEducationCategory(input model.NewEducationCategory) *model.EducationCategory{
	db := config.ConnectDatabase()
	defer db.Close()

	newEducationCategory := model.EducationCategory{
		Logo: input.Logo,
		Title: input.Title,
		OnDuty: 2,
	}
	db.Table("education_category").Create(&newEducationCategory)

	return &newEducationCategory
}

func DeleteEducationCategoryByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("education_category").Where("id = ?",id).Delete(&model.EducationCategory{}).Error

	return id,err
}
