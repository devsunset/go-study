package main

import (
	"net/smtp"
)

// SMTP Example
func main() {
	auth := smtp.PlainAuth("", "sender@live.com", "pwd", "smtp.live.com")

	from := "sender@live.com"

	to := []string{"receiver@live.com"}

	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.live.com:587", auth, from, to, msg)

	if err != nil {
		panic(err)
	}
}
