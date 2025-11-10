package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"errors"
)


func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var my_commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Prints out help",
		callback:    commandHelp,
	},
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() != false {
			my_string := scanner.Text()
			first_word := strings.Fields(strings.ToLower(my_string))[0]
			command, ok := my_commands[first_word]
			if ok {
				err := command.callback()
				if err != nil {
					fmt.Printf("Error running a command")
				}
			}
		}
	}

}

func cleanInput(text string) []string {
	final_list := []string(strings.Fields(text))
	return final_list
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
