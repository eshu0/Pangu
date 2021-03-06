package sqllite

import (
	"fmt"
	"strings"

	pangu "github.com/eshu0/pangu/pkg/structures"
	_ "github.com/mattn/go-sqlite3"
)

type CType int

const (
	Unknown CType = iota
	Integer
	Text
	Numeric
	VarChar
)

/*
func (d CType) String() string {
	return [...]string{"INTEGER", "TEXT", "NUMERIC", "VARCHAR"}[d]
}
*/

func CTypeFromString(typestring string) CType {

	if strings.Contains(typestring, "VARCHAR") {
		return VarChar
	} else {
		switch typestring {
		case "INTEGER":
			return Integer
		case "TEXT":
			return Text
		case "VARCHAR":
			return VarChar
		case "NUMERIC":
			return Numeric
		}
	}

	return Unknown

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

func (col *Column) ToProperty(table *Table) *pangu.Property {

	prop := &pangu.Property{}

	// Comment for the proptyer
	prop.Comment = fmt.Sprintf("%s (SQL TYPE: %s)", col.Name, col.CType)

	// property name
	prop.Name = strings.Title(col.Name)

	// if column is not null we cwill always return a json reponse
	if col.NotNull == 1 {
		prop.Json = "`json:\"" + strings.ToLower(col.Name) + "\"`"
	} else {
		prop.Json = "`json:\"" + strings.ToLower(col.Name) + ",omitempty\"`"
	}

	// this is not an identifier
	prop.IsIdentifier = false
	prop.Constant = table.CreateConstant(col)
	col.SetTypeFromCType(&prop)

	return prop
}

func (col *Column) SetTypeFromCType(prop *pangu.Property) {
	switch col.CType {
	case Integer:
		prop.GType = pangu.GTint64
		prop.UpdateValue = "11"
		break
	case Text:
		prop.GType = pangu.GTstring
		prop.UpdateValue = "\"Updated\""
		break
	case VarChar:
		prop.GType = pangu.GTstring
		prop.UpdateValue = "\"Updated\""
		break
	case Numeric:
		prop.GType = pangu.GTfloat64
		prop.UpdateValue = "1.11"
		break
	}
}
