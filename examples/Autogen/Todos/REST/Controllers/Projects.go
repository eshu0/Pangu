package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
)

// Controller

type ProjectsController struct {
	ProjectsHandler *ProjectsHandler // Storage handler
}

func NewProjectsController(handler *ProjectsHandler) *ProjectsController {
	ds := ProjectsController{}
	ds.ProjectsHandler = handler
	return &ds
}


func (controller *ProjectsController) HandleCreateRequest(request Request.ServerRequest) Project {
	result := controller.ProjectsHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


