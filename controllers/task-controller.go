package controllers

import (
	"net/http"

	"github.com/EmmaUmeh/Goka-CMS/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateUserTask handles the creation of a new task.
func CreateUserTask(c *gin.Context, db *gorm.DB) {
	var task models.Task

	// Parse the incoming request body into task
	if err := c.ShouldBindJSON(&task); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the task data
	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if task.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body is required"})
		return
	}

	// Attempt to create the task in the database
	if err := CreateTask(db, &task); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Task"})
		return
	}

	// If the task is successfully created, return a success response
	taskArray := []models.Task{task}
	c.JSON(http.StatusOK, gin.H{
		"Task created successfully": taskArray,
	})
}

func GetTaskByID(c *gin.Context, db *gorm.DB, id string) {
	var task models.Task

	// Attempt to find the task in the database
	if err := db.First(&task, id).Error; err!= nil {
		// If the task is not found, return a 404 Not Found status
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := db.Preload("User").Where("id =?", id).First(&task).Error; err!= nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// If the task is found, return the task data
	c.JSON(http.StatusOK, gin.H{
		"Task": task,
	})
}

// CreateTask attempts to create a new task in the database.
func CreateTask(db *gorm.DB, task *models.Task) error {
	return db.Create(task).Error
}

// ListTaskById retrieves a task by its ID.
func ListTaskById(db *gorm.DB, task_dataById *models.Task) error {
	return db.Find(task_dataById).Error
}
