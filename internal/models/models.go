package models

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type FriendRequest struct {
	UserId   int `json:"user_id"`
	FriendId int `json:"friend_id"`
}
