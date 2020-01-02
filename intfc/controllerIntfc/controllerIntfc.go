package controllerIntfc

type ControllerIntfc interface {
	ShowAll() map[string]interface{}
  	Delete(id string)
  	Add(map[string]string)
  	Get(id string)interface{}
}
