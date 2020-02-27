package app

import "github.com/mombe090/todo/backend/controllers/todo_controller"

func urlMapper() {
	//Todos API/endPoint
	router.GET("/todos", todo_controller.Index)
	router.POST("/todos", todo_controller.Store)
	router.GET("/todos/:id", todo_controller.Show)
	router.PUT("/todos/:id", todo_controller.Update)
	router.DELETE("/todos/:id", todo_controller.Destroy)
}
