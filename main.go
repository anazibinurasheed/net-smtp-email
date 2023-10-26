package main

import (
	"fmt"
	"net/smtp"
)

const (
	email = ""
	pass  = ""
)

func sendMail(subject, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		email,
		pass,
		"smtp.gmail.com",
	)

	//ipvp rtgh xjmv mled
	msg := subject + "\n" + body
	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		email,
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func sendHtmlMail(subject, html string, to []string) {
	auth := smtp.PlainAuth(
		"",
		email,
		pass,
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html
	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		email,
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	toMail := ""
	sendMail("", "<h1>I'm a heading</h1><p>I'm a paragraph</p>", []string{toMail})
}
