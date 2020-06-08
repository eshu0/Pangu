package pgucontrollers

import (
	"fmt"
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"

)

// Controller

type JobsController struct {
	JobsHandler *hndlr.JobsHandler // Storage handler
}

func NewJobsController(handler *hndlr.JobsHandler) *JobsController {
	ds := JobsController{}
	ds.JobsHandler = handler
	return &ds
}

func (controller *JobsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)
	
	if request.Request.Method == "POST" {

		data := request.Payload.(models.Job)

		result := controller.JobsHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result
	}

}

func (controller *JobsController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)
	data.Archived = 1
	result := controller.JobsHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *JobsController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)

	result := controller.JobsHandler.Create(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *JobsController) HandleUpdateRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)

	result := controller.JobsHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *JobsController) HandleFindByIdRequest(request Request.ServerRequest) per.IQueryResult { 
	data := request.Payload.(models.Job)

	result := controller.JobsHandler.FindById(data.Id)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *JobsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	result := controller.JobsHandler.ReadAll()
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


