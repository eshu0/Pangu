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

type TasksController struct {
	TasksHandler *hndlr.TasksHandler // Storage handler
	Server *RSServer.RServer
}

func NewTasksController(handler *hndlr.TasksHandler, Server *RSServer.RServer) *TasksController {
	ds := TasksController{}
	ds.TasksHandler = handler
	ds.Server = Server
	return &ds
}

func (controller *TasksController) HandleRequest(request Request.ServerRequest) per.IQueryResult {  //.Task {
	data := request.Payload.(models.Task)
	
	if request.Request.Method == "POST" {

		result := controller.TasksHandler.Create(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "PUT" { 
	
		result := controller.TasksHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		result := controller.TasksHandler.Update(data)
		fmt.Println("----")
		fmt.Println("Result")
		fmt.Println("----")
		fmt.Println(result)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			result := controller.TasksHandler.FindById(int64(*Id))
			fmt.Println("----")
			fmt.Println("Result")
			fmt.Println("----")
			fmt.Println(result)
			return result
		}
	}

	return "Failed"
}

func (controller *TasksController) HandleRemoveRequest(request Request.ServerRequest) per.IQueryResult {  //.Task {
	data := request.Payload.(models.Task)
	result := controller.TasksHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *TasksController) HandleCreateRequest(request Request.ServerRequest) per.IQueryResult {  //.Task {
	data := request.Payload.(models.Task)

	result := controller.TasksHandler.Create(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *TasksController) HandleUpdateRequest(request Request.ServerRequest) per.IQueryResult {  //.Task {
	data := request.Payload.(models.Task)

	result := controller.TasksHandler.Update(data)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


func (controller *TasksController) HandleFindByIdRequest(request Request.ServerRequest) per.IQueryResult { 
	data := request.Payload.(models.Task)

	result := controller.TasksHandler.FindById(data.Id)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}

func (controller *TasksController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	result := controller.TasksHandler.ReadAll()
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


