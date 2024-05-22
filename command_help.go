package main

import "fmt"

func cmdHelp(b []byte) string {
	commands := getCommands()
	helpText := ""

	//TODO: Add header information including version

	for _, command := range commands {
		helpText += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}

	return helpText
}
