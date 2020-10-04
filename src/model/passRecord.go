package model

type PassRecord struct {
	Id uint `json:"id"`
	Nickname string `json:"nickname"`
	Website string `json:"website"`
	Email string `json:"email"`
	Password string	`json:"password"`
	OtherData string `json:"other_data"`

}
