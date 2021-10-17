package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	example1()
	example2()
	example3()
}

func example1() {
	userJson := `{"username": "jhon", "email": "jhon@email.com", "password": "123456"}`

	fmt.Println(userJson)

	var userObj User

	err := json.Unmarshal([]byte(userJson), &userObj)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userObj)
}

func example2() {
	userObj := User{
		Username: "jhon",
		Email:    "jhon@email.com",
		Password: "123456",
	}

	fmt.Println(userObj)

	userJson, err := json.Marshal(userObj)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", userJson)
}

func example3() {
	type Weapon struct {
		Name  string `json:"weapon_name"`
		Level int    `json:"weapon_level"`
	}

	type Character struct {
		Name   string `json:"name"`
		Class  string `json:"class"`
		Level  int    `json:"level"`
		Weapon Weapon `json:"weapon"`
	}

	userObj := Character{
		Name:  "ring0",
		Class: "assassin",
		Level: 50,
		Weapon: Weapon{
			Name:  "blade of chaos",
			Level: 2,
		},
	}

	fmt.Println(userObj)

	userJson, err := json.Marshal(userObj)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", userJson)
}
