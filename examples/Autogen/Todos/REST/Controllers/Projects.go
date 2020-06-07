package pgucontrollers

import (
	"fmt"
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"

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


func (controller *ProjectsController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(*models.Project)

	result := controller.ProjectsHandler.Create(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


