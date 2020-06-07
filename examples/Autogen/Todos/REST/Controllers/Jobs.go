package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

// Controller

type JobsController struct {
	JobsHandler *hndlr.JobsHandler // Storage handler
}

func NewJobsController(handler *hndlr.JobsHandler) *JobsController {
	ds := JobsController{}
	ds.JobsHandler = handler
	return &ds
}


func (controller *JobsController) HandleCreateRequest(request Request.ServerRequest) models.Job {
	result := controller.JobsHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


