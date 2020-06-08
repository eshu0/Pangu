package pgumodel

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	per "github.com/eshu0/persist/pkg/interfaces"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE JobHasTasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, jobid INTEGER REFERENCES Jobs (id) NOT NULL, taskid INTEGER REFERENCES Tasks (id) NOT NULL)
//

// Data storage IDataItem

// Built from: JobHasTasks 
type JobHasTask struct {
	per.IDataItem `json:"-"`


	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// jobid (SQL TYPE: INTEGER)
	Jobid int64 `json:"jobid"`

	// taskid (SQL TYPE: INTEGER)
	Taskid int64 `json:"taskid"`

}

func (data JobHasTask) ConvertFromIDataItem(input per.IDataItem) JobHasTask {
	  res := input.(JobHasTask)
	  return res
}

func (data JobHasTask) Print() string {
	return fmt.Sprintf(" %s ",data) 
}

func (data *JobHasTask) String() string {
	str := ""
	
	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Id) 
	
	// jobid (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Jobid) 
	
	// taskid (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Taskid) 
	
	return str //fmt.Sprintf(" %v, ",data) 
}
