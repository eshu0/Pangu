package sqllite

import (
	"fmt"
	"strings"

	pangu "github.com/eshu0/pangu/pkg/structures"
)

type Table struct {
	Name      string
	TableName string
	Sql       string
	Columns   []*Column
	HasPK     bool
	PKColumn  *Column
}

func (table *Table) CreateConstants() ([]*pangu.Constant, *pangu.Constant) {
	var cnsts []*pangu.Constant
	//pk := &Constant{}
	//pk.Name = strings.ToLower(table.Name) + strings.Title(table.PKColumn.Name + "CName")
	//pk.Value = table.PKColumn.Name
	pk := table.CreateConstant(table.PKColumn)
	pk.Comment = fmt.Sprintf("Primay Key: %s", table.PKColumn.Name)

	for _, col := range table.Columns {
		cnst := table.CreateConstant(col)
		cnsts = append(cnsts, cnst)
	}

	return cnsts, pk
}

func (table *Table) CreateConstant(col *Column) *pangu.Constant {

	cnst := &pangu.Constant{}
	cnst.Comment = fmt.Sprintf("%s", col.Name)
	cnst.Name = strings.ToLower(table.Name) + strings.Title(col.Name+"CName")
	cnst.Value = col.Name

	return cnst
}

func (table *Table) CreateStructDetails() *pangu.StructDetails {
	name := strings.Title(table.Name)
	if last := len(name) - 1; last >= 0 && name[last] == 's' {
		name = name[:last]
	}
	stru := pangu.StructDetails{Name: name}
	stru.Comment = fmt.Sprintf("Built from: %s", table.Name)

	var props []*pangu.Property
	var functions []*pangu.Function

	prop := table.PKColumn.ToProperty(table)
	prop.IsIdentifier = true

	stru.ID = prop

	props = append(props, prop)

	for _, col := range table.Columns {
		props = append(props, col.ToProperty(table))
	}

	stru.Properties = props

	ConvertFromIDataItem := pangu.Function{}
	ConvertFromIDataItem.Data = fmt.Sprintf("func (data *%s) ConvertFromIDataItem(input per.IDataItem) %s { \n\tres := input.(%s) \n\treturn res \n}", stru.Name, stru.Name, stru.Name)

	Print := pangu.Function{}
	Print.Data = "func (data " + stru.Name + ") Print() string { \n\t return fmt.Sprintf(\"%s\",data) \n}"

	String := pangu.Function{}
	String.Data = "func (data *" + stru.Name + ") String() string {\n\t str := \"\""

	for _, p := range stru.Properties {
		String.Data += "\n\t// " + p.Comment
		String.Data += p.PrintString()
	}

	String.Data += "\n\treturn str //fmt.Sprintf(\" %v, \",data) \n}"

	functions = append(functions, &ConvertFromIDataItem)
	functions = append(functions, &Print)
	functions = append(functions, &String)
	stru.Functions = functions
	return &stru
}
