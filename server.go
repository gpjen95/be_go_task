package main

import (
	"be_go_task/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "root page, server oke",
		})
	})

	V1 := router.Group("/v1")
	V1.GET("/task", controllers.GetTask)
	V1.GET("/task/:id", controllers.GetTaskById)
	V1.POST("/task", controllers.NewTask)
	V1.PATCH("/task/:id", controllers.UpdateTask)
	V1.DELETE("/task/:id", controllers.DeleteTask)

	router.Run(":5000")
}
