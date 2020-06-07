package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

// Controller

type ProjectsController struct {
	ProjectsHandler *hndlr.ProjectsHandler // Storage handler
}

func NewProjectsController(handler *hndlr.ProjectsHandler) *ProjectsController {
	ds := ProjectsController{}
	ds.ProjectsHandler = handler
	return &ds
}


func (controller *ProjectsController) HandleCreateRequest(request Request.ServerRequest) models.Project {
	result := controller.ProjectsHandler.Create(request.Payload)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


