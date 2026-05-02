package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: task-cli <command> [args]")
		fmt.Fprintln(os.Stderr, "Commands: add, list, done, delete")
		os.Exit(1)
	}

	tasks, err := load()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading tasks: ", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task-cli add \"task title\"")
			os.Exit(1)
		}
		addTask(&tasks, os.Args[2])

	case "list":
		listTasks(tasks)

	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task-cli done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid ID", err)
			os.Exit(1)
		}
		markDone(&tasks, id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task-cli done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid ID", err)
			os.Exit(1)
		}
		deleteTask(&tasks, id)
	}
}

func addTask(tasks *TaskList, title string) {
	task := Task{
		ID:    tasks.NextID(),
		Title: title,
	}
	*tasks = append(*tasks, task)
	if err := save(*tasks); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving task", err)
		os.Exit(1)
	}
}

func listTasks(tasks TaskList) {
	if len(tasks) == 0 {
		fmt.Println("No Tasks yet, Add one with: task-cli add <task name>")
		return
	}
	for _, task := range tasks {
		status := "[]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%s #%d: %s\n", status, task.ID, task.Title)
	}
}

func markDone(tasks *TaskList, id int) {
	idx := tasks.FindIndex(id)
	if idx == -1 {
		fmt.Fprintf(os.Stderr, "No task found with ID %d\n", id)
		os.Exit(1)
	}
	(*tasks)[idx].Done = true
	if err := save(*tasks); err != nil {
		fmt.Fprintln(os.Stderr, "Error while saving task", err)
		os.Exit(1)
	}
	fmt.Printf("Task #%d marked as done.\n", id)
}

func deleteTask(tasks *TaskList, id int) {
	idx := tasks.FindIndex(id)
	if idx == -1 {
		fmt.Fprintf(os.Stderr, "No task found with ID %d\n", id)
		os.Exit(1)
	}
	*tasks = append((*tasks)[:idx], (*tasks)[idx+1:]...)
	if err := save(*tasks); err != nil {
		fmt.Fprintln(os.Stderr, "Error while saving task", err)
		os.Exit(1)
	}
	fmt.Printf("Task #%d deleted.\n", id)
}
