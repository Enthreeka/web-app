package entity

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
