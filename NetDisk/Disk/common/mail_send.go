package common

import (
	"cloud_go/Disk/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

//	func SendCode(toUserEmail, code string) error {
//		e := email.NewEmail()
//		e.From = "验证码获取 <18559680907@163.com>"
//		e.To = []string{toUserEmail}
//		e.Subject = "咪咪网盘"
//		e.HTML = []byte("你的验证码是<h1>" + code + "</h1>")
//
//		//返回eof时，关闭ssl
//		return e.SendWithTLS("smtp.163.com:465",
//			smtp.PlainAuth("", "18559680907@163.com", "IBXQLJWYKMYYHJOO", "smtp.163.com"),
//			&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
//	}
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "验证码获取 <2649475267@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "咪咪网盘"
	e.HTML = []byte("你的验证码是<h1>" + code + "</h1>")

	//返回eof时，关闭ssl
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "2649475267@qq.com", "fhtgncpegyjeecad", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
