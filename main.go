package main

import (
	"fmt"
	"game-app/entity"
	"game-app/repository/mysql"
)

func main() {

}
func testUserMysql() {
	mysqlRepo := mysql.New()
	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "09382715999",
		Name:        "hosein",
	})
	if err != nil {
		fmt.Println("err register user: ", err)
	} else {
		fmt.Println("createdUser user", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique error", err)
	}
	fmt.Println("is phone number uniques ", isUnique)
}
