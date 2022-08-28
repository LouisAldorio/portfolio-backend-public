package service

import (
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/utils"
)

func Login(input model.OwnerCredential) (*model.Owner, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	var password []string

	db.Table("owner").Where("username = ?", input.Username).Pluck("password",&password)

	if len(password) == 0{
		return nil, fmt.Errorf("You Are Not The Owner Of this Site!")
	}

	isValid := utils.CheckPassword(password[0],input.Password)
	if isValid == false {
		return nil, fmt.Errorf("You Are Not The Owner Of this Site!")
	}

	token, err := utils.CreateToken(input.Username)
	if err != nil {
		return nil, err
	}

	return &model.Owner{
		Username:    input.Username,
		AccessToken: token,
	}, nil
}
