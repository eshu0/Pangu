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


