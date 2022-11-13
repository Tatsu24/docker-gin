package crud

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateNewTask(c *gin.Context) {
	res := gin.H{"post": "task"}
  	c.JSON(200, res)
}
func GetAllTask(c *gin.Context) {
	res := gin.H{"get": "all-task"}
  	c.JSON(200, res)
}
func GetTask(c *gin.Context) {
	taskId, err := strconv.Atoi((c.Param("task-id")))
	if err != nil {
		c.String(http.StatusBadRequest, "task is not found")
		return
	}
	res := gin.H{"get": "task","taskId": taskId}
  	c.JSON(200, res)
}
func UpdateTask(c *gin.Context) {
	taskId, err := strconv.Atoi((c.Param("task-id")))
	if err != nil {
		c.String(http.StatusBadRequest, "task is not found")
		return
	}
	res := gin.H{"update": "task","taskId": taskId}
  	c.JSON(200, res)
}
func DeleteTask(c *gin.Context) {
	taskId, err := strconv.Atoi((c.Param("task-id")))
	if err != nil {
		c.String(http.StatusBadRequest, "task is not found")
		return
	}
	res := gin.H{"delete": "task","taskId": taskId}
  	c.JSON(200, res)
}