package sqllite

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"strings"

	sl "github.com/eshu0/simplelogger/pkg"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseAnalyser struct {
	database *sql.DB
	Filename string
	sl.AppLogger
}

type Column struct {
	PTableName string
	Name       string
	CType      string
	NotNull    int
	Default    string
	PrimaryKey int
}

type Database struct {
	Name            string
	Filename        string
	FilenameTrimmed string
}

type Table struct {
	Name      string
	TableName string
	Sql       string
	Columns   []*Column
	HasPK     bool
	PKColumn  *Column
}

type View struct {
	Name      string
	TableName string
	Sql       string
}

type DatabaseStructure struct {
	Tables   []*Table
	Views    []*View
	Database *Database
}

func (daa *DatabaseAnalyser) Create() {
	db, err := sql.Open("sqlite3", daa.Filename)
	if err != nil {
		daa.LogErrorEf("Create", "Create %s ", err)
		daa.database = nil
	} else {
		daa.database = db
	}
}

func (daa *DatabaseAnalyser) GetDatabaseStructure() *DatabaseStructure {
	statement, _ := daa.database.Prepare("SELECT name,type,tbl_name,sql FROM sqlite_master")
	rows, _ := statement.Query()
	dbs := daa.parseStructureRows(rows)
	for _, tbl := range dbs.Tables {
		cols, pk := daa.GetColumns(tbl.TableName)
		tbl.Columns = cols
		tbl.PKColumn = pk
	}

	dbs.Database = daa.GetDatabase()
	return dbs
}

func (daa *DatabaseAnalyser) GetColumns(tablename string) ([]*Column, *Column) {
	statement, _ := daa.database.Prepare("PRAGMA table_info(" + tablename + ")")
	rows, _ := statement.Query() //(tablename)
	return daa.parseTableColumsRows(rows, tablename)
}

func (daa *DatabaseAnalyser) GetDatabase() *Database {
	statement, _ := daa.database.Prepare("PRAGMA database_list")
	rows, _ := statement.Query()
	return daa.parseDBRows(rows)
}

func (daa *DatabaseAnalyser) parseDBRows(rows *sql.Rows) *Database {

	var cId int
	var name string
	var filename string

	db := Database{}

	for rows.Next() {

		rows.Scan(&cId, &name, &filename)

		daa.LogDebug("parseDBRows", "READ: id: "+strconv.Itoa(cId)+"-  name: "+name+" - filename: "+filename)

		db.Name = name
		db.Filename = strings.Title(filepath.Base(filename))
		db.FilenameTrimmed = strings.TrimSuffix(db.Filename, filepath.Ext(db.Filename))
	}

	return &db

}

func (daa *DatabaseAnalyser) parseTableColumsRows(rows *sql.Rows, PTableName string) ([]*Column, *Column) {

	var cId int
	var name string
	var cType string
	var notNull int
	var dftvalue []byte
	var primaryKey int

	var cols []*Column
	var pk *Column

	for rows.Next() {

		rows.Scan(&cId, &name, &cType, &notNull, &dftvalue, &primaryKey)
		daa.LogDebug("parseTableColumsRows", "READ: id: "+strconv.Itoa(cId)+"- type:"+cType+"- notnull:"+strconv.Itoa(notNull)+"- default: name: "+name+"- primaryKey: "+strconv.Itoa(primaryKey))

		col := Column{}

		//
		col.Name = name
		col.CType = cType
		col.NotNull = notNull
		col.PrimaryKey = primaryKey

		//
		col.PTableName = PTableName
		if col.PrimaryKey == 1 {
			pk = &col
		} else {
			cols = append(cols, &col)
		}
	}

	return cols, pk
}

func (daa *DatabaseAnalyser) parseStructureRows(rows *sql.Rows) *DatabaseStructure {

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
