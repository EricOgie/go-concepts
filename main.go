package main

import (
	"fmt"
	"strings"
)

// Define a function with double return value
func getInitials(fullname string) (string, string) {

	var initials []string
	capitalizeName := strings.ToUpper(fullname)

	names := strings.Split(capitalizeName, " ")

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {

	firstInitial1, surnameInitials1 := getInitials("SAM OGIE")
	firstInitial2, surnameInitials2 := getInitials("GRAHAM Samson James")
	firstInitial3, surnameInitials3 := getInitials("James")

	fmt.Println(firstInitial1, surnameInitials1)
	fmt.Println(firstInitial2, surnameInitials2)
	fmt.Println(firstInitial3, surnameInitials3)

}
