package sqllite

import (
	"database/sql"
	"strconv"

	sl "github.com/eshu0/simplelogger/pkg"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseAnalyser struct {
	database *sql.DB
	Filename string
	sl.AppLogger
}

type View struct {
	Name      string
	TableName string
	Sql       string
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
	dbs := ParseStructureRows(rows)
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
	return ParseDBRows(rows)
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
		col.CType = CTypeFromString(cType)
		col.NotNull = notNull
		col.DefaultValue = dftvalue
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
