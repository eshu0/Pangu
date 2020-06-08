package main

import (
	"flag"

	ds "github.com/eshu0/Pangu/examples/Autogen/Todos/DataStore"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"	
	Controllers "github.com/eshu0/Pangu/examples/Autogen/Todos/REST/Controllers"	
	
	"github.com/eshu0/RESTServer/pkg/commands"	
	RSConfig "github.com/eshu0/RESTServer/pkg/config"
	RSServer "github.com/eshu0/RESTServer/pkg/server"
)

func main() {

	dbname := flag.String("db", "./todos.db", "Database defaults to ./todos.db")
	flag.Parse()

	// create a new server
	conf := RSConfig.NewRServerConfig()

	// Create a new REST Server
	server := RSServer.NewRServer(conf)

	//defer the close till the server has closed
	defer server.Log.CloseAllChannels()

	// load this first
	server.ConfigFilePath = "./config.json"

	ok := server.LoadConfig()

	if !ok {
		server.Log.LogErrorf("Main", "Error : %s","Failed to load configuration server not started")
		return
	}
	
	// add the defaults here
	RESTCommands.AddDefaults(server)
	RESTCommands.SetDefaultFunctionalMap(server)

	fds := ds.CreateDataStorage(server.Log, *dbname)

	

	ProjectsHandler := fds.GetProjectsHandler()
	ProjectsController := Controllers.NewProjectsController(ProjectsHandler)
	server.Register("ProjectsController",ProjectsController)

	

	JobsHandler := fds.GetJobsHandler()
	JobsController := Controllers.NewJobsController(JobsHandler)
	server.Register("JobsController",JobsController)

	

	TasksHandler := fds.GetTasksHandler()
	TasksController := Controllers.NewTasksController(TasksHandler)
	server.Register("TasksController",TasksController)

	

	JobHasTasksHandler := fds.GetJobHasTasksHandler()
	JobHasTasksController := Controllers.NewJobHasTasksController(JobHasTasksHandler)
	server.Register("JobHasTasksController",JobHasTasksController)

	

	ProjectHasJobsHandler := fds.GetProjectHasJobsHandler()
	ProjectHasJobsController := Controllers.NewProjectHasJobsController(ProjectHasJobsHandler)
	server.Register("ProjectHasJobsController",ProjectHasJobsController)

	

	

	newProject  := data.Project{}
	server.AddJSONFunctionHandler("/Project/Create/","HandleCreateRequest","POST","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Project/Update/","HandleUpdateRequest","POST","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Project/","HandleReadAllRequest","GET","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Project/Find/","HandleFindByIdRequest","POST","ProjectsController",newProject)

	

	newJob  := data.Job{}
	server.AddJSONFunctionHandler("/Job/Create/","HandleCreateRequest","POST","JobsController",newJob)
	server.AddJSONFunctionHandler("/Job/Update/","HandleUpdateRequest","POST","JobsController",newJob)
	server.AddJSONFunctionHandler("/Job/","HandleReadAllRequest","GET","JobsController",newJob)
	server.AddJSONFunctionHandler("/Job/Find/","HandleFindByIdRequest","POST","JobsController",newJob)

	

	newTask  := data.Task{}
	server.AddJSONFunctionHandler("/Task/Create/","HandleCreateRequest","POST","TasksController",newTask)
	server.AddJSONFunctionHandler("/Task/Update/","HandleUpdateRequest","POST","TasksController",newTask)
	server.AddJSONFunctionHandler("/Task/","HandleReadAllRequest","GET","TasksController",newTask)
	server.AddJSONFunctionHandler("/Task/Find/","HandleFindByIdRequest","POST","TasksController",newTask)

	

	newJobHasTask  := data.JobHasTask{}
	server.AddJSONFunctionHandler("/JobHasTask/Create/","HandleCreateRequest","POST","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/Update/","HandleUpdateRequest","POST","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/","HandleReadAllRequest","GET","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/Find/","HandleFindByIdRequest","POST","JobHasTasksController",newJobHasTask)

	

	newProjectHasJob  := data.ProjectHasJob{}
	server.AddJSONFunctionHandler("/ProjectHasJob/Create/","HandleCreateRequest","POST","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/Update/","HandleUpdateRequest","POST","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/","HandleReadAllRequest","GET","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/Find/","HandleFindByIdRequest","POST","ProjectHasJobsController",newProjectHasJob)

	


	// start Listen Server, this build the mapping and creates Handler/
	// also fires the "http listen and server method"
	server.ListenAndServe()

}


