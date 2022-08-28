package service

import (
	"myapp/config"
	"myapp/graph/model"
)

//dataloader
func GetPhotosByWhereInCategoryID(params []*model.PhotoLoaderParams) ([][]*model.Photo, []error) {
	db := config.ConnectDatabase()
	defer db.Close()

	categoryIDs := []int{}
	status := params[0].Status
	for _, v := range params {
		categoryIDs = append(categoryIDs, v.ID)
	}

	var result []*model.Photo
	db.Table("images").Where("images_category_id IN (?) AND on_duty IN (?)", categoryIDs, status).Find(&result)

	photoMappedByCategoryId := map[int][]*model.Photo{}
	for _, v := range result {
		photoMappedByCategoryId[v.ImagesCategoryID] = append(photoMappedByCategoryId[v.ImagesCategoryID], v)
	}

	response := make([][]*model.Photo, len(categoryIDs))
	for i, categoryId := range categoryIDs {
		response[i] = photoMappedByCategoryId[categoryId]
	}

	return response, nil
}

func GetImagesCategoryByWhereIn(imagesCategoryID []int) ([]*model.ImagesCategory, []error) {
	db := config.ConnectDatabase()
	defer db.Close()

	var result []*model.ImagesCategory
	db.Table("images_category").Where("id IN (?)", imagesCategoryID).Find(&result)

	ImagesCategoryDetailMappedByID := map[int]*model.ImagesCategory{}

	for _, v := range result {
		ImagesCategoryDetailMappedByID[v.ID] = v
	}

	array := make([]*model.ImagesCategory, len(imagesCategoryID))
	for index, imagesCategoryID := range imagesCategoryID {
		array[index] = ImagesCategoryDetailMappedByID[imagesCategoryID]
	}

	return array, nil
}

//service
func GetCarouselTopImages(status []*int) []*model.TopImage {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2, 1}
	if status != nil {
		defaultStatus = []int{}
		for _, v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.TopImage
	db.Table("top_images").Where("on_duty IN (?)", defaultStatus).Find(&result)

	return result
}

func GetImageCategory(status []*int) []*model.ImagesCategory {
	db := config.ConnectDatabase()
	defer db.Close()

	defaultStatus := []int{2, 1}
	if status != nil {
		defaultStatus = []int{}
		for _, v := range status {
			defaultStatus = append(defaultStatus, *v)
		}
	}

	var result []*model.ImagesCategory
	db.Table("images_category").Where("on_duty IN (?)", defaultStatus).Find(&result)

	return result
}

//mutations
func CreateTopImage(input model.NewTopImage) *model.TopImage {
	db := config.ConnectDatabase()
	defer db.Close()

	newTopImage := model.TopImage{
		Src:     input.Src,
		Caption: input.Caption,
		OnDuty:  2,
	}
	db.Table("top_images").Create(&newTopImage)

	return &newTopImage
}

func CreateImageCategory(input model.NewImagesCategory) *model.ImagesCategory {
	db := config.ConnectDatabase()
	defer db.Close()

	newImageCategory := model.ImagesCategory{
		Category: input.Category,
		OnDuty:   2,
	}

	db.Table("images_category").Create(&newImageCategory)

	return &newImageCategory
}

func CreateImage(input model.NewPhoto) *model.Photo {
	db := config.ConnectDatabase()
	defer db.Close()

	newPhoto := model.Photo{
		Src:              input.Src,
		Width:            input.Width,
		Height:           input.Height,
		ImagesCategoryID: input.ImagesCategoryID,
		OnDuty:           2,
	}

	db.Table("images").Create(&newPhoto)

	return &newPhoto
}

func UpdateTopImageByID(input model.EditedTopImage) (*model.TopImage, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedTopImage := model.TopImage{
		ID:      input.ID,
		Src:     input.Src,
		Caption: input.Caption,
		OnDuty:  input.OnDuty,
	}

	err := db.Table("top_images").Where("id = ?", input.ID).Update(&editedTopImage).Error

	return &editedTopImage, err
}

func UpdateImageCategoryByID(input model.EditedImagesCategory) (*model.ImagesCategory, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedImageCategory := model.ImagesCategory{
		ID:       input.ID,
		Category: input.Category,
		OnDuty:   input.OnDuty,
	}

	err := db.Table("images_category").Where("id = ?", input.ID).Update(&editedImageCategory).Error

	return &editedImageCategory, err
}

func UpdateImageByID(input model.EditedPhoto) (*model.Photo, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	editedPhoto := model.Photo{
		ID:               input.ID,
		Src:              input.Src,
		Width:            input.Width,
		Height:           input.Height,
		ImagesCategoryID: input.ImagesCategoryID,
		OnDuty:           input.OnDuty,
	}

	err := db.Table("images").Where("id = ?", input.ID).Update(&editedPhoto).Error

	return &editedPhoto, err
}

func GetImageWithoutCategory() []*model.Photo {
	db := config.ConnectDatabase()
	defer db.Close()

	var images []*model.Photo

	db.Table("images").Find(&images)

	return images
}

func DeleteImageByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("images").Where("id = ?",id).Delete(&model.Photo{}).Error

	return id,err
}

func DeleteTopImagesByID(id int) (int,error){
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("top_images").Where("id = ?",id).Delete(&model.TopImage{}).Error

	return id,err
}

func DeleteImageCategoryByID(id int) (int,error) {
	db := config.ConnectDatabase()
	defer db.Close()

	err := db.Table("images_category").Where("id = ?",id).Delete(&model.ImagesCategory{}).Error

	return id,err
}
