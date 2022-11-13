package routes

import (
	"src/crud"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default() // *gin.Engineを返す関数
	router.POST("/task", crud.CreateNewTask)
	router.GET("/task", crud.GetAllTask)
	router.GET("/task/:task-id", crud.GetTask)
	router.PATCH("/task/:task-id", crud.UpdateTask)
	router.DELETE("/task/:task-id", crud.DeleteTask)
	return router
}