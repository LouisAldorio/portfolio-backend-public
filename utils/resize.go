package utils

import "myapp/graph/model"

func Resize(link string) *model.File{
	return &model.File{
		Small: "https://res.cloudinary.com/demo/image/fetch/w_250,f_auto/" + link,
		Medium: "https://res.cloudinary.com/demo/image/fetch/w_500,f_auto/" + link,
		Large: "https://res.cloudinary.com/demo/image/fetch/w_1000,f_auto/" + link,
		ExtraLarge: "https://res.cloudinary.com/demo/image/fetch/w_1300,f_auto/" + link,
	}
}