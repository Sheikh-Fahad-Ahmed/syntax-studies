package main

import "fmt"

func main() {

}

var classes = map[string][]string{
	"knight": {"100", "sword", "strength"},
	"mage":   {"60", "staff", "intelligence"},
	"rogue":  {"80", "knives", "dexterity"},
}

func ChooseAClass() {
	fmt.Println("Welcome to Class Selector")

	printClasses()

	fmt.Println("Choose Your class OR Create one")

	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("You choose knight! ")
	}
}

func printClasses() {
	i := 1
	for class, details := range classes {
		fmt.Printf("%d. %s: %v\n", i, class, details)
		i += 1
	}
}
