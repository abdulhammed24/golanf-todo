package controllers

import (
	"net/http"
	"todo-app/database"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// CreateTodo handles creating a new Todo
func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }

    c.JSON(http.StatusCreated, todo)
}


// GetTodos fetches all Todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// // GetTodo fetches a single Todo by ID
func GetTodoByID(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

// UpdateTodo updates a Todo by ID
func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a Todo by ID
func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
