package mongoUtils

import (
	mgo "gopkg.in/mgo.v2"
)

const conStr string = "mongodb://{user}:{pass}@ds258963.mlab.com:258963/tasks"

// connectMongo Connect to MongoDB
func connectMongo() *mgo.Session {

	session, err := mgo.Dial(getConnString())
	if err != nil {
		panic(err)
	}

	//defer session.Close()
	return session
}

// getConnString get the connection string
func getConnString() string {
	return conStr
}

// getDBName Get database name
func getDBName() string {
	return "tasks"
}

// GetSessionCollection Get session and collection reference
func GetSessionCollection(collection string) (*mgo.Session, *mgo.Collection) {
	session := connectMongo()
	_collection := session.DB(getDBName()).C(collection)
	return session, _collection
}
