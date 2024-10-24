package main

import "fmt"

// uppercase/lowercase also applies to structs
// uppercase fields are exported and can be used outside the package
// lowercase fields are private to the package
type User struct {
	Name  string
	Role  string
	Email string
	Age   int
}

func (u User) Salary() int {
	switch u.Role {
	case "admin":
		return 2000
	case "editor":
		return 1500

	}
	return 0
}

// needs to be a pointer to the struct
// otherwise it will not update the struct
// because u User is a copy of the struct
func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func main() {
	user := User{
		Name:  "quiyet",
		Role:  "admin",
		Email: "qb@email.com",
		Age:   27,
	}
	fmt.Println(user)
	fmt.Println(user.Name)
	user.Name = "quiyet2"
	fmt.Println(user.Name)
	fmt.Println(user.Salary())

	user2 := User{Name: "john", Role: "editor"}
	fmt.Println(user2)
	fmt.Println(user2.Name)
	fmt.Println(user2.Salary())
}
