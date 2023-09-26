package service

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"okbang/util"
	"time"
)

func SendMail(message, to string) {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", "schenglee@qq.com")
	//接收人
	m.SetHeader("To", to)
	//主题
	m.SetHeader("Subject", i.Home+"("+time.Now().Format("2006年01月")+")")
	//内容
	m.SetBody("text/html", message)
	//邮箱服务地址,端口,发送人,token
	do := gomail.NewDialer("smtp.qq.com", 587, "schenglee@qq.com", "tgeuhbipbrrieajc")
	err := do.DialAndSend(m)
	if err != nil {
		util.Logger.Error(err.Error())
		return
	}
	util.Logger.Info(fmt.Sprintf("%s send successed.", to))
}
