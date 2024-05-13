package model

type (
	VQuestion struct {
		Id           int    `json:"id"`
		QuestionText string `json:"question_text"`
	}
	VQuestionAndAnswer struct {
		Question   string `json:"question"`
		AnswerPath string `json:"answer_path"`
	}
	QuestionAnswer struct {
		VQuestionId int
		UserId      int
		Path        string
	}
	TestQuestion struct {
		Id       int                  `json:"id"`
		Question string               `json:"question"`
		Answers  []TestQuestionAnswer `json:"answers"`
	}
	TestQuestionAnswer struct {
		Id     int    `json:"id"`
		Answer string `json:"answer"`
	}
	UserTestAnswer struct {
		TestQuestionId int `json:"test_question_id"`
		TestAnswerId   int `json:"test_answer_id"`
	}
	PonomarResult struct {
		Hysteroid   float64 `json:"hysteroid"`
		Epileptoid  float64 `json:"epileptoid"`
		Paranoid    float64 `json:"paranoid"`
		Emotional   float64 `json:"emotional"`
		Schizoid    float64 `json:"schizoid"`
		Hyperthymic float64 `json:"hyperthymic"`
		Anxious     float64 `json:"anxious"`
	}
)
