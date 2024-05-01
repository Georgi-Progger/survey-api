package model

type (
	VQuestion struct {
		Id           int    `json:"id"`
		QuestionText string `json:"question_text"`
	}

	QuestionAnswer struct {
		VQuestionId int
		UserId      int
		Path        string
	}
	TestQuestion struct {
		Id int `json:"id"`
		Question string `json:"question"`
		Answers []TestQuestionAnswer `json:"answers"`
	}
	TestQuestionAnswer struct {
		Id int `json:"id"`
		Answer string `json:"answer"`
	}
	UserTestAnswer struct {
		TestQuestionId int `json:"test_question_id"`
		TestAnswerId int `json:"test_answer_id"`
	}
	PonomarResult struct {
		Hysteroid int `json:"hysteroid"`
		Epileptoid int `json:"epileptoid"`
		Paranoid int `json:"paranoid"`
		Emotional int `json:"emotional"`
		Schizoid int `json:"schizoid"`
		Hyperthymic int `json:"hyperthymic"`
		Anxious int `json:"anxious"`
	}
)
