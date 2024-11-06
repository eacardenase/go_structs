package main

import (
	"fmt"

	"github.com/eacardenase/go_structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)

		return
	}

	admin := user.NewAdmin("Edwin", "Cardenas", "03/20/1996", "eacardenase@gmail.com", "superpassword")
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	var value string

	fmt.Print(promptText)
	fmt.Scanln(&value)

	return value
}
