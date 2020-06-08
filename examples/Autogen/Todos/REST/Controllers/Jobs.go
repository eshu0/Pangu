package pgucontrollers

import (
	"fmt"
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"
	RSServer "github.com/eshu0/RESTServer/pkg/server"
)

// Controller

type JobsController struct {
	JobsHandler *hndlr.JobsHandler // Storage handler
	Server *RSServer.RServer
}

func NewJobsController(handler *hndlr.JobsHandler, Server *RSServer.RServer) *JobsController {
	ds := JobsController{}
	ds.JobsHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *JobsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)
	
	if request.Request.Method == "POST" {

		result := controller.JobsHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "PUT" { 
	
		result := controller.JobsHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		result := controller.JobsHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			result := controller.JobsHandler.FindById(int64(*Id))
			fmt.Println("----")
			fmt.Println("Result")
			fmt.Println("----")
			fmt.Println(result)
			return result
		}
	}

	return "Failed"
}

func (controller *JobsController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.Job {
	data := request.Payload.(models.Job)
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


