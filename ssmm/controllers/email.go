package controllers
import (
	"fmt"
	"net/smtp"
	"strings"
	"strconv"
)

var Emailuser, Emailpass, Smtp, Domain string

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "@gfw404.com\r\nSubject: " +subject+ "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

//welcome
func SendWelcome(mail, username, password string) bool {
	body := `Dear `+username+`,
Thank you for signing up with us. Your new account has been setup and you can now login to our client area using the details below.

Email Address: `+mail+`
Password: `+password+`

To login, visit http://www.gfw404.com/user

Kindest Regards,
GFW404
www.gfw404.com
		`
	err := SendToMail(Emailuser, Emailpass, Smtp, mail, "Welcome", body, "plain")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Send mail success!")
		return true
	}
}

//account detail
func SendAccountDetail(mail, username, ip string, port int, pass, authtype string) bool {
	body := `Dear `+username+`,
We are pleased to tell you that the service you ordered has now been set up and is operational.

New Service Information
=======================

IP Address: `+ip+`
Port: `+strconv.Itoa(port)+`
Password: `+pass+`
Encryption method: `+authtype+`

To login, visit http://www.gfw404.com/user

Kindest Regards,
GFW404
www.gfw404.com
		`
	err := SendToMail(Emailuser, Emailpass, Smtp, mail, "New Service", body, "plain")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Send mail success!")
		return true
	}
}

func SendBillInfo(mail string, username string, bid string, price float64, createtime, expiretime string) bool {
	body := `Dear `+username+`,
This is a billing reminder that your invoice no. `+bid+` which was generated on `+createtime+` is due on `+expiretime+`.

Invoice: `+bid+`
Amount Due: ￥`+strconv.FormatFloat(price, 'f', 2, 64) +`RMB
Due Date: `+expiretime+`

To login, visit http://www.gfw404.com/user

Kindest Regards,
GFW404
www.gfw404.com
		`
	err := SendToMail(Emailuser, Emailpass, Smtp, mail, "Invoice Payment Reminder", body, "plain")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Send mail success!")
		return true
	}
}

func SendBillComfirm(mail string, username string, bid string, price float64) bool {
	body := `Dear `+username+`,
This is a payment receipt for Invoice `+bid+`.

Invoice: `+bid+`
Amount Due: ￥`+strconv.FormatFloat(price, 'f', 2, 64) +`RMB

Note: This email will serve as an official receipt for this payment.

To login, visit http://www.gfw404.com/user

Kindest Regards,
GFW404
www.gfw404.com
		`
	err := SendToMail(Emailuser, Emailpass, Smtp, mail, "Invoice Payment Confirmation", body, "plain")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Send mail success!")
		return true
	}
}

func main() {
	user := "admin"
	password := "19901124ubuntu"
	host := "mail.gfw404.com:25"
	to := "799770174@qq.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}