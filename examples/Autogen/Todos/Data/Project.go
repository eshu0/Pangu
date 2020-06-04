package pgudata

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	per "github.com/eshu0/persist/pkg/interfaces"
)

//
// Built from:
// main - Todos.Db
// CREATE TABLE Projects (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, displayname TEXT NOT NULL, description TEXT, archived INTEGER NOT NULL DEFAULT (0), completed INTEGER DEFAULT (0) NOT NULL)
//

// Data storage IDataItem

// Built from: Projects 
type Project struct {
	per.IDataItem


	// id (SQL TYPE: INTEGER)
	Id int64 `id`

	// displayname (SQL TYPE: TEXT)
	Displayname string `displayname`

	// description (SQL TYPE: TEXT)
	Description string `description`

	// archived (SQL TYPE: INTEGER)
	Archived int64 `archived`

	// completed (SQL TYPE: INTEGER)
	Completed int64 `completed`

}

func (data Project) ConvertFromIDataItem(input per.IDataItem) Project {
	  res := input.(Project)
	  return res
}

func (data Project) Print() string {
	return fmt.Sprintf(" %s ",data) 
}

func (data *Project) String() string {
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
