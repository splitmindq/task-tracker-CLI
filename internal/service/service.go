package service

import (
	"task-tracker/internal/entities"
	"task-tracker/internal/storage"
	"time"
)

func NewTask(description string) (*entities.Task, error) {
	task := &entities.Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.DateTime),
		UpdatedAt:   time.Now().Format(time.DateTime),
		IsPinned:    false,
	}

	if err := storage.AddTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func MarkTaskAsDone(id int) error {
	task, err := storage.GetTaskByID(id)
	if err != nil {
		return err
	}
	task.Status = "done"
	return storage.UpdateTask(&task)
}

func MarkTaskAsInProgress(id int) error {
	task, err := storage.GetTaskByID(id)
	if err != nil {
		return err
	}
	task.Status = "in-progress"
	return storage.UpdateTask(&task)
}

func DeleteTask(id int) error {
	return storage.DeleteTask(id)
}

func GetAllTasks() {
	storage.GetAllTasks()
}

func UpdateTask(id int, description string) error {
	task, err := storage.GetTaskByID(id)
	if err != nil {
		return err
	}
	task.Description = description
	task.UpdatedAt = time.Now().Format(time.DateTime)
	return storage.UpdateTask(&task)
}

func PinTask(id int) error {
	task, err := storage.GetTaskByID(id)
	if err != nil {
		return err
	}
	task.IsPinned = true
	return storage.UpdateTask(&task)
}

func UnpinTask(id int) error {
	task, err := storage.GetTaskByID(id)
	if err != nil {
		return err
	}
	task.IsPinned = false
	return storage.UpdateTask(&task)
}

func GetPinnedTasks() {
	storage.GetPinnedTasks()
}
