package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type task struct {
	gorm.Model
	decription string
	status     string
	createdAt  time.Time
	updatedAt  time.Time
}

func addTask(description string) error {
	return nil
}

func deleteTask(id string) error {
	return nil
}

func updateTask(id, description string) error {
	return nil
}

func listTasks() error {
	return nil
}

func listTaskStatus(status string) error {
	return nil
}

func markInProgress(id string) error {
	return nil
}

func markDone(id string) error {
	return nil
}
