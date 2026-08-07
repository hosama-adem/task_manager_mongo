package controllers

import (
	"Task_manager/data"
	"Task_manager/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// To Get All Tasks
func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// To Get Task by ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := data.GetTaskByID(idInt)
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"task": task})
	}

}

// To Update Task by ID
func UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	task := data.UpdateTaskByID(idInt, updatedTask)
	if task == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Task not found")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// To Remove Task by ID
func RemoveTaskByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := data.RemoveTaskByID(idInt)
	if !data {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Task not found")})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed successfully"})
}

// To add new Task
func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newtask, err := data.AddNewTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task": newtask})
}
