package utils

import (
	"gopkg.in/gomail.v2"

	"fmt"
)

var (
	smtpHost    = "smtp.163.com"
	smtpPort    = 25
	senderEmail = "18810535172@163.com"
	// senderPassword = "password123@"
	senderPassword = "MERQTNPVJDVBNCIH"

	msgPrefix = messagePrefix()
	dialer    = gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
)

func SendVerifyMail(to, code []byte) error {
	// return sendMail(to, append(msgPrefix, code...))

	fmt.Printf("[%s]:[%s]\n", string(to), string(code))

	return nil
}

func sendMail(to, message []byte) error {
	m := gomail.NewMessage()

	m.SetHeader("From", senderEmail)
	m.SetHeader("To", string(to))
	m.SetHeader("Subject", "Verify Code")
	m.SetBody("text/plain", string(message))

	return dialer.DialAndSend(m)
}

func messagePrefix() []byte {
	warning := "You are register our service.If you didn't do it, ignore this mail.\n"
	info := "verify code:"
	return []byte(warning + info)
}
