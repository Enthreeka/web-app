package entity

type Account struct {
	Id              int    `json:"id"`
	UserId          string `json:"user_id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Photo           string `json:"photo"`
	Subscribe       bool   `json:"subscribe"`
	NameTask        string `json:"name_task"`
	DescriptionTask string `json:"description_task"`
	DateSignUp      string `json:"date_signup"`
}
