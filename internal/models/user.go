package models

type User struct {
	Id         string `json:"user_id"`
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
}
