package pgumodel

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	per "github.com/eshu0/persist/pkg/interfaces"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE Tasks (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, archived INTEGER DEFAULT (0) NOT NULL, completed INTEGER DEFAULT (0) NOT NULL)
//

// Data storage IDataItem

// Built from: Tasks 
type Task struct {
	per.IDataItem


	// id (SQL TYPE: INTEGER)
	Id int64 `json:id`

	// displayname (SQL TYPE: TEXT)
	Displayname string `displayname`

	// archived (SQL TYPE: INTEGER)
	Archived int64 `archived`

	// completed (SQL TYPE: INTEGER)
	Completed int64 `completed`

}

func (data Task) ConvertFromIDataItem(input per.IDataItem) Task {
	  res := input.(Task)
	  return res
}

func (data Task) Print() string {
	return fmt.Sprintf(" %s ",data) 
}

func (data *Task) String() string {
	str := ""
	
	// id (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Id) 
	
	// displayname (SQL TYPE: TEXT)
	str = str + fmt.Sprintf(" %s ",data.Displayname) 
	
	// archived (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Archived) 
	
	// completed (SQL TYPE: INTEGER)
	str = str + fmt.Sprintf(" %s ",data.Completed) 
	
	return str //fmt.Sprintf(" %v, ",data) 
}
