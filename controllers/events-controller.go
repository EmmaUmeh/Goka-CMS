package controllers

import (
	"net/http"

	"github.com/EmmaUmeh/Goka-CMS/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateUserTask handles the creation of a new task.
func CreateUserEvent(c *gin.Context, db *gorm.DB) {
	var event models.Events

	// Parse the incoming request body into task
	if err := c.ShouldBindJSON(&event); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the task data
	if event.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if event.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	// Attempt to create the task in the database
	if err := CreateEvent(db, &event); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Event"})
		return
	}

	// If the task is successfully created, return a success response
	taskArray := []models.Events{event}
	c.JSON(http.StatusOK, gin.H{
		"Task created successfully": taskArray,
	})
}

func GetEventByID(c *gin.Context, db *gorm.DB, id string) {
	var event models.Events

	// Attempt to find the event in the database
	if err := db.First(&event, id).Error; err!= nil {
		// If the event is not found, return a 404 Not Found status
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := db.Preload("User").Where("id =?", id).First(&event).Error; err!= nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// If the task is found, return the task data
	c.JSON(http.StatusOK, gin.H{
		"Task": event,
	})
}

// CreateTask attempts to create a new task in the database.
func CreateEvent(db *gorm.DB, task *models.Events) error {
	return db.Create(task).Error
}

// ListTaskById retrieves a task by its ID.
// func ListTaskById(db *gorm.DB, task_dataById *models.Task) error {
// 	return db.Find(task_dataById).Error
// }
