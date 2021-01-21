package sqllite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseStructure struct {
	Tables   []*Table
	Views    []*View
	Database *Database
}

func ParseStructureRows(rows *sql.Rows) *DatabaseStructure {

	var name string
	var tblname string
	var sql string
	var ttype string

	var tables []*Table
	var views []*View
	dbs := &DatabaseStructure{}

	for rows.Next() {

		rows.Scan(&name, &ttype, &tblname, &sql)
		if name != "sqlite_sequence" {

			if ttype == "table" {
				tbl := Table{}
				tbl.Name = name
				tbl.TableName = tblname
				tbl.Sql = sql
				tables = append(tables, &tbl)
			}

			if ttype == "view" {
				viw := View{}
				viw.Name = name
				viw.TableName = tblname
				viw.Sql = sql
				views = append(views, &viw)
			}
		}

	}

	dbs.Tables = tables
	dbs.Views = views

	return dbs
}
