package models

type ToDo struct {
	Id   int
	Note string `json:"note"`
	Date string
}
