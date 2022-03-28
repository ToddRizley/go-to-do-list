package controllers

import (
	"github.com/gin-gonic/gin"
	"go-to-do-list/models"
	"net/http"
)

func GetTodoItems(c *gin.Context) {
	var todo []models.TodoItem
	err := models.GetAllTodoItems(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateATodoItem(c *gin.Context) {
	var todo models.TodoItem
	c.BindJSON(&todo)
	err := models.CreateATodoItem(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodoItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.TodoItem
	err := models.GetATodoItem(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodoItem(c *gin.Context) {
	var todo models.TodoItem
	id := c.Params.ByName("id")
	err := models.GetATodoItem(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = models.UpdateATodoItem(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetAPIHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "API Healthy")
}

func DeleteATodoItem(c *gin.Context) {
	var todo models.TodoItem
	id := c.Params.ByName("id")
	err := models.DeleteATodoItem(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
