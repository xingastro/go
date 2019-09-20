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

type admin struct {
	name string
	email string
}

func (a admin) notify() {
	fmt.Println("Sending admin email to", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	u.notify()
	u1 := &user{"Adam", "Adam@email.com"}
	u1.notify()

	a := admin{"Bill", "bill@email.com"}
	a.notify()
	a2 := &admin{"Adam", "Adam@email.com"}
	a2.notify()

	var iu notifier = &user{"Bill", "bill@email.com"}
	iu.notify()

	var ia notifier = admin{"Bill", "bill@email.com"}
	ia.notify()
}
