package sqllite

import (
	_ "github.com/mattn/go-sqlite3"
)

type CType int

const (
	Integer CType = iota
	Text
	Numeric
)

func (d CType) String() string {
	return [...]string{"INTEGER", "TEXT", "NUMERIC"}[d]
}

type Column struct {
	PTableName   string
	Name         string
	CType        CType
	NotNull      int
	Default      string
	PrimaryKey   int
	DefaultValue []byte
}
