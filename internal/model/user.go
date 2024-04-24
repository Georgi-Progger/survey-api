package model

type User struct {
	Id          int
	RoleId      int    `json:"role_id"`
	Email       string `json:"-"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"-"`
}

type UserWithName struct {
	Id          int
	RoleId      int    `json:"role_id"`
	Email       string `json:"-"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"-"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
}
