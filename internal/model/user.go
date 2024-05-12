package model

type User struct {
	Id          int
	RoleId      int    `json:"role_id"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"-"`
}

type UserWithInfo struct {
	Id          int
	RoleId      int    `json:"role_id"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"-"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	City string `json:"city"`
	Education string `json:"education"`
	ReasonDismissal string `json:"reason_dismissal"`
	YearWorkExperience string `json:"year_work_experience"`
	ResumePath string `json:"resume_path"`
	CreationDate string `json:"creation_date"`
}
