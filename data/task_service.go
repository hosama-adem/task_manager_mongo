package data

import (
	"Task_manager/models"
	"errors"
)

var tasks = []models.Task{}

// To get All Tasks
func GetAllTasks() []models.Task {
	return tasks
}

// To get Task by ID
func GetTaskByID(id int) *models.Task {
	for i, _ := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

// To update Task by ID
func UpdateTaskByID(id int, updatedTask models.Task) *models.Task {
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			return &tasks[i]

		}

	}
	return nil
}

// To Delete Task by ID
func RemoveTaskByID(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false

}

// To Add New Task
func AddNewTask(task models.Task) (models.Task, error) {
	taskid := task.ID

	for _, t := range tasks {
		if t.ID == taskid {
			return models.Task{}, errors.New("Task with this ID already exists")
		}
	}

	tasks = append(tasks, task)
	return task, nil
}
