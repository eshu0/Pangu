package main

import (
	"flag"
	"fmt"

	ds "github.com/eshu0/Pangu/examples/Autogen/Todos/DataStore"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

func main() {

	dbname := flag.String("db", "./somedb.db", "Database defaults to ./somedb.db")

	flag.Parse()

	fds := ds.CreateDataStorage(*dbname)

	
	ProjectsHandler := fds.GetProjectsHandler()
	
	JobsHandler := fds.GetJobsHandler()
	
	TasksHandler := fds.GetTasksHandler()
	
	JobHasTasksHandler := fds.GetJobHasTasksHandler()
	
	ProjectHasJobsHandler := fds.GetProjectHasJobsHandler()
	

	fmt.Println("----")
	fmt.Println("Create")
	fmt.Println("----")
	
	newProject  := data.Project{}
	fmt.Println(newProject)

	insProject := ProjectsHandler.Create(newProject)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(insProject)


	if insProject.QuerySucceeded() {
		newProject.Id = insProject.LastInsertId

		fmt.Println("----")
		fmt.Printf("Find By Id %d\n", insProject.LastInsertId)
		fmt.Println("----")
		fresProjectsHandler := ProjectsHandler.FindById(insProject.LastInsertId)
		for _, res1 := range fresProjectsHandler.Results  {	

			fmt.Println("----")
			fmt.Println("Update")
			fmt.Println("----")
			
			res :=data.Project{}
			res = res.ConvertFromIDataItem(res1)
			fmt.Println(res)
			
			newProject.Displayname = "Updated"
			
			newProject.Description = "Updated"
			
			affectedProject := ProjectsHandler.Update(newProject)
			fmt.Println(affectedProject)
		}
	}else{
		fmt.Println("----")
		fmt.Println("Query not succeeded")
		fmt.Println("----")
	}
	
	
	newJob  := data.Job{}
	fmt.Println(newJob)

	insJob := JobsHandler.Create(newJob)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(insJob)


	if insJob.QuerySucceeded() {
		newJob.Id = insJob.LastInsertId

		fmt.Println("----")
		fmt.Printf("Find By Id %d\n", insJob.LastInsertId)
		fmt.Println("----")
		fresJobsHandler := JobsHandler.FindById(insJob.LastInsertId)
		for _, res1 := range fresJobsHandler.Results  {	

			fmt.Println("----")
			fmt.Println("Update")
			fmt.Println("----")
			
			res :=data.Job{}
			res = res.ConvertFromIDataItem(res1)
			fmt.Println(res)
			
			newJob.Displayname = "Updated"
			
			newJob.Description = "Updated"
			
			affectedJob := JobsHandler.Update(newJob)
			fmt.Println(affectedJob)
		}
	}else{
		fmt.Println("----")
		fmt.Println("Query not succeeded")
		fmt.Println("----")
	}
	
	
	newTask  := data.Task{}
	fmt.Println(newTask)

	insTask := TasksHandler.Create(newTask)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(insTask)


	if insTask.QuerySucceeded() {
		newTask.Id = insTask.LastInsertId

		fmt.Println("----")
		fmt.Printf("Find By Id %d\n", insTask.LastInsertId)
		fmt.Println("----")
		fresTasksHandler := TasksHandler.FindById(insTask.LastInsertId)
		for _, res1 := range fresTasksHandler.Results  {	

			fmt.Println("----")
			fmt.Println("Update")
			fmt.Println("----")
			
			res :=data.Task{}
			res = res.ConvertFromIDataItem(res1)
			fmt.Println(res)
			
			newTask.Displayname = "Updated"
			
			affectedTask := TasksHandler.Update(newTask)
			fmt.Println(affectedTask)
		}
	}else{
		fmt.Println("----")
		fmt.Println("Query not succeeded")
		fmt.Println("----")
	}
	
	
	newJobHasTask  := data.JobHasTask{}
	fmt.Println(newJobHasTask)

	insJobHasTask := JobHasTasksHandler.Create(newJobHasTask)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(insJobHasTask)


	if insJobHasTask.QuerySucceeded() {
		newJobHasTask.Id = insJobHasTask.LastInsertId

		fmt.Println("----")
		fmt.Printf("Find By Id %d\n", insJobHasTask.LastInsertId)
		fmt.Println("----")
		fresJobHasTasksHandler := JobHasTasksHandler.FindById(insJobHasTask.LastInsertId)
		for _, res1 := range fresJobHasTasksHandler.Results  {	

			fmt.Println("----")
			fmt.Println("Update")
			fmt.Println("----")
			
			res :=data.JobHasTask{}
			res = res.ConvertFromIDataItem(res1)
			fmt.Println(res)
			
			affectedJobHasTask := JobHasTasksHandler.Update(newJobHasTask)
			fmt.Println(affectedJobHasTask)
		}
	}else{
		fmt.Println("----")
		fmt.Println("Query not succeeded")
		fmt.Println("----")
	}
	
	
	newProjectHasJob  := data.ProjectHasJob{}
	fmt.Println(newProjectHasJob)

	insProjectHasJob := ProjectHasJobsHandler.Create(newProjectHasJob)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(insProjectHasJob)


	if insProjectHasJob.QuerySucceeded() {
		newProjectHasJob.Id = insProjectHasJob.LastInsertId

		fmt.Println("----")
		fmt.Printf("Find By Id %d\n", insProjectHasJob.LastInsertId)
		fmt.Println("----")
		fresProjectHasJobsHandler := ProjectHasJobsHandler.FindById(insProjectHasJob.LastInsertId)
		for _, res1 := range fresProjectHasJobsHandler.Results  {	

			fmt.Println("----")
			fmt.Println("Update")
			fmt.Println("----")
			
			res :=data.ProjectHasJob{}
			res = res.ConvertFromIDataItem(res1)
			fmt.Println(res)
			
			affectedProjectHasJob := ProjectHasJobsHandler.Update(newProjectHasJob)
			fmt.Println(affectedProjectHasJob)
		}
	}else{
		fmt.Println("----")
		fmt.Println("Query not succeeded")
		fmt.Println("----")
	}
	
	

	fmt.Println("----")
	fmt.Println("Update")
	fmt.Println("----")

	
	if  insProject.LastInsertId > 0 {
		newProject.Id = insProject.LastInsertId

		
		
		newProject.Displayname = "Updated"
		
		newProject.Description = "Updated"
		
		affectedProject := ProjectsHandler.Update(newProject)
		fmt.Println(affectedProject)
	} else {

	}
	
	if  insJob.LastInsertId > 0 {
		newJob.Id = insJob.LastInsertId

		
		
		newJob.Displayname = "Updated"
		
		newJob.Description = "Updated"
		
		affectedJob := JobsHandler.Update(newJob)
		fmt.Println(affectedJob)
	} else {

	}
	
	if  insTask.LastInsertId > 0 {
		newTask.Id = insTask.LastInsertId

		
		
		newTask.Displayname = "Updated"
		
		affectedTask := TasksHandler.Update(newTask)
		fmt.Println(affectedTask)
	} else {

	}
	
	if  insJobHasTask.LastInsertId > 0 {
		newJobHasTask.Id = insJobHasTask.LastInsertId

		
		
		affectedJobHasTask := JobHasTasksHandler.Update(newJobHasTask)
		fmt.Println(affectedJobHasTask)
	} else {

	}
	
	if  insProjectHasJob.LastInsertId > 0 {
		newProjectHasJob.Id = insProjectHasJob.LastInsertId

		
		
		affectedProjectHasJob := ProjectHasJobsHandler.Update(newProjectHasJob)
		fmt.Println(affectedProjectHasJob)
	} else {

	}
	

	fmt.Println("----")
	fmt.Println("Get All")
	fmt.Println("----")

	
	resProjectsHandler := ProjectsHandler.ReadAll()
	for _, res := range resProjectsHandler.Results {	
		fmt.Println(res)
	}
	
	resJobsHandler := JobsHandler.ReadAll()
	for _, res := range resJobsHandler.Results {	
		fmt.Println(res)
	}
	
	resTasksHandler := TasksHandler.ReadAll()
	for _, res := range resTasksHandler.Results {	
		fmt.Println(res)
	}
	
	resJobHasTasksHandler := JobHasTasksHandler.ReadAll()
	for _, res := range resJobHasTasksHandler.Results {	
		fmt.Println(res)
	}
	
	resProjectHasJobsHandler := ProjectHasJobsHandler.ReadAll()
	for _, res := range resProjectHasJobsHandler.Results {	
		fmt.Println(res)
	}
	
/*
	fmt.Println("----")
	fmt.Println("Clear All")
	fmt.Println("----")

	
	ProjectsHandlerrowsaffected := ProjectsHandler.Wipe()
	fmt.Println(ProjectsHandlerrowsaffected.RowsAffected)
	
	JobsHandlerrowsaffected := JobsHandler.Wipe()
	fmt.Println(JobsHandlerrowsaffected.RowsAffected)
	
	TasksHandlerrowsaffected := TasksHandler.Wipe()
	fmt.Println(TasksHandlerrowsaffected.RowsAffected)
	
	JobHasTasksHandlerrowsaffected := JobHasTasksHandler.Wipe()
	fmt.Println(JobHasTasksHandlerrowsaffected.RowsAffected)
	
	ProjectHasJobsHandlerrowsaffected := ProjectHasJobsHandler.Wipe()
	fmt.Println(ProjectHasJobsHandlerrowsaffected.RowsAffected)
	
*/
	fmt.Println("----")
	fmt.Println("Get All")
	fmt.Println("----")

	
	resProjectsHandler = ProjectsHandler.ReadAll()
	for _, res := range resProjectsHandler.Results {	
		fmt.Println(res)
	}
	
	resJobsHandler = JobsHandler.ReadAll()
	for _, res := range resJobsHandler.Results {	
		fmt.Println(res)
	}
	
	resTasksHandler = TasksHandler.ReadAll()
	for _, res := range resTasksHandler.Results {	
		fmt.Println(res)
	}
	
	resJobHasTasksHandler = JobHasTasksHandler.ReadAll()
	for _, res := range resJobHasTasksHandler.Results {	
		fmt.Println(res)
	}
	
	resProjectHasJobsHandler = ProjectHasJobsHandler.ReadAll()
	for _, res := range resProjectHasJobsHandler.Results {	
		fmt.Println(res)
	}
	
	
}
