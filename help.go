package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"add": {
			name:        "add",
			description: "Add a Task to the Todo list.\n\t`task-cli add <content-to-add>`",
		},
		"update": {
			name:        "update",
			description: "Update an existing Task in the Todo list.\n\t`task-cli add <id> <content-to-update>`",
		},
		"delete": {
			name:        "delete",
			description: "Delete an existing Task from the Todo list.\n\t`task-cli delete <id>`",
		},
		"mark-in-progress": {
			name:        "mark-in-progress",
			description: "Mark a Task 'in-progress' in the Todo list.\n\t`task-cli mark-in-progress <id>`",
		},
		"mark-done": {
			name:        "mark-done",
			description: "Mark a Task 'Completed' in the Todo list.\n\t`task-cli mark-done <id>`",
		},
		"list": {
			name:        "list",
			description: "List the tasks in the Todo list.\n\t`task-cli list`\n\t`task-cli delete <status>`",
		},
	}
}

func commandHelp() error {
	for _, cmd := range getCommands() {
		fmt.Printf("%s:\n \t%s \n", cmd.name, cmd.description)
	}
	return nil
}
