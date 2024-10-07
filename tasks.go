package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Description string
	Status      string
}

func addTask(description string) error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	if description == " " {
		return fmt.Errorf("No description given.")
	}

	db := DBConnection
	task := Task{
		Description: description,
		Status:      "todo",
	}
	db.Create(&task)
	fmt.Printf("TASK ADDED: \n\tTask:%v \n\tStatus:%v \n\tCreated at:%v \n\tUpdated at:%v \n", task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	return nil
}

func deleteTask(id string) error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	db := DBConnection
	var task Task

	db.First(&task, id)
	if task.Description == "" {
		return fmt.Errorf("No Task found for id %v", id)
	}
	db.Delete(&task)
	fmt.Printf("TASK DELETED: \n\tTask:%v \n\tStatus:%v \n", task.Description, task.Status)
	return nil
}

func updateTask(id, description string) error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	db := DBConnection
	var task Task

	db.First(&task, id)
	if task.Description == "" {
		return fmt.Errorf("No Task found for id %v", id)
	}
	task.Description = description
	task.UpdatedAt = time.Now()
	db.Model(&task).Update("description", description)
	db.Model(&task).Update("updatedAt", time.Now())
	fmt.Printf("TASK UPDATED: \n\tTask:%v \n\tStatus:%v \n\tCreated at:%v \n\tUpdated at:%v \n", task.Description, task.Status, task.CreatedAt, task.UpdatedAt)

	return nil
}

func listTasks() error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	db := DBConnection
	var tasks []Task

	db.Find(&tasks)
	if len(tasks) == 0 {
		return fmt.Errorf("No Tasks Available")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)

	fmt.Fprintln(writer, "ID\tTASK\tSTATUS\tCREATED AT\tUPDATED AT")
	for _, task := range tasks {
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}

	writer.Flush()

	return nil
}

func listTaskStatus(status string) error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	db := DBConnection
	var tasks []Task
	valid := false
	STATUS := []string{"todo", "in-progress", "done"}
	for _, s := range STATUS {
		if s == status {
			valid = true
		}
	}

	if !valid {
		return fmt.Errorf("Invalid argument: status must be one of 'todo', 'in-progress', or 'done'")
	}
	db.Where("status = ?", status).Find(&tasks)
	if len(tasks) == 0 {
		return fmt.Errorf("No Tasks Available")
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)

	fmt.Fprintln(writer, "ID\tTASK\tSTATUS\tCREATED AT\tUPDATED AT")
	for _, task := range tasks {
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}

	writer.Flush()

	return nil
}

func markInProgress(id string) error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	db := DBConnection
	var task Task

	db.First(&task, id)
	if task.Description == "" {
		return fmt.Errorf("No Task found for id %v", id)
	}
	task.Status = "in-progress"
	task.UpdatedAt = time.Now()
	db.Model(&task).Update("status", "in-progress")
	db.Model(&task).Update("updatedAt", time.Now())
	fmt.Printf("TASK UPDATED: \n\tTask:%v \n\tStatus:%v \n\tCreated at:%v \n\tUpdated at:%v \n", task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	return nil
}

func markDone(id string) error {
	if DBConnection == nil {
		return fmt.Errorf("No database connection")
	}

	db := DBConnection
	var task Task

	db.First(&task, id)
	if task.Description == "" {
		return fmt.Errorf("No Task found for id %v", id)
	}
	task.Status = "done"
	task.UpdatedAt = time.Now()
	db.Model(&task).Update("status", "done")
	db.Model(&task).Update("updatedAt", time.Now())
	fmt.Printf("TASK UPDATED: \n\tTask:%v \n\tStatus:%v \n\tCreated at:%v \n\tUpdated at:%v \n", task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	return nil
}
