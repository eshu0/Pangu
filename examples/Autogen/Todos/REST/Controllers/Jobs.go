package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
)

// Controller

type JobsController struct {
	JobsHandler *JobsHandler // Storage handler
}

func NewJobsController(handler *JobsHandler) *JobsController {
	ds := JobsController{}
	ds.JobsHandler = handler
	return &ds
}


func (controller *JobsController) HandleCreateRequest(request Request.ServerRequest) Job {
	result := controller.JobsHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


