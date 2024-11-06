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

func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}
}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(userFirstName, userLastName, userBirthdate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scan(&value)

	return value
}
