package entity

type Task struct {
	Id              int    `json:"id"`
	AccountId       string `json:"user_id"`
	NameTask        string `json:"name_task"`
	DescriptionTask string `json:"description_task"`
	DateCreate      string `json:"date_task"`
}
