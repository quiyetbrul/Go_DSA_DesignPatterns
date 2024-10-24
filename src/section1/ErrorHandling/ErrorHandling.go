package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Role  string
	Email string
	Age   int
}

func (u User) Salary() (int, error) {
	if u.Role == "" {
		return 0, errors.New("empty role")
	}
	switch u.Role {
	case "admin":
		return 2000, nil
	case "editor":
		return 1500, nil
	}
	return 0, errors.New(
		fmt.Sprintf("undefined role: %s", u.Role))
}

func main() {
	user := User{Name: "quiyet"}
	// can define a car but the scope is only within the if block
	if salary, err := user.Salary(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(salary)
	}
}
