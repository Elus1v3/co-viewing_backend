package models

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
