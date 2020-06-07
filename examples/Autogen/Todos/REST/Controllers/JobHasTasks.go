package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

// Controller

type JobHasTasksController struct {
	JobHasTasksHandler *hndlr.JobHasTasksHandler // Storage handler
}

func NewJobHasTasksController(handler *hndlr.JobHasTasksHandler) *JobHasTasksController {
	ds := JobHasTasksController{}
	ds.JobHasTasksHandler = handler
	return &ds
}


func (controller *JobHasTasksController) HandleCreateRequest(request Request.ServerRequest) models.JobHasTask {
	result := controller.JobHasTasksHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


