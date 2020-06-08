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
	Server *RSServer.RServer
}

func NewJobHasTasksController(handler *hndlr.JobHasTasksHandler, Server *RSServer.RServer) *JobHasTasksController {
	ds := JobHasTasksController{}
	ds.JobHasTasksHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *JobHasTasksController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.JobHasTask {
	data := request.Payload.(models.JobHasTask)
	
	if request.Request.Method == "POST" {

		result := controller.JobHasTasksHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "PUT" { 
	
		result := controller.JobHasTasksHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		result := controller.JobHasTasksHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")

		result := controller.JobHasTasksHandler.FindById(int64(Id))
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result
	}

}

func (controller *JobHasTasksController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.JobHasTask {
	data := request.Payload.(models.JobHasTask)
	result := controller.JobHasTasksHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *JobHasTasksController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.JobHasTask {
	data := request.Payload.(models.JobHasTask)

	result := controller.JobHasTasksHandler.Create(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *JobHasTasksController) HandleUpdateRequest(request Request.ServerRequest) per.IQueryResult {  //.JobHasTask {
	data := request.Payload.(models.JobHasTask)

	result := controller.JobHasTasksHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *JobHasTasksController) HandleFindByIdRequest(request Request.ServerRequest) per.IQueryResult { 
	data := request.Payload.(models.JobHasTask)

	result := controller.JobHasTasksHandler.FindById(data.Id)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *JobHasTasksController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	result := controller.JobHasTasksHandler.ReadAll()
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


