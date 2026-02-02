package main

import (
	"fmt"
	"os"
	"todo-list/functions"
)

func main() {

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: projectName add \"task description\"")
			os.Exit(1)
		}

		task := os.Args[2]
		err := functions.AddTask(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}
		fmt.Printf("Added task: %s\n", task)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: projectName update <task_id> \"new task description\"")
			os.Exit(1)
		}

		taskID := os.Args[2]
		newDescription := os.Args[3]

		err := functions.UpdateTask(taskID, newDescription, "")
		if err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}
		fmt.Printf("Updated task %s to: %s\n", taskID, newDescription)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: projectName delete <task_id>")
			os.Exit(1)
		}
		taskID := os.Args[2]
		// Here you could delete the task from a file, a database, etc.
		err := functions.DeleteTaskById(taskID)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted task with ID: %s\n", taskID)

	case "list":
		var option string
		if len(os.Args) > 2 {
			option = os.Args[2]
		} else {
			option = ""
		}

		var tasks []functions.Task
		var err error

		switch option {
		case "done":
			tasks, err = functions.GetTasksByStatus(functions.StatusDone)
		case "in-progress":
			tasks, err = functions.GetTasksByStatus(functions.StatusInProgress)
		case "todo":
			tasks, err = functions.GetTasksByStatus(functions.StatusTodo)
		default:
			tasks, err = functions.GetTasks()
		}

		if err != nil {
			fmt.Println("Error retrieving tasks:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			os.Exit(0)
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: projectName mark-done <task_id>")
			os.Exit(1)
		}
		taskID := os.Args[2]
		// Here you could mark the task as done in a file, a database, etc.
		err := functions.UpdateTask(taskID, "", "done")
		if err != nil {
			fmt.Println("Error marking task as done:", err)
			os.Exit(1)
		}
		fmt.Printf("Marked task %s as done\n", taskID)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: projectName mark-in-progress <task_id>")
			os.Exit(1)
		}
		taskID := os.Args[2]
		// Here you could mark the task as in-progress in a file, a database, etc.
		err := functions.UpdateTask(taskID, "", "in-progress")
		if err != nil {
			fmt.Println("Error marking task as in-progress:", err)
			os.Exit(1)
		}
		fmt.Printf("Marked task %s as in-progress\n", taskID)
	default:
		fmt.Println("Unknown command:", command)
	}
	fmt.Println("Command executed successfully.")
	os.Exit(0)
}
