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
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// appUser := new(user) // creates a pointer
	// appUser.firstName = firstName
	// appUser.lastName = lastName
	// appUser.birthdate = birthdate
	// appUser.createdAt = time.Now()

	outputUserDetails(appUser)
}

func outputUserDetails(user user) {
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scan(&value)

	return value
}
