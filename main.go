package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		commandHelp()
	} else if len(args) > 1 {
		command, exists := getCommands()[args[1]]
		if exists {
			commandParser(command.name, args)
		} else {
			fmt.Printf("Error: '%s' is an Unknown command\nUse `task-cli help` to know all the available commands.", args[1])
		}
	}
}
