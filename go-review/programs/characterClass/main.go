package characterclass

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	var choice int

	fmt.Println("Welcome to Class Selector:")
	for {
		fmt.Println("1. List all classes")
		fmt.Println("2. Choose A Class")
		fmt.Println("3. Create your custom class")

		fmt.Println("Enter your choice")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println()
			printClasses()
		case 2:
			ChooseAClass()
		case 3:
			createAClass()
		}
	}
}

var classes = map[string][]string{
	"knight": {"100", "sword", "strength"},
	"mage":   {"60", "staff", "intelligence"},
	"rogue":  {"80", "knives", "dexterity"},
}

func ChooseAClass() {
	var orderedClasses []string
	for class := range classes {
		orderedClasses = append(orderedClasses, class)
	}

	slices.Sort(orderedClasses)

	for i, class := range orderedClasses {
		fmt.Printf("%d. %s: %v\n", i+1, class, classes[class])
	}

	fmt.Println("Choose Your class ")

	var choice int
	fmt.Scanln(&choice)
	playerClass := orderedClasses[choice-1]
	playerHealth := classes[playerClass][0]
	playerWeapon := classes[playerClass][1]
	playerAttribute := classes[playerClass][2]

	fmt.Printf("\n\n\n")
	fmt.Printf("You chose %s!\nYour health is: %s\nYour weapon: %s\nYour attribute: %s", playerClass, playerHealth,
		playerWeapon, playerAttribute)
	fmt.Printf("\n\n\n")
	os.Exit(1)
}

func printClasses() {

	for class, details := range classes {
		fmt.Printf("%s: %v\n", class, details)
	}
	fmt.Println()
}

func createAClass() {
	var className string
	var health string
	var weapon string
	var attribute string

	fmt.Printf("\n\nname Your class:")
	fmt.Scanln(&className)
	_, ok := classes[className]
	if ok {
		fmt.Println("class already exists")
		return
	}

	fmt.Println("How much health does your class name: ")
	fmt.Scanln(&health)
	fmt.Println("Whats your weapon:")
	fmt.Scanln(&weapon)
	fmt.Println("Whats your attribute: ")
	fmt.Scanln(&attribute)

	classes[className] = []string{health, weapon, attribute}
}
