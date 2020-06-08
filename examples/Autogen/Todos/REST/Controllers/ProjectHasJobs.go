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
}

func NewProjectHasJobsController(handler *hndlr.ProjectHasJobsHandler) *ProjectHasJobsController {
	ds := ProjectHasJobsController{}
	ds.ProjectHasJobsHandler = handler
	return &ds
}

func (controller *ProjectHasJobsController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)
	
	if request.Request.Method == "POST" {

		data := request.Payload.(models.ProjectHasJob)

		result := controller.ProjectHasJobsHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result
	}

}

func (controller *ProjectHasJobsController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.ProjectHasJob {
	data := request.Payload.(models.ProjectHasJob)
	data.Archived = 1
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


