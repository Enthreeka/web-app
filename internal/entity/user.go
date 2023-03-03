package entity

type User struct {
	Id       int      `json:"id"`
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Account  *Account `json:"account"`
}
