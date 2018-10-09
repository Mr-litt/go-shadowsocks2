package monitor

import (
	"strings"
	"fmt"
	"net/smtp"
)

type AuthConfig  struct{
	emailSwitch bool
	username string
	password string
	host string
	addr string
	to []string
	subject string
}

type Email struct {
	authConfig *AuthConfig
}

func InitEmail(emailSwitch bool, emailUserName string, emailPassword string) * Email{
	email := &Email{}
	authConfig := &AuthConfig{}
	authConfig.emailSwitch = emailSwitch
	authConfig.username = emailUserName
	authConfig.password = emailPassword
	authConfig.host = "smtp.qq.com"
	authConfig.addr = "smtp.qq.com:25"
	authConfig.to = [] string {emailUserName}
	authConfig.subject = "proxy service monitor"

	if emailUserName == "" || emailPassword == "" {
		authConfig.emailSwitch = false
	}

	email.authConfig = authConfig
	return email
}


func (email * Email) Send (body string) {

	// 消息体
	msg := []byte("To: " + strings.Join(email.authConfig.to, ",") + "\r\nFrom: " + email.authConfig.username +
		"<" + email.authConfig.username + ">\r\nSubject: " + email.authConfig.subject + "\r\n" + "Content-Type: text/plain; charset=UTF-8" +
		"\r\n\r\n" + body)

	fmt.Printf("send msg body: %s", msg)

	// 打开开关时发送邮件
	if email.authConfig.emailSwitch {
		auth := smtp.PlainAuth("", email.authConfig.username, email.authConfig.password, email.authConfig.host)
		err := smtp.SendMail(email.authConfig.addr, auth, email.authConfig.username, email.authConfig.to, msg)
		if err != nil {
			fmt.Printf("send mail error: %v", err)
		}
		fmt.Println("send msg Success:")
	}
}

func test() {
	/*
	auth := smtp.PlainAuth("", "953637695@qq.com", "password", "smtp.qq.com")
    to := []string{"874560965@qq.com"}
    nickname := "test"
    user := "953637695@qq.com"
    subject := "test mail"
    content_type := "Content-Type: text/plain; charset=UTF-8"
    body := "This is the email body."
    msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
        "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
    if err != nil {
        fmt.Printf("send mail error: %v", err)
    }*/
}
