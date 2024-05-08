// // routers/route.go
// package routers

// import (

// 	"github.com/EmmaUmeh/Goka-CMS/controllers"
//     "github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	// "github.com/gofiber/fiber"
// )

// func SetupRoutes(router *gin.Engine, db *gorm.DB) {
//     router.POST("/auth/signup", controllers.Signup)
//     // Add other routes here if needed
// }

// routers/route.go

package routers

import (
	"github.com/EmmaUmeh/Goka-CMS/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TaskRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/create/task", func(c *gin.Context) {
		controllers.CreateUserTask(c, db)
	})
	// Add other routes here if needed
}
