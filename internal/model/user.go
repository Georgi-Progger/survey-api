package model

type User struct {
	Id          int
	RoleId      int    `json:"role_id"`
	Email       string `json:"-"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"-"`
}
