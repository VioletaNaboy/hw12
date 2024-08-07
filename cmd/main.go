package main

import (
	"bufio"
	"fmt"
	"main/internal/password"
	"os"
	"strings"
)

func main() {
	mgr := password.NewManager("passwords.json")

	for {
		fmt.Println("Menu:")
		fmt.Println("1. List passwords")
		fmt.Println("2. Save password")
		fmt.Println("3. Retrieve password")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			mgr.ListPasswords()
		case "2":
			savePasswordFlow(mgr)
		case "3":
			retrievePasswordFlow(mgr)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func savePasswordFlow(mgr *password.Manager) {
	fmt.Print("Enter the name of the password: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter the password: ")
	pass, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	pass = strings.TrimSpace(pass)

	mgr.SavePassword(name, pass)
	fmt.Println("Password saved successfully.")
}

func retrievePasswordFlow(mgr *password.Manager) {
	fmt.Print("Enter the name of the password you want to retrieve: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.TrimSpace(name)

	mgr.RetrievePassword(name)
}
