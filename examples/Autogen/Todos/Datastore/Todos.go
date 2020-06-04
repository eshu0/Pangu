package pgudatastore

import (
	sli "github.com/eshu0/simplelogger/interfaces"
	SQLL "github.com/eshu0/persist/pkg/sqllite"
	//per "github.com/eshu0/persist/pkg/interfaces"
)


type TodosDatastore struct {
	Datastore *SQLL.SQLLiteDatastore

	//
	//ProjectsHandler *ProjectsHandler
	//
	//JobsHandler *JobsHandler
	//
	//TasksHandler *TasksHandler
	//
	//JobHasTasksHandler *JobHasTasksHandler
	//
	//ProjectHasJobsHandler *ProjectHasJobsHandler
	//
}

func CreateDataStorage(log sli.ISimpleLogger,filename string) *TodosDatastore {
	ds := SQLL.CreateOpenSQLLiteDatastore(log, filename)
	
	// tests the example
	ds.SetStorageHander("Generic",SQLL.NewSQLLiteTableHandler(ds)) 
	
	ds.SetStorageHander("Projects",NewProjectsHandler(ds))
	
	ds.SetStorageHander("Jobs",NewJobsHandler(ds))
	
	ds.SetStorageHander("Tasks",NewTasksHandler(ds))
	
	ds.SetStorageHander("JobHasTasks",NewJobHasTasksHandler(ds))
	
	ds.SetStorageHander("ProjectHasJobs",NewProjectHasJobsHandler(ds))
	


	ds.CreateStructures()

	res := TodosDatastore{}
	res.Datastore = ds
	
	return &res
}


func (fds *TodosDatastore) GetProjectsHandler() *ProjectsHandler {

	data, ok := fds.Datastore.GetStorageHandler("Projects")
	if ok {
	  res := data.(*ProjectsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetJobsHandler() *JobsHandler {

	data, ok := fds.Datastore.GetStorageHandler("Jobs")
	if ok {
	  res := data.(*JobsHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetTasksHandler() *TasksHandler {

	data, ok := fds.Datastore.GetStorageHandler("Tasks")
	if ok {
	  res := data.(*TasksHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetJobHasTasksHandler() *JobHasTasksHandler {

	data, ok := fds.Datastore.GetStorageHandler("JobHasTasks")
	if ok {
	  res := data.(*JobHasTasksHandler)
	  return res
	}
	return nil
}

func (fds *TodosDatastore) GetProjectHasJobsHandler() *ProjectHasJobsHandler {

	data, ok := fds.Datastore.GetStorageHandler("ProjectHasJobs")
	if ok {
	  res := data.(*ProjectHasJobsHandler)
	  return res
	}
	return nil
}



