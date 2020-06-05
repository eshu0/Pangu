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
// CREATE TABLE ProjectHasJobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, projectid INTEGER REFERENCES Projects (id) NOT NULL, jobid INTEGER REFERENCES Jobs (id) NOT NULL)
//

// Table fields

// ProjectHasJobs
const projecthasjobsTName = "ProjectHasJobs"

// Primay Key: id
const projecthasjobsIdCName = "id"


// projectid
const projecthasjobsProjectidCName = "projectid"

// jobid
const projecthasjobsJobidCName = "jobid"



// HANDLER

type ProjectHasJobsHandler struct {
	per.IStorageHandler
	Parent *SQLL.SQLLiteDatastore
	Executor *SQLL.SQLLightQueryExecutor
}

func NewProjectHasJobsHandler(datastore *SQLL.SQLLiteDatastore) *ProjectHasJobsHandler {
	ds := ProjectHasJobsHandler{}
	ds.Parent = datastore
	ds.Executor = SQLL.NewSQLLightQueryExecutor(datastore)
	return &ds
}

// Start IStorage Handler 
func (handler *ProjectHasJobsHandler) GetPersistantStorage() per.IPersistantStorage {
	return handler.Parent
}

func (handler *ProjectHasJobsHandler) SetPersistantStorage(persistant per.IPersistantStorage) { // per.IStorageHandler {
	res := persistant.(*SQLL.SQLLiteDatastore)
	handler.Parent = res
	//return handler
}

// This function creates the database table for ProjectHasJob
func (handler *ProjectHasJobsHandler) CreateStructures() per.IQueryResult {
	handler.Parent.GetLog().LogDebug("CreateStructures","Executing Query")
	return handler.Executor.ExecuteQuery("CREATE TABLE IF NOT EXISTS ProjectHasJobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, projectid INTEGER REFERENCES Projects (id) NOT NULL, jobid INTEGER REFERENCES Jobs (id) NOT NULL)")
}

// End Istorage 

// This function ProjectHasJob removes all data for the table
func (handler *ProjectHasJobsHandler) Wipe() SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteQuery("DELETE FROM " + projecthasjobsTName))
}

// This adds ProjectHasJob to the database 
func (handler *ProjectHasJobsHandler) Create(Data per.IDataItem) SQLL.SQLLiteQueryResult {
	data := Data.(data.ProjectHasJob)
	return handler.ConvertResult(handler.Executor.ExecuteInsertQuery("INSERT INTO " + projecthasjobsTName + " ( "+ "["+projecthasjobsProjectidCName+"]" +  ",["+projecthasjobsJobidCName+"]" +" ) VALUES (?,?)", data.Projectid,data.Jobid))
}

func (handler *ProjectHasJobsHandler) Update(Data per.IDataItem) SQLL.SQLLiteQueryResult  {
	data := Data.(data.ProjectHasJob)
	return handler.ConvertResult(handler.Executor.ExecuteQuery("UPDATE " + projecthasjobsTName + " SET "+ "["+projecthasjobsProjectidCName+"] = ? " +  ",["+projecthasjobsJobidCName+"] = ? " +"  WHERE [" + projecthasjobsIdCName + "] = ?",data.Projectid,data.Jobid,data.Id))
}

func (handler *ProjectHasJobsHandler) ConvertResult(data per.IQueryResult) SQLL.SQLLiteQueryResult {
	// this needs to be implemented
	return  SQLL.ResultToSQLLiteQueryResult(data)
}








func (handler *ProjectHasJobsHandler) FindById(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projecthasjobsIdCName+"]," + "["+projecthasjobsProjectidCName+"]" +  ",["+projecthasjobsJobidCName+"]" +"  FROM " + projecthasjobsTName + " WHERE " + projecthasjobsIdCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *ProjectHasJobsHandler) FindByProjectid(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projecthasjobsIdCName+"]," + "["+projecthasjobsProjectidCName+"]" +  ",["+projecthasjobsJobidCName+"]" +"  FROM " + projecthasjobsTName + " WHERE " + projecthasjobsProjectidCName + " = ?",handler.ParseRows,SearchData))
}


func (handler *ProjectHasJobsHandler) FindByJobid(SearchData int64)  SQLL.SQLLiteQueryResult   {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projecthasjobsIdCName+"]," + "["+projecthasjobsProjectidCName+"]" +  ",["+projecthasjobsJobidCName+"]" +"  FROM " + projecthasjobsTName + " WHERE " + projecthasjobsJobidCName + " = ?",handler.ParseRows,SearchData))
}




func (handler *ProjectHasJobsHandler) ReadAll()  SQLL.SQLLiteQueryResult {
	return handler.ConvertResult(handler.Executor.ExecuteResult("SELECT "+ "["+projecthasjobsIdCName+"]," + "["+projecthasjobsProjectidCName+"]" +  ",["+projecthasjobsJobidCName+"]" +"  FROM " + projecthasjobsTName, handler.ParseRows))
}

func (handler *ProjectHasJobsHandler) ParseRows(rows *sql.Rows) per.IQueryResult {
	
	var Id int64
	
	var Projectid int64
	
	var Jobid int64
	
	results := []per.IDataItem{} //ProjectHasJobs{}

	for rows.Next() {
		rows.Scan(&Id,&Projectid,&Jobid)
		//fmt.Println("READ: id: " + string(id) + "- Displayname:"+  displayname + "- Description:" + description)

		res := data.ProjectHasJob{}
		
		res.Id = Id
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Id",Id)
		
		res.Projectid = Projectid
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Projectid",Projectid)
		
		res.Jobid = Jobid
		handler.Parent.GetLog().LogDebugf("ParseRows","Set '%v' for Jobid",Jobid)
		

		results = append(results, res)
	}
	return SQLL.NewDataQueryResult(true,results)
}
