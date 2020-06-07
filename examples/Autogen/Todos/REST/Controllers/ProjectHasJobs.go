package pgucontrollers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
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

func (controller *ProjectHasJobsController) HandleCreateRequest(DataIn ProjectHasJob) ProjectHasJob {
	result := controller.ProjectHasJobsHandler.Create(DataIn)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


