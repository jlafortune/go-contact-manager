// contacts.go
// Basic contact manager
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var contacts = []contact{}

type contact struct {
	firstName string
	lastName  string
	age       int
	telephone string
}

func main() {
	mainMenu()
}

func mainMenu() {
	fmt.Println("Welcome to Contact Manager. Select an option:")
	fmt.Println("\t[1] Find a Contact")
	fmt.Println("\t[2] Create a Contact")
	fmt.Println("\t[q] Quit")
	fmt.Print("Your choice: ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		switch {
		case input == "1":
			findContact()
		case input == "2":
			createContact()
		case input == "q" || input == "Q":
			os.Exit(0)
		default:
			fmt.Println("Unrecognized selection.")
			mainMenu()
		}
	}
}

func findContact() {
	for {
		for i, c := range contacts {
			fmt.Printf("\t[%d] %s %s\n", i, c.firstName, c.lastName)
		}
		fmt.Println("\t[m] Main Menu")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "m" {
			mainMenu()
		} else {
			i, err := strconv.Atoi(choice)
			if err != nil {
				fmt.Println("Invalid selection.")
				continue
			}
			editContact(&contacts[i])
		}
	}
}

func editContact(c *contact) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("First name [%s]: ", c.firstName)
	scanner.Scan()
	firstName := scanner.Text()
	if firstName != "" && firstName != c.firstName {
		c.firstName = firstName
	}

	fmt.Printf("Last name [%s]: ", c.lastName)
	scanner.Scan()
	lastName := scanner.Text()
	if lastName != "" && lastName != c.lastName {
		c.lastName = lastName
	}

	for {
		fmt.Printf("Age [%d]: ", c.age)
		scanner.Scan()
		age := scanner.Text()
		if age != "" && age != strconv.Itoa(c.age) {
			ageNum, err := strconv.Atoi(age)
			if err != nil {
				fmt.Println("Invalid age.")
				continue
			} else {
				c.age = ageNum
				break
			}
		} else {
			break
		}
	}

	fmt.Printf("Telephone [%s]: ", c.telephone)
	scanner.Scan()
	telephone := scanner.Text()
	if telephone != "" && telephone != c.telephone {
		c.telephone = telephone
	}

	fmt.Println("Contact saved.")
}

func createContact() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("First name: ")
	scanner.Scan()
	firstName := scanner.Text()

	fmt.Print("Last name: ")
	scanner.Scan()
	lastName := scanner.Text()

	fmt.Print("Age: ")
	scanner.Scan()
	age := scanner.Text()

	fmt.Print("Telephone: ")
	scanner.Scan()
	telephone := scanner.Text()

	ageNum, _ := strconv.Atoi(age)
	newContact := contact{firstName, lastName, ageNum, telephone}
	contacts = append(contacts, newContact)

	fmt.Println("Contact saved.")
	mainMenu()
}
