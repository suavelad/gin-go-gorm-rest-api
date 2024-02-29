package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)



func SendEmailTask(input interface{}) interface{} {
	fmt.Println("SendEmailTask started")
	params := input.([]interface{})
	to := fmt.Sprintf("%v", params[0])
	attachment_path := fmt.Sprintf("%v", params[1])
	subject := fmt.Sprintf("%v", params[2])
	body := fmt.Sprintf("%v", params[3])


	SendEmail(to, attachment_path, subject, body)
	fmt.Println("SendEmailTask done")
	return nil
}

func SendEmail(to string, attachment_path string, subject string, body string) {
	fmt.Println("SendEmail started")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Error converting SMTP_PORT to int:", err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "Sunday Support <sunnexajayi@gmail.com>")
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if attachment_path != "" {
		m.Attach(attachment_path)

	}

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), smtpPort, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return
}
