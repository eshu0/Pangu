package main

import (
	"flag"
	"strings"
	"encoding/json"
	"fmt"
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
	fmt.Printf("res: %+v\n", res) 
	rest := res.(data.Project)
	fmt.Printf("Id: %+v\n", rest.Id) 

	//return
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

	
func TestJSON(Data interface{}) interface{} {
	

	jsonstr := `{"id" : -1,"displayname" : "Hello","description" : "Something","archived" : 0,"completed" : 0}`
	
	//
	d := map[string]interface{}{}
	json.Unmarshal([]byte(jsonstr), &d)
	
	//
	//obj := data.Project{}
	firstArg := reflect.TypeOf(Data)
	s := reflect.New(firstArg).Elem()

	structPtr := reflect.New(firstArg).Elem()
	//instance := structPtr.Interface()

	//s := reflect.ValueOf(&Data).Elem()
	typeOfT := s.Type()
	//
	for i := 0; i < s.NumField(); i++ {
		for j, f := range d {
			fmt.Printf("j :%+v\n", j) 
			fmt.Printf("%v - %v - %v - %v\n",typeOfT,typeOfT.Field(i),typeOfT.Field(i).Tag,typeOfT.Field(i).Tag.Get("json"))
			fmt.Printf("%v - %v - %v - %v\n",typeOfT,typeOfT.Field(i),typeOfT.Field(i).Tag,typeOfT.Field(i).Tag.Get("json"))

			withoutomit:= typeOfT.Field(i).Tag.Get("json")
			withoutomit = strings.Replace(withoutomit,",omitempty","",-1)
			if withoutomit == j {
				fmt.Printf("Name :%+v\n", typeOfT.Field(i).Name) 

				fl := structPtr.FieldByName(typeOfT.Field(i).Name)
				fmt.Printf("Kind :%+v\n", fl.Kind()) 

				switch fl.Kind() {
					case reflect.Bool:
						fl.SetBool(f.(bool))
					case reflect.Int, reflect.Int64:
						c, _ := f.(float64)
						fmt.Printf("c :%+v\n",c) 

						fl.SetInt(int64(c))
					case reflect.String:
						fmt.Printf("f :%+v\n",f) 
						fl.SetString(f.(string))
				}
			}
		}
	}
	fmt.Printf("%+v\n", structPtr) 

	return structPtr.Interface()

}

