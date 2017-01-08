// Basic CLI contact manager, which stores data in a local file contacts.json
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var contacts = []contact{}

type contact struct {
	FirstName string
	LastName  string
	Age       int
	Telephone string
}

func main() {
	loadContacts()
	mainMenu()
}

func loadContacts() {
	input, err := os.Open("contacts.json")
	defer input.Close()
	if err != nil {
		fmt.Println("No previous contacts found. Will create a new save.")
	} else {
		jsonParser := json.NewDecoder(input)
		if err = jsonParser.Decode(&contacts); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func mainMenu() {
	for {
		fmt.Println("Select an option:")
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
			}
		}
	}
}

func findContact() {
	for {
		for i, c := range contacts {
			fmt.Printf("\t[%d] %s %s, Age: %d, Phone: %s\n", i,
				c.FirstName, c.LastName, c.Age, c.Telephone)
		}
		fmt.Println("\t[m] Main Menu")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "m" {
			break
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

	fmt.Printf("First name [%s]: ", c.FirstName)
	scanner.Scan()
	firstName := scanner.Text()
	if firstName != "" && firstName != c.FirstName {
		c.FirstName = firstName
	}

	fmt.Printf("Last name [%s]: ", c.LastName)
	scanner.Scan()
	lastName := scanner.Text()
	if lastName != "" && lastName != c.LastName {
		c.LastName = lastName
	}

	// Executes in a loop until the user enters a number
	// or just hits enter, which indicates no change
	for {
		fmt.Printf("Age [%d]: ", c.Age)
		scanner.Scan()
		age := scanner.Text()
		if age != "" && age != strconv.Itoa(c.Age) {
			ageNum, err := strconv.Atoi(age)
			if err != nil {
				fmt.Println("Invalid age.")
				continue
			} else {
				c.Age = ageNum
				break
			}
		} else {
			break
		}
	}

	fmt.Printf("Telephone [%s]: ", c.Telephone)
	scanner.Scan()
	telephone := scanner.Text()
	if telephone != "" && telephone != c.Telephone {
		c.Telephone = telephone
	}

	save()
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

	save()
}

func save() {
	output, _ := json.Marshal(contacts)
	ioutil.WriteFile("contacts.json", output, 0644)
	fmt.Println("Changes saved.")
}
