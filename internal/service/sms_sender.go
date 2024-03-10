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

func (s *SmsSenderService) Send(destination, message string) error {
	request := fmt.Sprintf("https://%s:%s@gate.smsaero.ru/v2/sms/send?number=%s&text=%s&sign=SMS Aero",
		os.Getenv("SMSAERO_LOGIN"), os.Getenv("SMSAERO_API_KEY"), destination, message)
	resp, err := http.Get(request)
	if err != nil {
		log.Fatal("Не удалось отправить сообщение:" + err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}
