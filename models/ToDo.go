package models

type ToDo struct {
	Id   int
	Note string `json:"todo,omitempty"`
	Date string
}
