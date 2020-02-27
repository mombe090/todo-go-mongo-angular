package todo_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mombe090/todo/backend/domains/todo"
	"github.com/mombe090/todo/backend/services/todo_service"
	"github.com/mombe090/utils/errors_utils"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Index(c *gin.Context) {
	todos, restErr := todo_service.GetTodos(todo.Todo{})

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if len(*todos) == 0 {
		c.JSON(http.StatusOK, []todo.Todo{})
		return
	}
	c.JSON(http.StatusOK, todos)
	return
}

func Show(c *gin.Context) {
	todo, restErr := todo_service.GetTodo(c.Param("id"))

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if todo.Id == "" {
		c.JSON(http.StatusOK, bson.M{})
		return
	}

	c.JSON(http.StatusOK, todo)
	return
}

func Store(c *gin.Context) {
	var todo todo.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, errors_utils.GetBadRequest(fmt.Sprintf("Binding %s", err.Error())))
	}

	t, restErr := todo_service.SaveTodo(todo)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, t)
	return
}

func Update(c *gin.Context) {
	var todo todo.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, errors_utils.GetBadRequest(fmt.Sprintf("Binding %s", err.Error())))
		return
	}

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, errors_utils.GetBadRequest("Title must not be empty or null"))
		return
	}

	t, restErr := todo_service.UpdateTodo(todo, c.Param("id"))

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, t)
	return
}

func Destroy(c *gin.Context) {
	restErr := todo_service.DeleteTodo(c.Param("id"))

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, bson.M{"value": "Successfully deleted"})
	return
}
