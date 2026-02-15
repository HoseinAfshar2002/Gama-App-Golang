package phonenumber

import (
	"fmt"
	"strconv"
)

func IsPhoneNumberValid(phoneNumber string) bool {
	//TODO -
	if len(phoneNumber) != 11 {
		fmt.Println("is phone number invalid")
		return false
	}

	if phoneNumber[0:2] != "09" {
		fmt.Println("is phone number not starts with 09")
		return false
	}

	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		fmt.Println("is phone number invalid2")
		return false
	}

	return true

}
