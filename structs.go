package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("new user validation failed")
	}

	return &user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
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

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)

		return
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scanln(&value)

	return value
}
