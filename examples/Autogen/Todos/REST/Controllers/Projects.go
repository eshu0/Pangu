package pgucontrollers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
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

func (controller *ProjectsController) HandleCreateRequest2(DataIn *Project) Project {
	result := controller.ProjectsHandler.Create(DataIn)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *ProjectsController) HandleCreateRequest1(DataIn interface{}) Project {

	item := DataIn.(Project)

	result := controller.ProjectsHandler.Create(item)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *ProjectsController) HandleCreateRequest(DataIn Project) Project {
	result := controller.ProjectsHandler.Create(DataIn)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


