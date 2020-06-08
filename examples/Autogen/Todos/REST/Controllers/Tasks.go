package pgucontrollers

import (
	Request "github.com/eshu0/RESTServer/pkg/request"
	hndlr "github.com/eshu0/Pangu/examples/Autogen/Todos/Handlers"
	models "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
	per "github.com/eshu0/persist/pkg/interfaces"
	RSServer "github.com/eshu0/RESTServer/pkg/server"
	"github.com/eshu0/persist/pkg/sqllite"

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
		controller.Server.Log.LogDebug("HandleRequest", "Calling to insert a new Task")
		result := controller.TasksHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.Log.LogDebug("HandleRequest", "Calling to update Task")
		result := controller.TasksHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.Log.LogDebug("HandleRequest", "Calling to update the (DELETE) Task")
		result := controller.TasksHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.Log.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.TasksHandler.FindById(int64(*Id))
			fmt.Println("----")
			fmt.Println("Result")
			fmt.Println("----")
			fmt.Println(result)
			return result
		} else {
			controller.Server.Log.LogError("HandleRequest", "Id was nil")
		}
	}
	
	controller.Server.Log.LogError("HandleRequest", "Failed returning empty SQLLiteResult")
	return SQLL.NewEmptyFailedSQLLiteQueryResult()
}

func (controller *TasksController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.Log.LogDebug("HandleRequest", "Calling to read all Task")
	result := controller.TasksHandler.ReadAll()
	return result
}


