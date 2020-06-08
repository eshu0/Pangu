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
		controller.Server.Log.LogDebug("HandleRequest", "Calling to insert a new JobHasTask")
		result := controller.JobHasTasksHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.Log.LogDebug("HandleRequest", "Calling to update JobHasTask")
		result := controller.JobHasTasksHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.Log.LogDebug("HandleRequest", "Calling to update the (DELETE) JobHasTask")
		result := controller.JobHasTasksHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.Log.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.JobHasTasksHandler.FindById(int64(*Id))
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

func (controller *JobHasTasksController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.Log.LogDebug("HandleRequest", "Calling to read all JobHasTask")
	result := controller.JobHasTasksHandler.ReadAll()
	return result
}


