package main

import (
	"fmt"
)

func commandParser(commandName string, args []string) {
	switch commandName {
	case "add":
		if len(args) == 3 {
			err := addTask(args[2])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Printf("Error: Invalid arguments")
		}
	case "delete":
		if len(args) == 3 {
			err := deleteTask(args[2])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Print("Error: Invalid arguments")
		}

	case "update":
		if len(args) == 4 {
			err := updateTask(args[2], args[3])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Print("Error: Invalid arguments")
		}

	case "list":
		if len(args) == 2 {
			err := listTasks()
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else if len(args) == 3 {
			err := listTaskStatus(args[2])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Print("Error: Invalid arguments")
		}

	case "mark-in-progress":
		if len(args) == 3 {
			err := markInProgress(args[2])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Print("Error: Invalid arguments")
		}
	case "mark-done":
		if len(args) == 3 {
			err := markDone(args[2])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Print("Error: Invalid arguments")
		}
	}
}
