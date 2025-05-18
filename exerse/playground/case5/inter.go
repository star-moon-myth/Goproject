package main

import (
	"errors"
	"fmt"
	"strings"
)

type Notifier interface {
	send(string) error
}

type EmailNotify struct {
	emailaddress string
}

func (e EmailNotify) send(message string) error {
	if e.emailaddress== "" {
		return fmt.Errorf("无效邮箱地址: %s", e.emailaddress)
	}
	
	if !strings.Contains(e.emailaddress,"@"){
		return errors.New("邮箱格式无效")
	}
	fmt.Printf("邮箱发送到%s:%s\n",e.emailaddress,message)
	return nil 

}
type SmsNotify struct{
	phonenumber string
}
func (p SmsNotify) send(message string ) error{
	fmt.Printf("短信发送到%s：%s\n",p.phonenumber,message)
	return nil
}
func usernotify(f Notifier,message string){
	fmt.Println("Ready to start:")
	err:=f.send(message)
	if err!=nil{
		fmt.Println("failed:",err)
		return 
	}
	fmt.Println("It is Ok🌸")
}
func main(){
	email:=EmailNotify{emailaddress: ""}
	number:=SmsNotify{phonenumber: "16605692861"}
	usernotify(email,"这是邮件📧发送机制")
	usernotify(number,"这是短信发送机制")
}