package main

import (
	"flag"

	ds "github.com/eshu0/Pangu/examples/Autogen/Todos/DataStore"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"	
	
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
	server.Register("ProjectsHandler",&ProjectsHandler)

	

	JobsHandler := fds.GetJobsHandler()
	server.Register("JobsHandler",&JobsHandler)

	

	TasksHandler := fds.GetTasksHandler()
	server.Register("TasksHandler",&TasksHandler)

	

	JobHasTasksHandler := fds.GetJobHasTasksHandler()
	server.Register("JobHasTasksHandler",&JobHasTasksHandler)

	

	ProjectHasJobsHandler := fds.GetProjectHasJobsHandler()
	server.Register("ProjectHasJobsHandler",&ProjectHasJobsHandler)

	

	

	newProject  := data.Project{}
	server.AddJSONFunctionHandler("/Project/Create/","HandleCreateRequest","POST","ProjectsHandler",newProject)
	server.AddJSONFunctionHandler("/Project/Create1/","HandleCreateRequest1","POST","ProjectsHandler",newProject)
	server.AddJSONFunctionHandler("/Project/Create2/","HandleCreateRequest2","POST","ProjectsHandler",newProject)

	

	newJob  := data.Job{}
	server.AddJSONFunctionHandler("/Job/Create/","HandleCreateRequest","POST","JobsHandler",newJob)
	server.AddJSONFunctionHandler("/Job/Create1/","HandleCreateRequest1","POST","JobsHandler",newJob)
	server.AddJSONFunctionHandler("/Job/Create2/","HandleCreateRequest2","POST","JobsHandler",newJob)

	

	newTask  := data.Task{}
	server.AddJSONFunctionHandler("/Task/Create/","HandleCreateRequest","POST","TasksHandler",newTask)
	server.AddJSONFunctionHandler("/Task/Create1/","HandleCreateRequest1","POST","TasksHandler",newTask)
	server.AddJSONFunctionHandler("/Task/Create2/","HandleCreateRequest2","POST","TasksHandler",newTask)

	

	newJobHasTask  := data.JobHasTask{}
	server.AddJSONFunctionHandler("/JobHasTask/Create/","HandleCreateRequest","POST","JobHasTasksHandler",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/Create1/","HandleCreateRequest1","POST","JobHasTasksHandler",newJobHasTask)
	server.AddJSONFunctionHandler("/JobHasTask/Create2/","HandleCreateRequest2","POST","JobHasTasksHandler",newJobHasTask)

	

	newProjectHasJob  := data.ProjectHasJob{}
	server.AddJSONFunctionHandler("/ProjectHasJob/Create/","HandleCreateRequest","POST","ProjectHasJobsHandler",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/Create1/","HandleCreateRequest1","POST","ProjectHasJobsHandler",newProjectHasJob)
	server.AddJSONFunctionHandler("/ProjectHasJob/Create2/","HandleCreateRequest2","POST","ProjectHasJobsHandler",newProjectHasJob)

	


	// start Listen Server, this build the mapping and creates Handler/
	// also fires the "http listen and server method"
	server.ListenAndServe()

}


