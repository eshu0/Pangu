package sqllite

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Name            string
	Filename        string
	FilenameTrimmed string
}

func ParseDBRows(daa *DatabaseAnalyser, rows *sql.Rows) *Database {

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
