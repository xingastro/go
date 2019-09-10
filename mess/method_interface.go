package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Println("Sending user email to", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	u.notify()
	sendNotification(&u)
}
