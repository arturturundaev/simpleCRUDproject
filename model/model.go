package model

import (
	uuid "github.com/google/uuid"
)

type Entity struct {
	Id       string
	Login    string
	Password string
}

var EntityArr = make(map[string]interface{})

/* Get entity by id from storage */
func Get(id string) interface{} {
	return EntityArr[id]
}

/* Return all entities from storage */
func GetAll() map[string]interface{} {
	return EntityArr
}

/* Update entity by id */
func Update(data map[string]string) {
	user := getEntity(data)
	EntityArr[user.Id] = user
}

/* Insert new entity into storage */
func Create(data map[string]string) {
	user := getEntity(data)
	EntityArr[user.Id] = user
}

/* Delete entity from storage */
func Delete(id string) {
	delete(EntityArr, id)
}

func getEntity(data map[string]string) Entity {
	if data["Id"] == "" {
		data["Id"] = uuid.New().String()
	}

	return Entity{Id: data["Id"], Login: data["Login"], Password: data["Password"]}
}
