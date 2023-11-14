package routers

import (
	"finalProject1/controllers"

	"github.com/gin-gonic/gin"

	_ "finalProject1/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Todo API
// @version 1.0
// @description This is a sevice for managing todos
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func Router() {
	router := gin.Default()

	controllers := controllers.NewHandlerImpl()

	router.GET("/todo/:id", controllers.GetOneTodo)

	router.POST("/todo", controllers.CreateTodo)

	router.GET("/todo", controllers.GetAllTodo)

	router.PATCH("/todo/:id", controllers.UpdateTodo)

	router.DELETE("/todo/:id", controllers.DeleteTodo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
