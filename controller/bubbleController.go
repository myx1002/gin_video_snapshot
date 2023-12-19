package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_video_snapshot/model"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	todo := new(model.TODO)
	_ = c.BindJSON(&todo)
	if err := model.CreateTodo(todo); err == nil {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}
}

func GetTodoList(c *gin.Context) {
	var todoList []model.TODO
	if err := model.GetTodoList(&todoList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := model.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	fmt.Println(111111)
	err = c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println(22222)

	err = model.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	todo := new(model.TODO)
	id := c.Param("id")

	if err := model.DeleteTodo(id, todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}
