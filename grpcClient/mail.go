package grpcClient

import (
	"context"
	"myapp/graph/model"

	"github.com/LouisAldorio/louisaldorio-utils-go/mail"
)

func SendMail(input model.NewMail) string {

	MailCon, con := mail.Connect(mail.ConnectionOption{})
	defer con.Close()

	response, err := MailCon.SendMail(context.Background(), &mail.SendMailRequest{
		Email:       input.DestinationEmail,
		FirstName:   "",
		LastName:    "",
		PhoneNumber: "",
		Messages:    "",
		Subject:     "",
	})

	if err != nil {
		panic(err)
	}

	return response.StatusMessage
}
