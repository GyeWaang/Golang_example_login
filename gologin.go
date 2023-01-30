package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var attempts int = 3
	var username string = "gyewong"
	var password string = "secret"

	for i := 0; i < attempts; i++ {
		fmt.Print("Enter username: ")
		reader := bufio.NewReader(os.Stdin)
		inputUsername, _ := reader.ReadString('\n')
		inputUsername = inputUsername[:len(inputUsername)-1]

		fmt.Print("Enter password: ")
		inputPassword, _ := reader.ReadString('\n')
		inputPassword = inputPassword[:len(inputPassword)-1]

		if inputUsername == username && inputPassword == password {
			fmt.Println("Login Successful")
			break
		} else {
			fmt.Println("Login Failed. Attempts left: ", attempts-i-1)
		}
	}
}
