package pgucontrollers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
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


