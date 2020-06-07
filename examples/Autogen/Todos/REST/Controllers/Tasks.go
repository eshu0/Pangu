package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

// Controller

type TasksController struct {
	TasksHandler *hndlr.TasksHandler // Storage handler
}

func NewTasksController(handler *hndlr.TasksHandler) *TasksController {
	ds := TasksController{}
	ds.TasksHandler = handler
	return &ds
}


func (controller *TasksController) HandleCreateRequest(request Request.ServerRequest) models.Task {
	result := controller.TasksHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


