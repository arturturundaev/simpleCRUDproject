package userController

import (
    user "../../model"
)

type UserController struct {
}

/* Get all entities from storage */
func (cntr UserController) ShowAll() map[string]interface{} {
    return user.GetAll()
}

/* Delete entity by id */
func (cntr UserController) Delete(id string) {
    user.Delete(id)
}

/* Insert new entity into storage or update */
func (cntr UserController) Add(formData map[string]string) {
    if formData["Id"] != "" {
        user.Update(formData)
    } else {
        user.Create(formData)
    }
}

/* Get entity from storage by id */
func (cntr UserController) Get(id string) interface{} {
    return user.Get(id)
}
