package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//inheretence or parent class type thing not available in golang
	human := User{"sanskar", "sanskarsharma3110@gmail.com", true, 21}
	fmt.Println(human)
	fmt.Printf("naam toh sunna hoga %v umer %v\n", human.Name, human.Age)
	if sanskar := human.Age; sanskar == 21 {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}

	//switch case
	rand.Seed(time.Now().UnixNano())
	diceNo := rand.Intn(6) + 1
	switch diceNo {
	case 1:
		fmt.Println("its a 1")
	case 2:
		fmt.Println("its a 2")
	case 3:
		fmt.Println("its a 3")
		fallthrough
	case 4:
		fmt.Println("its a 4")
	case 5:
		fmt.Println("its a 5")
	case 6:
		fmt.Println("its a 6")
	default:
		fmt.Println("sahi number daalay mandbudhi")
	}
	human.GetStatus()
	fmt.Println("Email of the user is currently : ", human.Email)
	human.NewMail()
	fmt.Println("Email of the user is currently : ", human.Email)
}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

type Values struct {
	X int
	Y int
}

func (u User) GetStatus() {
	fmt.Println("The status of the user is :", u.Verified)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is updated to :", u.Email)
}
