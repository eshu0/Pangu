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


