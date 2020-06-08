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

type ProjectsController struct {
	ProjectsHandler *hndlr.ProjectsHandler // Storage handler
	Server *RSServer.RServer
}

func NewProjectsController(handler *hndlr.ProjectsHandler, Server *RSServer.RServer) *ProjectsController {
	ds := ProjectsController{}
	ds.ProjectsHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *ProjectsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)
	
	if request.Request.Method == "POST" {

		result := controller.ProjectsHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "PUT" { 
	
		result := controller.ProjectsHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		result := controller.ProjectsHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")

		result := controller.ProjectsHandler.FindById(int64(Id))
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result
	}

}

func (controller *ProjectsController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)
	result := controller.ProjectsHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *ProjectsController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)

	result := controller.ProjectsHandler.Create(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *ProjectsController) HandleUpdateRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)

	result := controller.ProjectsHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *ProjectsController) HandleFindByIdRequest(request Request.ServerRequest) per.IQueryResult { 
	data := request.Payload.(models.Project)

	result := controller.ProjectsHandler.FindById(data.Id)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *ProjectsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	result := controller.ProjectsHandler.ReadAll()
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


