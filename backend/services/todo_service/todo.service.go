package todo_service

import (
	"fmt"
	"github.com/mombe090/todo/backend/domains/todo"
	"github.com/mombe090/utils/errors_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(t todo.Todo) (*[]todo.Todo, *errors_utils.RestError)  {
	todos, restErr := t.GetAll()

	if restErr != nil{
		return  nil, restErr
	}

	return todos, nil
}

func GetTodo(id string) (*todo.Todo, *errors_utils.RestError)  {
	ojb, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors_utils.GetBadRequest(fmt.Sprintf("Converting id to ObjectID not working %s", err))
	}

	todo, restErr := todo.GetOne(bson.M{"_id": ojb})

	if restErr != nil{
		return  nil, restErr
	}


	return todo, nil
}

func SaveTodo(todo todo.Todo) (*todo.Todo, *errors_utils.RestError)  {
	restErr := todo.Save()

	if restErr != nil{
		return  nil, restErr
	}

	return &todo, nil
}

func UpdateTodo(todo todo.Todo, id string) (*todo.Todo, *errors_utils.RestError)  {
	restErr := todo.Update(id)

	if restErr != nil{
		return  nil, restErr
	}

	return &todo, nil
}

func DeleteTodo(id string) *errors_utils.RestError  {
	var todo = todo.Todo{Id: id}
	return todo.DeleteOne()
}
