package pgucontrollers

import (
	"fmt"
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"

)

// Controller

type JobHasTasksController struct {
	JobHasTasksHandler *hndlr.JobHasTasksHandler // Storage handler
}

func NewJobHasTasksController(handler *hndlr.JobHasTasksHandler) *JobHasTasksController {
	ds := JobHasTasksController{}
	ds.JobHasTasksHandler = handler
	return &ds
}


func (controller *JobHasTasksController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.JobHasTask {
	data := request.Payload.(*models.JobHasTask)

	result := controller.JobHasTasksHandler.Create(*data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


