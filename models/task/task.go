package model

import "gopkg.in/mgo.v2/bson"

// Task struct for tasks
type Task struct {
	_ID         bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Status      bool          `bson:"status" json:"status"`
}
