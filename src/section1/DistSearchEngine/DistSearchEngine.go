package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Email string
	Name  string
}

var DataBase = []User{
	{Email: "john.doe@example.com", Name: "John Doe"},
	{Email: "jane.smith@example.com", Name: "Jane Smith"},
	{Email: "harper.main@example.com", Name: "Harper Main"},
	{Email: "michael.johnson@example.com", Name: "Michael Johnson"},
	{Email: "emily.davis@example.com", Name: "Emily Davis"},
	{Email: "william.brown@example.com", Name: "William Brown"},
	{Email: "olivia.jones@example.com", Name: "Olivia Jones"},
	{Email: "james.garcia@example.com", Name: "James Garcia"},
	{Email: "isabella.martinez@example.com", Name: "Isabella Martinez"},
	{Email: "benjamin.rodriguez@example.com", Name: "Benjamin Rodriguez"},
	{Email: "sophia.hernandez@example.com", Name: "Sophia Hernandez"},
	{Email: "alexander.lopez@example.com", Name: "Alexander Lopez"},
	{Email: "mia.gonzalez@example.com", Name: "Mia Gonzalez"},
	{Email: "daniel.wilson@example.com", Name: "Daniel Wilson"},
	{Email: "amelia.anderson@example.com", Name: "Amelia Anderson"},
	{Email: "henry.thomas@example.com", Name: "Henry Thomas"},
	{Email: "ava.taylor@example.com", Name: "Ava Taylor"},
	{Email: "jackson.moore@example.com", Name: "Jackson Moore"},
	{Email: "charlotte.jackson@example.com", Name: "Charlotte Jackson"},
	{Email: "sebastian.martin@example.com", Name: "Sebastian Martin"},
	{Email: "harper.lee@example.com", Name: "Harper Lee"},
}

type Worker struct {
	users []User
	ch    chan *User
}

func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	for _, user := range w.users {
		if strings.Contains(user.Email, email) {
			w.ch <- &user
		}
	}
}

func main() {
	email := os.Args[1] // go run DistSearchEngine.go invalid@email

	ch := make(chan *User)

	log.Printf("Searching for user with email: %s", email)
	go NewWorker(DataBase[:5], ch).Find(email)   // Non-overlapping range
	go NewWorker(DataBase[5:10], ch).Find(email) // Non-overlapping range
	go NewWorker(DataBase[10:15], ch).Find(email) // Non-overlapping range
	go NewWorker(DataBase[15:], ch).Find(email) // Non-overlapping range


	for {
		select {
		case user := <-ch:
			log.Printf("User found: %s", user.Name)
		case <-time.After(1 * time.Second):
			return
		}
	}
}
