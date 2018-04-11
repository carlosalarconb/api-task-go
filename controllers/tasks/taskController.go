package controller

import (
	utils "api-tasks/common/mongo"
	model "api-tasks/models/task"

	"gopkg.in/mgo.v2/bson"
)

const collectionName string = "tasks"

// GetAll get all tasks
func GetAll() []model.Task {
	session, collection := utils.GetSessionCollection(collectionName)
	defer session.Close()

	var result []model.Task

	err := collection.Find(nil).All(&result)

	if err != nil {
		panic("Error calling to BD")
	}

	return result
}

//GetOne get one task by id
func GetOne(id string) []model.Task {
	session, collection := utils.GetSessionCollection(collectionName)
	defer session.Close()

	var result []model.Task

	err := collection.Find(bson.M{"id": id}).One(&result)

	if err != nil {
		panic("Error calling to BD")
	}

	return result
}
