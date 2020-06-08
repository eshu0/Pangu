package pgucontrollers

import (
	"fmt"
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"

)

// Controller

type ProjectHasJobsController struct {
	ProjectHasJobsHandler *hndlr.ProjectHasJobsHandler // Storage handler
	Server *RSServer.RServer
}

func NewProjectHasJobsController(handler *hndlr.ProjectHasJobsHandler, Server *RSServer.RServer) *ProjectHasJobsController {
	ds := ProjectHasJobsController{}
	ds.ProjectHasJobsHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *ProjectHasJobsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)
	
	if request.Request.Method == "POST" {

		result := controller.ProjectHasJobsHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "PUT" { 
	
		result := controller.ProjectHasJobsHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		result := controller.ProjectHasJobsHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")

		result := controller.ProjectHasJobsHandler.FindById(int64(Id))
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result
	}

}

func (controller *ProjectHasJobsController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)
	result := controller.ProjectHasJobsHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *ProjectHasJobsController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)

	result := controller.ProjectHasJobsHandler.Create(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *ProjectHasJobsController) HandleUpdateRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)

	result := controller.ProjectHasJobsHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *ProjectHasJobsController) HandleFindByIdRequest(request Request.ServerRequest) per.IQueryResult { 
	data := request.Payload.(models.ProjectHasJob)

	result := controller.ProjectHasJobsHandler.FindById(data.Id)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *ProjectHasJobsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	result := controller.ProjectHasJobsHandler.ReadAll()
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


