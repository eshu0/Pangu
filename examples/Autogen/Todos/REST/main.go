package main

import (
	"flag"
	"fmt"
	"encoding/json"
	"reflect"

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
	res := TestJSON(newProject)
	fmt.Printf("Resut: %v \n",res)

	tres := res.(*data.Project)
	fmt.Printf("Resutt: %v \n",tres)
	return

	server.AddJSONFunctionHandler("/Project/Create/","HandleCreateRequest","POST","ProjectsController",newProject)

	

	newJob  := data.Job{}
	server.AddJSONFunctionHandler("/Job/Create/","HandleCreateRequest","POST","JobsController",newJob)

	

	newTask  := data.Task{}
	server.AddJSONFunctionHandler("/Task/Create/","HandleCreateRequest","POST","TasksController",newTask)

	

	newJobHasTask  := data.JobHasTask{}
	server.AddJSONFunctionHandler("/JobHasTask/Create/","HandleCreateRequest","POST","JobHasTasksController",newJobHasTask)

	

	newProjectHasJob  := data.ProjectHasJob{}
	server.AddJSONFunctionHandler("/ProjectHasJob/Create/","HandleCreateRequest","POST","ProjectHasJobsController",newProjectHasJob)

	


	// start Listen Server, this build the mapping and creates Handler/
	// also fires the "http listen and server method"
	server.ListenAndServe()

}

var jsonData = []byte(`{"id" : -1,"displayname" : "Hello","description" : "Something","archived" : 0,"completed" : 0}`)

func TestJSON(Data interface{}) interface{}{

	//rv := reflect.ValueOf(&Data)
	rv := reflect.ValueOf(&Data).Elem()
	fmt.Printf("TestJSON - Got the following rv %v\n",rv)

	// Get first arg of the function
	firstArg := reflect.TypeOf(Data)//.In(0)

	fmt.Printf("TestJSON - Got the following firstArg %v\n",firstArg)

	// Get the PtrTo to the first function parameter
	structPtr := reflect.New(firstArg)//rv.Elem().Type())
	fmt.Printf("TestJSON - Got the following structPtr %v\n",structPtr)


	// Convert to Interface
	// Note that I can't assert this to .(myStruct) type
	instance := structPtr.Interface()
	fmt.Printf("TestJSON - Got the following instance %v\n",instance)

	//b := []byte(str)
	//fmt.Printf("TestJSON - str %s\n",str)

	//err := json.NewDecoder(string(body)).Decode(&Data)
	err := json.Unmarshal(jsonData, &instance)
	if err != nil {
		fmt.Printf("TestJSON - Unmarshal -> Got the following error while unmarchsalling JSON %s\n",err.Error())
		return nil
	}else{
		fmt.Printf("TestJSON - Unmarshal -> Got the following %v\n",Data)
	}

	vfn := reflect.ValueOf(instance)
	fmt.Printf("TestJSON - Got the following instance %v\n",vfn)
	/*
		err := json.Unmarshal(jsonData, &Data)
	if err != nil {
		fmt.Printf("TestJSON - Unmarshal -> Got the following error while unmarchsalling JSON %s\n",err.Error())
		return nil
	}else{
		fmt.Printf("TestJSON - Unmarshal -> Got the following %v\n",Data)
	}
	s := reflect.ValueOf(&Data).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		for j, f := range Data {
			if typeOfT.Field(i).Tag.Get("json") == j {
				fl := s.FieldByName(typeOfT.Field(i).Name)
				switch fl.Kind() {
				case reflect.Bool:
					fl.SetBool(f.(bool))
				case reflect.Int, reflect.Int64:
					c, _ := f.(float64)
					fl.SetInt(int64(c))
				case reflect.String:
					fl.SetString(f.(string))
				}
			}
		}
	}
	*/

	return vfn.Interface()
}


