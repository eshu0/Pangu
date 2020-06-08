package pgumodel

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	per "github.com/eshu0/persist/pkg/interfaces"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE Jobs (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)
//

// Data storage IDataItem

// Built from: Jobs 
type Job struct {
	per.IDataItem `json:"-"`


	// id (SQL TYPE: INTEGER)
	Id int64 `json:"id"`

	// displayname (SQL TYPE: TEXT)
	Displayname string `json:"displayname"`

	// description (SQL TYPE: TEXT)
	Description string `json:"description,omitempty"`

	// archived (SQL TYPE: INTEGER)
	Archived int64 `json:"archived"`

	// completed (SQL TYPE: INTEGER)
	Completed int64 `json:"completed"`

}

func (data Job) ConvertFromIDataItem(input per.IDataItem) Job {
	  res := input.(Job)
	  return res
}

func (data Job) Print() string {
	return fmt.Sprintf(" %s ",data) 
}

func (data *Job) String() string {
	str := ""
	
	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Id) 
	
	// displayname (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ",data.Displayname) 
	
	// description (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ",data.Description) 
	
	// archived (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Archived) 
	
	// completed (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Completed) 
	
	return str //fmt.Sprintf(" %v, ",data) 
}
