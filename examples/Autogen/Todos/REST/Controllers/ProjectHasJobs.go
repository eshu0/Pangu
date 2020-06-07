package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
)

// Controller

type ProjectHasJobsController struct {
	ProjectHasJobsHandler *ProjectHasJobsHandler // Storage handler
}

func NewProjectHasJobsController(handler *ProjectHasJobsHandler) *ProjectHasJobsController {
	ds := ProjectHasJobsController{}
	ds.ProjectHasJobsHandler = handler
	return &ds
}


func (controller *ProjectHasJobsController) HandleCreateRequest(request Request.ServerRequest) ProjectHasJob {
	result := controller.ProjectHasJobsHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


