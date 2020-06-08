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
		controller.Server.Log.LogDebug("HandleRequest", "Calling to insert a new Job")
		result := controller.JobsHandler.Create(data)
		return result

	} else if request.Request.Method == "PUT" { 
	
		controller.Server.Log.LogDebug("HandleRequest", "Calling to update Job")
		result := controller.JobsHandler.Update(data)
		return result

	} else if request.Request.Method == "DELETE" { 
	
		controller.Server.Log.LogDebug("HandleRequest", "Calling to update the (DELETE) Job")
		result := controller.JobsHandler.Update(data)
		return result

	} else {
		
		Id := controller.Server.RequestHelper.GetRequestId(request.Request,"Id")
		if Id != nil {
			controller.Server.Log.LogDebugf("HandleRequest", "Id was not nil and have the following to lookup %d", *Id)
			result := controller.JobsHandler.FindById(int64(*Id))
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

func (controller *JobsController) HandleReadAllRequest(request Request.ServerRequest) per.IQueryResult { 
	controller.Server.Log.LogDebug("HandleRequest", "Calling to read all Job")
	result := controller.JobsHandler.ReadAll()
	return result
}


