package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
)

// Controller

type JobHasTasksController struct {
	JobHasTasksHandler *JobHasTasksHandler // Storage handler
}

func NewJobHasTasksController(handler *JobHasTasksHandler) *JobHasTasksController {
	ds := JobHasTasksController{}
	ds.JobHasTasksHandler = handler
	return &ds
}


func (controller *JobHasTasksController) HandleCreateRequest(request Request.ServerRequest) JobHasTask {
	result := controller.JobHasTasksHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


