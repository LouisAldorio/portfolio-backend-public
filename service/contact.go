package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func GetSocialMediaInfo() *model.SocialMediaInfo {
	db := config.ConnectDatabase()
	defer db.Close()

	var result []*model.SocialMediaInfo
	db.Table("social_media_info").Find(&result)

	if len(result) > 0 {
		return result[0]
	}

	return nil
}

func UpdateSocialMediaInfo(input model.EditSocialMediInfo) (*model.SocialMediaInfo, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	newSocialMediaInfo := model.SocialMediaInfo{
		ID:        input.ID,
		Phone:     input.Phone,
		Mail:      input.Mail,
		Location:  input.Location,
		Instagram: input.Instagram,
		Whatsapp:  input.Whatsapp,
	}

	err := db.Table("social_media_info").Where("id = ?", input.ID).Update(&newSocialMediaInfo).Error

	return &newSocialMediaInfo, err
}

func FindSubscriberByEmail(email string) (model.Subscriber, error) {
	db := config.ConnectDatabase()
	defer db.Close()

	var subscriber model.Subscriber
	err := db.Table("subscribers").Where("email = ?", email).First(&subscriber).Error

	return subscriber, err
}

func CreateSubscriber(input model.MailInput) string {
	var subscriber model.Subscriber

	subscriber, err := FindSubscriberByEmail(input.Email)
	if err != nil {
		db := config.ConnectDatabase()
		defer db.Close()

		subscriber = model.Subscriber{
			FirstName:   input.FirstName,
			LastName:    input.LastName,
			Email:       input.Email,
			PhoneNumber: input.PhoneNumber,
		}

		db.Table("subscribers").Create(&subscriber)
	}

	CreateMessage(input.Message, subscriber.ID)

	// response := grpcClient.SendMail(model.NewMail{
	// 	Subject:          "Subscribe To Louis Aldorio",
	// 	Message:          input.Message,
	// 	DestinationEmail: input.Email,
	// })

	return "Mail Sent!"
}

func CreateMessage(message string, subscriberID int) {
	db := config.ConnectDatabase()
	defer db.Close()

	newMessage := model.Mail{
		Message:      message,
		SubscriberID: subscriberID,
	}

	db.Table("messages").Create(&newMessage)
}

func GetMessages() []*model.Mail {
	db := config.ConnectDatabase()
	defer db.Close()

	var messages []*model.Mail
	db.Table("messages").Find(&messages)

	return messages
}

//dataloaders
func GetSubscriberByWhereIn(subscriberIds []int) ([]*model.Subscriber, []error) {
	db := config.ConnectDatabase()
	defer db.Close()

	var subscribers []*model.Subscriber

	db.Table("subscribers").Where("id IN (?)", subscriberIds).Find(&subscribers)

	subsciberMappedById := map[int]*model.Subscriber{}

	for _, v := range subscribers {
		subsciberMappedById[v.ID] = v
	}

	response := make([]*model.Subscriber, len(subscriberIds))
	for i, id := range subscriberIds {
		response[i] = subsciberMappedById[id]
	}

	return response, nil
}
