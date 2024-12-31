package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/todos", controllers.GetTodos)
    r.GET("/todos/:id", controllers.GetTodoByID)
    r.POST("/todos", controllers.CreateTodo)
    r.PUT("/todos/:id", controllers.UpdateTodo)
    r.DELETE("/todos/:id", controllers.DeleteTodo)
}
