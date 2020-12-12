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
/*
 CREATE TABLE Jobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)
 */
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

// This function creates the database table for Job 
func (handler *JobsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.LogDebug("CreateStructures","Executing Query")
	return handler.Executor.ExecuteQuery(`CREATE TABLE IF NOT EXISTS Jobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)`)
}

// End Istorage 

// This function Job removes all data for the table
func (handler *JobsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + jobsTName))
}

// This adds Job to the database 
func (handler *JobsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.Job)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO " + jobsTName + " ( "+ "["+jobsDisplaynameCName+"]" +  ",["+jobsDescriptionCName+"]" + ",["+jobsArchivedCName+"]" + ",["+jobsCompletedCName+"]" +" ) VALUES (?,?,?,?)", data.Displayname,data.Description,data.Archived,data.Completed))
}

func (handler *JobsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult  {
	data := Data.(data.Job)
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
	
	var Id *int64
	
	var Displayname *string
	
	var Description *string
	
	var Archived *int64
	
	var Completed *int64
	
	results := []per.IDataItem{} //Job{}

	for rows.Next() {
		err := rows.Scan(&Id,&Displayname,&Description,&Archived,&Completed)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)
		if err != nil {
			handler.Parent.LogErrorEf("ParseRows","Row Scan errr: %s ",err)
		} else {
			res := data.Job{}
			
				if Id != nil {
					res.Id = *Id
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Id",*Id)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Displayname != nil {
					res.Displayname = *Displayname
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Displayname",*Displayname)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Description != nil {
					res.Description = *Description
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Description",*Description)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Archived != nil {
					res.Archived = *Archived
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Archived",*Archived)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			
				if Completed != nil {
					res.Completed = *Completed
					handler.Parent.LogDebugf("ParseRows","Set '%v' for Completed",*Completed)
				} else {
					handler.Parent.LogDebugf("ParseRows","{.Name}} was NULL")
				}
			

			results = append(results, res)
		}

	}
	return SQLL.NewDataQueryResult(true,results)
}
