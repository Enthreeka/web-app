package entity

type Task struct {
	Id              int    `json:"id"`
	AccountId       int    `json:"account_id"`
	NameTask        string `json:"name_task"`
	DescriptionTask string `json:"description_task"`
	DateCreate      string `json:"date_task"`
}
