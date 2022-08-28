package model

type NewMail struct {
	Subject          string `json:"subject"`
	Message          string `json:"message"`
	DestinationEmail string `json:"destination_email"`
}

type EducationLoaderParams struct {
	ID     int      `json:"id"`
	Status []int `json:"status"`
}

type PhotoLoaderParams struct {
	ID     int      `json:"id"`
	Status []int `json:"status"`
}
