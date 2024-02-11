package survey

type experience uint

const (
	Not experience = iota
	LessThanAYear
	OneYear
	TwoYear
	ThreeYear
	FourYear
	MoreThanAYear
)

type (
	Candidate struct {
		Id                 uint
		FirstName          string
		LastName           string
		MiddleName         string
		BirthDate          string
		City               string
		Education          string
		ReasonDismissal    string
		Email              string
		PhoneNumper        string
		YearWorkExperience experience
		EmployeeInfo       string
	}
	Interview struct {
		Id            uint
		InterviewName string
		Questions     []InterviewQuestion
	}
	InterviewQuestion struct {
		Id          uint
		TextAnswer  string
		InterviewId Interview
	}
	ProblemsTheme struct {
		Id                uint
		ProblemSThemeName string
	}
	ProblemQuestion struct {
		Id              uint
		TextAnswer      string
		ProblemsThemeId ProblemsTheme
	}
	ProblemSolvingTasks struct {
		Id              uint
		TextQuestion    string
		ProblemsThemeId ProblemsTheme
	}
	File struct {
		Id   uint
		Path string
	}
)
