package main

import (
	"fmt"
	"time"

	"github.com/BryanMwangi/golang-write-to-local-json-example/DB/Models"
)

var NewUser = Models.User{
	Id:             "3134-434-55335",
	Email:          "jane@test.com",
	FirstName:      "Sophie",
	LastName:       "Doe",
	Password:       "34567",
	PhoneNumber:    "+3676453423",
	Deleted:        false,
	CreatedAt:      time.Now().Local().String(),
	UpdatedAt:      time.Now().Local().String(),
	ProfilePicture: "https://cdn.example.com/profilePictures/JaneDoe.jpg",
}

func main() {
	_, err := Models.Create(NewUser)
	if err != nil {
		fmt.Println(err)
	}

	user, err := Models.GetByEmail("test2@test.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

	users, err := Models.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)

	newUser, err := Models.Update(NewUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newUser)

	_, err = Models.Delete("3134-434-55335")
	if err != nil {
		fmt.Println(err)
		return
	}

}
