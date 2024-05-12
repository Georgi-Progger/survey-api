package mapper

import (
	"fmt"

	"github.com/Georgi-Progger/survey-api/internal/model"
)

func UserToUserWithInfo(user model.User, info model.UserWithInfo) model.UserWithInfo {
	info.Id = user.Id
	info.Password = user.Password
	info.Phonenumber = user.Phonenumber
	info.RoleId = user.RoleId
	return info
}

func CandidateToUserWithInfo(cnd model.Candidate, info model.UserWithInfo) model.UserWithInfo {
	info.DateOfBirth = cnd.BirthDate
	info.City = cnd.City
	info.CreationDate = cnd.CreationDate
	info.Education = cnd.Education
	info.Email = cnd.Email
	//cnd.EmployeeInfo
	info.FirstName = cnd.FirstName
	info.LastName = cnd.LastName
	info.MiddleName = cnd.MiddleName
	info.ReasonDismissal = cnd.ReasonDismissal
	info.ResumePath = cnd.ResumePath
	info.Id = cnd.UserId
	info.YearWorkExperience = fmt.Sprint(cnd.YearWorkExperience)
	return info
}
