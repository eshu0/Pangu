package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
)

// Controller

type TasksController struct {
	TasksHandler *TasksHandler // Storage handler
}

func NewTasksController(handler *TasksHandler) *TasksController {
	ds := TasksController{}
	ds.TasksHandler = handler
	return &ds
}


func (controller *TasksController) HandleCreateRequest(request Request.ServerRequest) Task {
	result := controller.TasksHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


