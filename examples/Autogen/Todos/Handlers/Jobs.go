package pguhandlers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Data"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE Jobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)
//

// Table fields

// Jobs
const jobsTName = "Jobs"

// Primay Key: id
const jobsIdCName = "id"


// displayname
const jobsDisplaynameCName = "displayname"

// description
const jobsDescriptionCName = "description"

// archived
const jobsArchivedCName = "archived"

// completed
const jobsCompletedCName = "completed"



// HANDLER

type JobsHandler struct {
	per.IStorageHandler
	Parent *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewJobsHandler(datastore *SQLL.SQLLiteDatastore) *JobsHandler {
	ds := JobsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler 
func (handler *JobsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *JobsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for JobsDBStruct 
func (handler *JobsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.GetLog().LogDebug("CreateStructures","Executing Query")
	return handler.Executor.ExecuteQuery("CREATE TABLE IF NOT EXISTS Jobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)")
}

// End Istorage 

// This function JobsDBStruct removes all data for the table
func (handler *JobsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + jobsTName))
}

// This adds JobsDBStruct to the database 
func (handler *JobsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(JobsDBStruct)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO " + jobsTName + " ( "+ "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +" ) VALUES (?,?,?,?)", data.Displayname,data.Description,data.Archived,data.Completed))
}

func (handler *JobsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult  {
	data := Data.(JobsDBStruct)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE " + jobsTName + " SET "+ "["+jobsDisplaynameCName+"] = ? " +  ",["+jobsDescriptionCName+"] = ? " + ",["+jobsArchivedCName+"] = ? " + ",["+jobsCompletedCName+"] = ? " +"  WHERE [" + jobsIdCName + "] = ?",data.Displayname,data.Description,data.Archived,data.Completed,data.Id))
}

func (handler *JobsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return  SQLL.ResultToSQLLiteQueryResult(data)
}








func (handler *JobsHandler) FindById(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobsIdCName+"]," + "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +"  FROM " + jobsTName + " WHERE " + jobsIdCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *JobsHandler) FindByDisplayname(SearchData string)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobsIdCName+"]," + "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +"  FROM " + jobsTName + " WHERE " + jobsDisplaynameCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *JobsHandler) FindByDescription(SearchData string)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobsIdCName+"]," + "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +"  FROM " + jobsTName + " WHERE " + jobsDescriptionCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *JobsHandler) FindByArchived(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobsIdCName+"]," + "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +"  FROM " + jobsTName + " WHERE " + jobsArchivedCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *JobsHandler) FindByCompleted(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobsIdCName+"]," + "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +"  FROM " + jobsTName + " WHERE " + jobsCompletedCName + " = ?",handler.ParseRows,SearchData))
}




func (handler *JobsHandler) ReadAll()  SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+jobsIdCName+"]," + "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +"  FROM " + jobsTName, handler.ParseRows))
}

func (handler *JobsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	
	var Id int64
	
	var Displayname string
	
	var Description string
	
	var Archived int64
	
	var Completed int64
	
	results := []per.IDataItem{} //JobsDBStruct{}

	for rows.Next() {
		rows.Scan(&Id,&Displayname,&Description,&Archived,&Completed)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)

		res := data.JobsDBStruct{}
		
		res.Id = Id
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Id",Id)
		
		res.Displayname = Displayname
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Displayname",Displayname)
		
		res.Description = Description
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Description",Description)
		
		res.Archived = Archived
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Archived",Archived)
		
		res.Completed = Completed
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Completed",Completed)
		

		results = append(results, res)
	}
	return SQLL.NewDataQueryResult(true,results)
}
