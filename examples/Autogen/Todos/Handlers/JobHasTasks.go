package pguhandlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE JobHasTasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, jobid INTEGER REFERENCES Jobs (id) NOT NULL, taskid INTEGER REFERENCES Tasks (id) NOT NULL)
//

// Table fields

// JobHasTasks
const jobhastasksTName = "JobHasTasks"

// Primay Key: id
const jobhastasksIdCName = "id"


// jobid
const jobhastasksJobidCName = "jobid"

// taskid
const jobhastasksTaskidCName = "taskid"



// HANDLER

type JobHasTasksHandler struct {
	per.IStorageHandler
	Parent *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewJobHasTasksHandler(datastore *SQLL.SQLLiteDatastore) *JobHasTasksHandler {
	ds := JobHasTasksHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler 
func (handler *JobHasTasksHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *JobHasTasksHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for JobHasTask 
func (handler *JobHasTasksHandler) CreateStructures() per.IQueryResult {
	handler.Parent.GetLog().LogDebug("CreateStructures","Executing Query")
	return handler.Executor.ExecuteQuery("CREATE TABLE IF NOT EXISTS JobHasTasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, jobid INTEGER REFERENCES Jobs (id) NOT NULL, taskid INTEGER REFERENCES Tasks (id) NOT NULL)")
}

// End Istorage 

// This function JobHasTask removes all data for the table
func (handler *JobHasTasksHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + jobhastasksTName))
}

// This adds JobHasTask to the database 
func (handler *JobHasTasksHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.JobHasTask)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO " + jobhastasksTName + " ( "+ "["+jobhastasksJobidCName+"]" +  ",["+jobhastasksTaskidCName+"]" +" ) VALUES (?,?)", data.Jobid,data.Taskid))
}

func (handler *JobHasTasksHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult  {
	data := Data.(data.JobHasTask)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE " + jobhastasksTName + " SET "+ "["+jobhastasksJobidCName+"] = ? " +  ",["+jobhastasksTaskidCName+"] = ? " +"  WHERE [" + jobhastasksIdCName + "] = ?",data.Jobid,data.Taskid,data.Id))
}

func (handler *JobHasTasksHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return  SQLL.ResultToSQLLiteQueryResult(data)
}








func (handler *JobHasTasksHandler) FindById(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobhastasksIdCName+"]," + "["+jobhastasksJobidCName+"]" +  ",["+jobhastasksTaskidCName+"]" +"  FROM " + jobhastasksTName + " WHERE " + jobhastasksIdCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *JobHasTasksHandler) FindByJobid(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobhastasksIdCName+"]," + "["+jobhastasksJobidCName+"]" +  ",["+jobhastasksTaskidCName+"]" +"  FROM " + jobhastasksTName + " WHERE " + jobhastasksJobidCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *JobHasTasksHandler) FindByTaskid(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobhastasksIdCName+"]," + "["+jobhastasksJobidCName+"]" +  ",["+jobhastasksTaskidCName+"]" +"  FROM " + jobhastasksTName + " WHERE " + jobhastasksTaskidCName + " = ?",handler.ParseRows,SearchData))
}




func (handler *JobHasTasksHandler) ReadAll()  SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobhastasksIdCName+"]," + "["+jobhastasksJobidCName+"]" +  ",["+jobhastasksTaskidCName+"]" +"  FROM " + jobhastasksTName, handler.ParseRows))
}

func (handler *JobHasTasksHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	
	var Id int64
	
	var Jobid int64
	
	var Taskid int64
	
	results := []per.IDataItem{} //JobHasTask{}

	for rows.Next() {
		rows.Scan(&Id,&Jobid,&Taskid)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)

		res := data.JobHasTask{}
		
		res.Id = Id
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Id",Id)
		
		res.Jobid = Jobid
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Jobid",Jobid)
		
		res.Taskid = Taskid
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Taskid",Taskid)
		

		results = append(results, res)
	}
	return SQLL.NewDataQueryResult(true,results)
}
