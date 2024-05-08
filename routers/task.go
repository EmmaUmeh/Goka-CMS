// routers/route.go
package routers

import (
	"github.com/EmmaUmeh/Goka-CMS/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/EmmaUmeh/Goka-CMS/middlewares"
)

func TaskRoutes(router *gin.Engine, db *gorm.DB) {
	// Create a sub-router for tasks with AuthMiddleware applied to the GET request
	taskRouter := router.Group("/tasks", middleware.AuthMiddleware())

	// Define your task POST route within the sub-router
	taskRouter.POST("/create", func(c *gin.Context) {
		controllers.CreateUserTask(c, db)
	})

	// Define your task GET route within the sub-router
	taskRouter.GET("/getTask/:id", func(c *gin.Context) {
		id := c.Param("id") 
		controllers.GetTaskByID(c, db, id)
	})
	// Add other task-related routes here if needed
}
