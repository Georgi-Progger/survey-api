package survey

type experience uint

const (
	LessThanAYear experience = iota
	OneYear
	TwoYear
	MoreThanAYear
)

type reasonDismissal uint

const (
	YourWish reasonDismissal = iota
	Dismissal
	NegativeGroup
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
		ReasonDismissal    reasonDismissal
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
)
