package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var filename = "./functions/tasks.json"

func MarshalJSON(tasks []Task) ([]byte, error) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marching JSON %w", err)
	}
	return jsonData, nil
}

func ReadJSONFromFile() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading from file %s: %w", filename, err)
	}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			fmt.Println("Error unmarshaling existing JSON:", err)
			return nil, err
		}
	}
	return tasks, nil
}

func WriteJSONToFile(jsonData []byte) error {
	err := os.WriteFile(filename, []byte(jsonData), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filename, err)
	}
	return nil
}

func Count() (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("error reading from file %s: %w", filename, err)
	}

	var tasks []Task
	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return 0, fmt.Errorf("error unmarshaling JSON: %w", err)
		}
	}
	return len(tasks), nil
}

func AddTask(title string) error {
	count, err := Count()
	if err != nil {
		return fmt.Errorf("error counting tasks: %w", err)
	}
	newTask := Task{
		Id:          fmt.Sprintf("%d", count+1),
		Description: title,
		Status:      StatusTodo,
		CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z"),
	}

	tasks, err := ReadJSONFromFile()
	if err != nil {
		return fmt.Errorf("error reading from file: %w", err)
	}

	tasks = append(tasks, newTask)

	jsonData, err := MarshalJSON(tasks)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = WriteJSONToFile(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil

}

func UpdateTask(id string, description string, status Status) error {
	tasks, err := ReadJSONFromFile()
	if err != nil {
		return fmt.Errorf("error reading from file: %w", err)
	}

	if len(tasks) == 0 {
		return fmt.Errorf("no tasks found to update")
	}

	if id == "" {
		return fmt.Errorf("task ID cannot be empty")
	}
	if description == "" && status == "" {
		return fmt.Errorf("at least one of description or status must be provided for update")
	}

	for i, task := range tasks {
		if task.Id == id {
			if description != "" {
				tasks[i].Description = description
			}
			if status != "" {
				tasks[i].Status = status
			}
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02T15:04:05Z")
			break
		}
	}

	jsonData, err := MarshalJSON(tasks)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = WriteJSONToFile(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil

}

func GetTasks() ([]Task, error) {
	tasks, err := ReadJSONFromFile()
	if err != nil {
		return nil, fmt.Errorf("error reading tasks: %w", err)
	}
	return tasks, nil
}

func GetTaskById(id string) (*Task, error) {
	tasks, err := ReadJSONFromFile()
	if err != nil {
		return nil, fmt.Errorf("error reading tasks: %w", err)
	}

	for _, task := range tasks {
		if task.Id == id {
			return &task, nil
		}
	}

	return nil, fmt.Errorf("task with id %s not found", id)

}

func GetTasksByStatus(status Status) ([]Task, error) {
	tasks, err := ReadJSONFromFile()
	if err != nil {
		return nil, fmt.Errorf("error reading tasks: %w", err)
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil

}

func DeleteTaskById(id string) error {
	tasks, err := ReadJSONFromFile()
	if err != nil {
		return fmt.Errorf("error reading tasks: %w", err)
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if task.Id != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	jsonData, err := MarshalJSON(updatedTasks)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = WriteJSONToFile(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
