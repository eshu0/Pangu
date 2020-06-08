package pgucontrollers

import (
	"fmt"
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"

)

// Controller

type ProjectsController struct {
	ProjectsHandler *hndlr.ProjectsHandler // Storage handler
}

func NewProjectsController(handler *hndlr.ProjectsHandler) *ProjectsController {
	ds := ProjectsController{}
	ds.ProjectsHandler = handler
	return &ds
}

func (controller *ProjectsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)
	
	if request.Request.Method == "POST" {

		data := request.Payload.(models.Project)

		result := controller.ProjectsHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result
	}

}

func (controller *ProjectsController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.Project {
	data := request.Payload.(models.Project)
	data.Archived = 1
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


