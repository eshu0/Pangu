package pgudata

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	per "github.com/eshu0/persist/pkg/interfaces"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE ProjectHasJobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, projectid INTEGER REFERENCES Projects (id) NOT NULL, jobid INTEGER REFERENCES Jobs (id) NOT NULL)
//

// Data storage IDataItem

// Built from: ProjectHasJobs 
type ProjectHasJob struct {
	per.IDataItem


	// id (SQL TYPE: INTEGER)
	Id int64 `id`

	// projectid (SQL TYPE: INTEGER)
	Projectid int64 `projectid`

	// jobid (SQL TYPE: INTEGER)
	Jobid int64 `jobid`

}

func (data ProjectHasJob) ConvertFromIDataItem(input per.IDataItem) ProjectHasJob {
	  res := input.(ProjectHasJob)
	  return res
}

func (data ProjectHasJob) Print() string {
	return fmt.Sprintf(" %s ",data) 
}

func (data *ProjectHasJob) String() string {
	str := ""
	
	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Id) 
	
	// projectid (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Projectid) 
	
	// jobid (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Jobid) 
	
	return str //fmt.Sprintf(" %v, ",data) 
}
