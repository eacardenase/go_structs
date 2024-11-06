package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scan(&value)

	return value
}
