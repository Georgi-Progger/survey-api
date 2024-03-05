package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type SmsSenderService struct {
}

func NewSmsSenderService() *SmsSenderService {
	return &SmsSenderService{}
}

func (s *SmsSenderService) Send(destination, message string) bool {
	str := fmt.Sprintf("https://%s:%s@gate.smsaero.ru/v2/sms/send?number=%s&text=%s&sign=SMS Aero",
		os.Getenv("SMSAERO_LOGIN"), os.Getenv("SMSAERO_API_KEY"), destination, message)
	resp, err := http.Get(str)
	if err != nil {
		log.Fatal("Не удалось отправить сообщение:" + err.Error())
		return false
	}
	defer resp.Body.Close()
	return true
}
