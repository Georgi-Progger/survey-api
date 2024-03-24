package model

type User struct {
	Id          int
	RoleId      int    `json:"role_id"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"password"`
}
