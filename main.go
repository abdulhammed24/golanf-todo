package main

import (
	"todo-app/database"
	"todo-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    database.Init() 

    r := gin.Default() 

    r.Use(cors.Default())

    routes.SetupRoutes(r) 

    r.Run(":8080") 
}
