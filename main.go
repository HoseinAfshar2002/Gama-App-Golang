package main

import (
	"encoding/json"
	"fmt"
	"game-app/repository/mysql"
	"game-app/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	log.Println("Starting server on port 8080...")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("ok")
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error": invalid method"}`)
	}
	//read is body request user
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"err": "%v"`, err.Error())))
	}
	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"err": "%v"`, err.Error())))
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"err": "%v"}`, err.Error())))
		return
	}
	writer.Write([]byte(fmt.Sprintf(`{"success": "user created successfully""}`)))
}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message": "server is good"}`)
}

//func testUserMysql() {
//	mysqlRepo := mysql.New()
//	createdUser, err := mysqlRepo.RegisterUser(entity.User{
//		ID:          0,
//		PhoneNumber: "09382715991",
//		Name:        "hosein",
//	})
//	if err != nil {
//		fmt.Println("err register user: ", err)
//	} else {
//		fmt.Println("createdUser user", createdUser)
//	}
//
//	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
//	if err != nil {
//		fmt.Println("unique error", err)
//	}
//	fmt.Println("is phone number uniques ", isUnique)
//}
