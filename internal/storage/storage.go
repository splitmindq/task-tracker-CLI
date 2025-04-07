package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"task-tracker/internal/entities"
	"task-tracker/internal/utils"
)

var (
	initOnce    sync.Once
	mu          sync.RWMutex
	storagePath = filepath.Join("data", "tasks.json")
	tasks       []entities.Task
)

func Initialize() error {
	var initErr error
	initOnce.Do(func() {
		if err := os.MkdirAll("data", 0755); err != nil {
			initErr = err
			return
		}

		if _, err := os.Stat(storagePath); os.IsNotExist(err) {
			tasks = []entities.Task{}
			return
		}

		file, err := os.Open(storagePath)
		if err != nil {
			initErr = err
			return
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&tasks); err != nil {
			initErr = err
		}
	})
	return initErr
}

func AddTask(task *entities.Task) error {
	mu.Lock()
	defer mu.Unlock()

	if len(tasks) > 0 {
		task.ID = tasks[len(tasks)-1].ID + 1
	} else {
		task.ID = 1
	}

	tasks = append(tasks, *task)
	return saveTasks()
}

func saveTasks() error {
	file, err := os.Create(storagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

func UpdateTask(task *entities.Task) error {
	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = *task
		}
	}
	return saveTasks()
}

func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return saveTasks()
}

func GetAllTasks() {
	mu.Lock()
	defer mu.Unlock()
	if tasks == nil {
		fmt.Println("tasks is empty")
	}

	table := utils.SetTableWriterOptions()

	var pinnedTasks []entities.Task
	var unpinnedTasks []entities.Task
	for _, task := range tasks {
		if task.IsPinned {
			pinnedTasks = append(pinnedTasks, task)
		} else {
			unpinnedTasks = append(unpinnedTasks, task)
		}
	}

	for _, task := range pinnedTasks {
		table.Append([]string{
			fmt.Sprintf("★ %d", task.ID),
			task.Description,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		})
	}

	for _, task := range unpinnedTasks {
		table.Append([]string{
			strconv.Itoa(task.ID),
			task.Description,
			task.Status,
			task.CreatedAt,
			task.UpdatedAt,
		})
	}

	table.Render()
}

func GetTaskByID(id int) (entities.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return entities.Task{}, errors.New("task not found")
}

func GetNotDoneTasks() {
	mu.Lock()
	defer mu.Unlock()
	if tasks == nil {
		fmt.Println("tasks is empty")
	}

	table := utils.SetTableWriterOptions()

	for _, task := range tasks {

		if task.Status == "todo" {
			table.Append([]string{
				strconv.Itoa(task.ID),
				task.Description,
				task.Status,
				task.CreatedAt,
			})
		}

	}
	table.Render()
}

func GetInProgressTasks() {
	mu.Lock()
	defer mu.Unlock()
	if tasks == nil {
		fmt.Println("tasks is empty")
	}

	table := utils.SetTableWriterOptions()

	for _, task := range tasks {

		if task.Status == "in-progress" {
			table.Append([]string{
				strconv.Itoa(task.ID),
				task.Description,
				task.Status,
				task.CreatedAt,
			})
		}

	}
	table.Render()
}

func GetDoneTasks() {
	mu.Lock()
	defer mu.Unlock()
	if tasks == nil {
		fmt.Println("tasks is empty")
	}

	table := utils.SetTableWriterOptions()

	for _, task := range tasks {

		if task.Status == "done" {
			table.Append([]string{
				strconv.Itoa(task.ID),
				task.Description,
				task.Status,
				task.CreatedAt,
			})
		}

	}

	table.Render()
}

func GetPinnedTasks() {

	mu.Lock()
	defer mu.Unlock()

	if tasks == nil {
		fmt.Println("tasks is empty")
	}

	table := utils.SetTableWriterOptions()

	for _, task := range tasks {
		if task.IsPinned {
			table.Append([]string{
				fmt.Sprintf("★ %d", task.ID),
				task.Description,
				task.Status,
				task.CreatedAt,
			})

		}
	}

	table.Render()

}
