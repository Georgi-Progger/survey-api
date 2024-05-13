package handler

import (
	"log"
	"net/http"

	"github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllTQuestions(c echo.Context) error {
	res, err := h.services.TQuestion.GetAll()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get all video test questions: " + err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) insertTQuestionAnswers(c echo.Context) error {
	var ans []model.UserTestAnswer
	if err := c.Bind(&ans); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON provided: " + err.Error()})
	}
	err := h.services.TQuestion.InsertAnswers(jwt.GetUserIdFromContext(c), ans)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to save answers in db: " + err.Error()})
	}
	return c.NoContent(http.StatusCreated)
}

func (h *Handler) getPonomarUserResult(c echo.Context) error {
	testAnsExplain := [][]string{
		[]string{"Hysteroid", "Emotional", "Epileptoid", "Hyperthymic", "Paranoid", "Anxious", "Schizoid"}, // A
		[]string{"Emotional", "Epileptoid", "Schizoid", "Hyperthymic", "Hysteroid", "Paranoid", "Anxious"}, // Б
		[]string{"Paranoid", "Schizoid", "Emotional", "Epileptoid", "Hysteroid", "Hyperthymic", "Anxious"}, // В
		[]string{"Anxious", "Epileptoid", "Hyperthymic", "Schizoid", "Emotional", "Paranoid", "Hysteroid"}, // Г
		[]string{"Epileptoid", "Hyperthymic", "Emotional", "Anxious", "Paranoid", "Hysteroid", "Schizoid"}, // Д
		[]string{"Hyperthymic", "Epileptoid", "Schizoid", "Anxious", "Paranoid", "Hysteroid", "Emotional"}, // Е
		[]string{"Epileptoid", "Hysteroid", "Emotional", "Paranoid", "Anxious", "Schizoid", "Hyperthymic"}, // Ж
		[]string{"Paranoid", "Anxious", "Epileptoid", "Schizoid", "Hysteroid", "Emotional", "Hyperthymic"}, // З
		[]string{"Hyperthymic", "Epileptoid", "Paranoid", "Anxious", "Emotional", "Schizoid", "Hysteroid"}, // И
		[]string{"Anxious", "Hysteroid", "Emotional", "Epileptoid", "Schizoid", "Hyperthymic", "Paranoid"}, // К
		[]string{"Hysteroid", "Epileptoid", "Anxious", "Hyperthymic", "Schizoid", "Paranoid", "Emotional"}, // Л
		[]string{"Schizoid", "Epileptoid", "Hysteroid", "Anxious", "Paranoid", "Emotional", "Hyperthymic"}, // М
		[]string{"Epileptoid", "Paranoid", "Anxious", "Hysteroid", "Hyperthymic", "Emotional", "Schizoid"}, // Н
	}

	userAns, err := h.services.TQuestion.GetUserAnswers(jwt.GetUserIdFromContext(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed get user answers from db: " + err.Error()})
	}
	var ansResult model.PonomarResult
	for _, ans := range userAns {
		switch testAnsExplain[ans.TestQuestionId-1][ans.TestAnswerId-1] {
		case "Hysteroid":
			ansResult.Hysteroid++
		case "Epileptoid":
			ansResult.Epileptoid++
		case "Paranoid":
			ansResult.Paranoid++
		case "Emotional":
			ansResult.Emotional++
		case "Schizoid":
			ansResult.Schizoid++
		case "Hyperthymic":
			ansResult.Hyperthymic++
		case "Anxious":
			ansResult.Anxious++
		}
	}
	kf := 100.0/13
	ansResult.Anxious *= kf 
	ansResult.Emotional *= kf 
	ansResult.Epileptoid *= kf 
	ansResult.Hyperthymic *= kf 
	ansResult.Hysteroid *= kf 
	ansResult.Paranoid *= kf 
	ansResult.Schizoid *= kf 

	return c.JSON(http.StatusOK, ansResult)
}
