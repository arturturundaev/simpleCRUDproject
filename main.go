package main

import (
    userController "./controller/userController"
    "./intfc/controllerIntfc"
    "fmt"
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
)

var tmp, _ = template.ParseFiles("view/user/showAll.html", "view/user/current.html")

/**
  Main method. ALl request will be here
*/
func main() {
    request := mux.NewRouter()
    request.HandleFunc("/{entity}", ParceHandler)
    request.HandleFunc("/{entity}/{action:add|save}", ParceHandler)
    request.HandleFunc("/{entity}/{action:view|edit|delete}/{id}", ParceHandler)

    http.Handle("/", request)

    fmt.Println("Start server on port 3030")
    log.Fatal(http.ListenAndServe(":3030", nil))
}

/** Get needly controller  */
func getNeedlyController(entity string) controllerIntfc.ControllerIntfc {
    var currentController controllerIntfc.ControllerIntfc

    switch entity {
    case "user":
        currentController = userController.UserController{}
    }

    return currentController
}

/**
  Parsing url params and call necessary function
*/
func ParceHandler(responseWriter http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    entity := vars["entity"] /* Type of entity. For example user/post */
    action := vars["action"] /* Type of action. What we want to do. For example Edit, Add, Delete, Show */
    id := vars["id"]         /* Id of necessary object. We found object and do our "action" */
    contr := getNeedlyController(entity)

    if entity == "favicon.ico" {
        return
    }

    switch action {
    case "delete":
        delete(contr, id)
    case "edit":
        showPageToEditEntity(contr, entity, id, responseWriter)
        return
    case "save":
        saveEntityAndRedirect(contr, request)
    case "add":
        showFormToAddNewEntity(entity, responseWriter)
        return
    case "":
        showAll(contr, entity, responseWriter)
        return
    }

    http.Redirect(responseWriter, request, "/"+entity, http.StatusSeeOther)
}

/* Get entity by id and show it */
func showPageToEditEntity(contr controllerIntfc.ControllerIntfc, entity string, id string, responseWriter http.ResponseWriter) {
    model := contr.Get(id)
    tmp.ExecuteTemplate(responseWriter, entity+"Current", model)
}

/* Save entity to memory */
func saveEntityAndRedirect(contr controllerIntfc.ControllerIntfc, request *http.Request) {
    request.ParseForm()
    data := getFormData(request)
    contr.Add(data)
}

/* Show page to add new entity */
func showFormToAddNewEntity(entity string, responseWriter http.ResponseWriter) {
    tmp.ExecuteTemplate(responseWriter, entity+"Current", nil)
}

/* Show all entities */
func showAll(contr controllerIntfc.ControllerIntfc, entity string, responseWriter http.ResponseWriter) {
    EntityArr := contr.ShowAll()
    tmp.ExecuteTemplate(responseWriter, entity+"showAll", EntityArr)
}

/* Delete entity */
func delete(contr controllerIntfc.ControllerIntfc, id string) {
    contr.Delete(id)
}

/* Parsing form data to array */
func getFormData(request *http.Request) map[string]string {
    var data = make(map[string]string, 0)
    request.ParseForm()
    for key, value := range request.Form {
        data[key] = value[0]
    }

    return data
}
