package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

type mail struct {
	user   string
	passwd string
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//初始化用户名和密码
func New(u string, p string) mail {
	temp := mail{user: u, passwd: p}
	return temp
}

// 标题 文本 目标邮箱
func (m mail) Send(title string, text string, toId string) {
	auth := smtp.PlainAuth("", m.user, m.passwd, "smtp.gmail.com")

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig)
	check(err)

	client, err := smtp.NewClient(conn, "smtp.gmail.com")
	check(err)

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(m.user); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(toId); err != nil {
		log.Panic(err)
	}

	w, err := client.Data()
	check(err)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", m.user, toId, title, text)

	_, err = w.Write([]byte(msg))
	check(err)

	err = w.Close()
	check(err)

	client.Quit()
}

func main(){
	// foo := New("username@gmail.com","passwd") //自己的邮箱地址和密码
	foo := New("klaus@sunday37.com","klauswyxdrqq0930") //自己的邮箱地址和密码
	// foo := New("jimfromchinagd@gmail.com","zyxfdbs123") //自己的邮箱地址和密码
	foo.Send("标题","邮箱内容","jimfromchinagd@gmail.com") //邮件标题 邮件内容 需要发送到的邮箱地址
}

