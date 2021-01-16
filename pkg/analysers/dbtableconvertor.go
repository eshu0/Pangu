package sqllite

import (
	"fmt"
	"strings"

	pangu "github.com/eshu0/pangu/pkg/structures"
)

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

	prop := ColumnToProperty(table.PKColumn)
	prop.IsIdentifier = true

	stru.Id = prop

	props = append(props, prop)

	for _, col := range table.Columns {
		props = append(props, ColumnToProperty(col))
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
		switch p.GType {
		case "int64":
			String.Data += "\n\tstr += fmt.Sprintf(\" %d \",data." + p.Name + ")"
		case "string":
			String.Data += "\n\tstr += fmt.Sprintf(\" %s \",data." + p.Name + ")"
			break
		case "float64":
			String.Data += "\n\tstr += fmt.Sprintf(\" %f \",data." + p.Name + ")"
			break
		default:
			String.Data += "\n\tstr += fmt.Sprintf(\" %v \",data." + p.Name + ")"
			break
		}
	}

	String.Data += "\n\treturn str //fmt.Sprintf(\" %v, \",data) \n}"

	functions = append(functions, &ConvertFromIDataItem)
	functions = append(functions, &Print)
	functions = append(functions, &String)
	stru.Functions = functions
	return &stru
}

func (table *Table) ColumnToProperty(col Column) *pangu.Property {

	prop := pangu.Property{}

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

	if strings.Contains(col.CType, "VARCHAR") {
		prop.GType = "string"
		prop.UpdateValue = "\"Updated\""
	} else {
		switch col.CType {
		case "INTEGER":
			prop.GType = "int64"
			prop.UpdateValue = "11"
			props = append(props, prop)
			break
		case "TEXT":
			prop.GType = "string"
			prop.UpdateValue = "\"Updated\""
			props = append(props, prop)
			break
		case "VARCHAR":
			prop.GType = "string"
			prop.UpdateValue = "\"Updated\""
			props = append(props, prop)
			break
		case "NUMERIC":
			prop.GType = "float64"
			prop.UpdateValue = "1.11"
			props = append(props, prop)
			break
		}
	}

	return &prop
}
