package main

import "fmt"

type Notifier interface {
	Send(messege string)
	GetName() string
}

type Email struct {
	Adress string
	Name   string
}

type Telegram struct {
	UserName string
	Name     string
}

type SMS struct {
	Phone string
	Name  string
}

func (e Email) Send(messege string) {
	fmt.Printf("Отправляем сообщение по Email %s: %s\n", e.Name, messege)
}

func (t Telegram) Send(messege string) {
	fmt.Printf("Отправляем сообщение в Telegram %s: %s\n", t.Name, messege)
}

func (s SMS) Send(messege string) {
	fmt.Printf("Отправляем сообщение по SMS %s: %s\n", s.Name, messege)
}

func (e Email) GetName() string {
	return e.Name
}

func (t Telegram) GetName() string {
	return t.Name
}

func (s SMS) GetName() string {
	return s.Name
}

func SendMes(n Notifier, messege string) {
	fmt.Printf("Уведомление для %s:\n", n.GetName())
	n.Send(messege)
}

func main() {
	email := Email{
		Adress: "malayegor@gmail.com",
		Name:   "Egor",
	}
	telega := Telegram{
		UserName: "yourusername",
		Name:     "Egor",
	}
	sms := SMS{
		Phone: "+79009009090",
		Name:  "Egor",
	}

	SendMes(email, "Привет!")
	SendMes(telega, "Привет!")
	SendMes(sms, "Привет!")
}
