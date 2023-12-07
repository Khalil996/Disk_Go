package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestEmailTest(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <18559680907@163.com>"
	e.To = []string{"2649475267@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码是<h1>123456</h1>")
	err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "18559680907@163.com", "GHHMBKGPETJVEDIQ", "smtp.163.com"))
	//返回eof时，关闭ssl
	err = e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "18559680907@163.com", "GHHMBKGPETJVEDIQ", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
