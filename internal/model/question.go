package model

type VQuestion struct {
	Id           int    `json:"id"`
	QuestionText string `json:"question_text"`
}

type QuestionAnswer struct {
	VQuestionId int
	UserId      int
	Path        string 
}
