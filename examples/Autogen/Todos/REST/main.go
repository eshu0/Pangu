package main

import (
	"flag"

	ds "github.com/eshu0/Pangu/examples/Autogen/Todos/DataStore"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"	
	Controllers "github.com/eshu0/Pangu/examples/Autogen/Todos/REST/Controllers"	
	
	RESTCommands "github.com/eshu0/RESTServer/pkg/commands"
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

	// load this first
	server.ConfigFilePath = "./config.json"

	ok := server.LoadConfig()

	if !ok {
		server.LogErrorf("Main", "Error : %s","Failed to load configuration server not started")
		return
	}
	
	// add the defaults here
	RESTCommands.AddDefaults(server)
	RESTCommands.SetDefaultFunctionalMap(server)

	fds := ds.CreateDataStorage(*dbname)

	

	ProjectsHandler := fds.GetProjectsHandler()
	ProjectsController := Controllers.NewProjectsController(ProjectsHandler, server)
	server.Register("ProjectsController",ProjectsController)

	

	JobsHandler := fds.GetJobsHandler()
	JobsController := Controllers.NewJobsController(JobsHandler, server)
	server.Register("JobsController",JobsController)

	

	TasksHandler := fds.GetTasksHandler()
	TasksController := Controllers.NewTasksController(TasksHandler, server)
	server.Register("TasksController",TasksController)

	

	JobHasTasksHandler := fds.GetJobHasTasksHandler()
	JobHasTasksController := Controllers.NewJobHasTasksController(JobHasTasksHandler, server)
	server.Register("JobHasTasksController",JobHasTasksController)

	

	ProjectHasJobsHandler := fds.GetProjectHasJobsHandler()
	ProjectHasJobsController := Controllers.NewProjectHasJobsController(ProjectHasJobsHandler, server)
	server.Register("ProjectHasJobsController",ProjectHasJobsController)

	

	

	newProject  := data.Project{}
	
	server.AddJSONFunctionHandler("/Project/{Id}/","HandleRequest","GET","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Project/","HandleRequest","","ProjectsController",newProject)
	server.AddJSONFunctionHandler("/Projects/","HandleReadAllRequest","GET","ProjectsController",newProject)

	

	newJob  := data.Job{}
	
	server.AddJSONFunctionHandler("/Job/{Id}/","HandleRequest","GET","JobsController",newJob)
	server.AddJSONFunctionHandler("/Job/","HandleRequest","","JobsController",newJob)
	server.AddJSONFunctionHandler("/Jobs/","HandleReadAllRequest","GET","JobsController",newJob)

	

	newTask  := data.Task{}
	
	server.AddJSONFunctionHandler("/Task/{Id}/","HandleRequest","GET","TasksController",newTask)
	server.AddJSONFunctionHandler("/Task/","HandleRequest","","TasksController",newTask)
	server.AddJSONFunctionHandler("/Tasks/","HandleReadAllRequest","GET","TasksController",newTask)

	

	newJobHasTask  := data.JobHasTask{}
	
	server.AddJSONFunctionHandler("/JobHasTask/{Id}/","HandleRequest","GET","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/","HandleRequest","","JobHasTasksController",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTasks/","HandleReadAllRequest","GET","JobHasTasksController",newJobHasTask)

	

	newProjectHasJob  := data.ProjectHasJob{}
	
	server.AddJSONFunctionHandler("/ProjectHasJob/{Id}/","HandleRequest","GET","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/","HandleRequest","","ProjectHasJobsController",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJobs/","HandleReadAllRequest","GET","ProjectHasJobsController",newProjectHasJob)

	


	// start Listen Server, this build the mapping and creates Handler/
	// also fires the "http listen and server method"
	server.ListenAndServe()

}


