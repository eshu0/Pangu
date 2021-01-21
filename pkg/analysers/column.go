package sqllite

import (
	_ "github.com/mattn/go-sqlite3"
)

type Column struct {
	PTableName   string
	Name         string
	CType        string
	NotNull      int
	Default      string
	PrimaryKey   int
	DefaultValue []byte
}
