package todo

import (
	"fmt"
	"github.com/mombe090/utils/errors_utils"
	"github.com/mombe090/utils/mongo_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"strings"
)

func (todo *Todo) Save() *errors_utils.RestError {
	res, err := mongo_utils.InsertOne(os.Getenv("MONGODB"), os.Getenv("MONGOCOLLECTION"), todo)
	if err != nil {
		return errors_utils.GetInternatServerError(fmt.Sprintf("Insertion Errors %s", err.Error()))
	}
	todo.Id = res

	return nil
}

func (todo *Todo) Update(id string) *errors_utils.RestError {
	todo.Id = id
	ojb, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors_utils.GetBadRequest(fmt.Sprintf("Can not convert %s into OBJECT ID %s", id, err))
	}

	err = mongo_utils.UpdateOne(
		os.Getenv("MONGODB"),
		os.Getenv("MONGOCOLLECTION"),
		bson.M{"_id": ojb},
		bson.D{
			{"$set", bson.D{
				{
					"title", todo.Title,
				},
				{
					"done", todo.Done,
				},
			},
			},
		})

	if err != nil {
		return errors_utils.GetInternatServerError(fmt.Sprintf("Update Errors %s", err.Error()))
	}

	return nil
}

func (todo *Todo) GetAll() (*[]Todo, *errors_utils.RestError) {
	curr, ctx, err := mongo_utils.FindMany(os.Getenv("MONGODB"), os.Getenv("MONGOCOLLECTION"), bson.D{})

	if err != nil {
		return nil, errors_utils.GetInternatServerError(fmt.Sprintf("Get Message null %s", err.Error()))
	}

	var todos []Todo
	for curr.Next(*ctx) {
		var resTmp map[string]interface{}
		err := curr.Decode(&resTmp)

		if err != nil {
			log.Print(err)
		}
		done, ok := resTmp["done"].(bool)

		if !ok {
			log.Printf("got data of type %T but wanted int", resTmp["done"])
			continue
		}

		id, ok := resTmp["_id"].(primitive.ObjectID)
		if !ok {
			log.Printf("got data of type %T but wanted int", resTmp["_id"])
			continue
		}

		todoTmp := Todo{
			Id:    id.Hex(),
			Title: fmt.Sprintf("%s", resTmp["title"]),
			Done:  done,
		}
		todos = append(todos, todoTmp)
	}

	return &todos, nil
}

func GetOne(m bson.M) (*Todo, *errors_utils.RestError) {
	singleResul := mongo_utils.FindOne(os.Getenv("MONGODB"), os.Getenv("MONGOCOLLECTION"), m)

	var todo Todo
	err := singleResul.Decode(&todo)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			log.Println("No document founded {}")
			return &Todo{}, nil
		}
		return nil, errors_utils.GetInternatServerError(err.Error())
	}

	return &todo, nil
}

func (todo *Todo) DeleteOne() *errors_utils.RestError {
	obj, err := primitive.ObjectIDFromHex(todo.Id)

	if err != nil {
		return errors_utils.GetBadRequest(fmt.Sprintf("Can not convert %s into OBJECT ID %s", todo.Id, err))
	}

	err = mongo_utils.DeleteOne(os.Getenv("MONGODB"), os.Getenv("MONGOCOLLECTION"), bson.M{"_id": obj})

	if err != nil {
		return errors_utils.GetInternatServerError(fmt.Sprintf("Impossible to delete %s", err))
	}

	return nil
}
