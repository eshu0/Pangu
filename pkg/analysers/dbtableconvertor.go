package sqllite 

import (
	"fmt"
	"strings"
	pangu "github.com/eshu0/pangu/pkg"
)

func (table *Table) CreateConstants() ([]*pangu.Constant, *pangu.Constant) {
	var cnsts []*pangu.Constant
	//pk := &Constant{}
	//pk.Name = strings.ToLower(table.Name) + strings.Title(table.PKColumn.Name + "CName")
	//pk.Value = table.PKColumn.Name
	pk := table.CreateConstant(table.PKColumn)
	pk.Comment = fmt.Sprintf("Primay Key: %s",table.PKColumn.Name) 

	for _, col := range table.Columns {
		cnst := table.CreateConstant(col)
		cnsts = append(cnsts, cnst)	
	}
	
	return cnsts,pk
}


func (table *Table) CreateConstant(col *Column)  *pangu.Constant {

	cnst := &pangu.Constant{}
	cnst.Comment = fmt.Sprintf("%s",col.Name) 
	cnst.Name = strings.ToLower(table.Name) + strings.Title(col.Name + "CName")
	cnst.Value = col.Name

	return cnst
}


func (table *Table) CreateStructDetails() *pangu.StructDetails {
	stru := pangu.StructDetails { Name: strings.Title(table.Name) }
	stru.Comment = fmt.Sprintf("Built from: %s",table.Name) 

	var props []*pangu.Property
	var uprops []*pangu.Property

	prop := &pangu.Property{}
	prop.Comment = fmt.Sprintf("%s (SQL TYPE: %s)",table.PKColumn.Name,table.PKColumn.CType) 
	prop.Name = strings.Title(table.PKColumn.Name)
	prop.Json = "`"+strings.ToLower(table.PKColumn.Name)+"`"
	prop.GType = "int64"
	prop.IsIdentifier = true
	prop.Constant = table.CreateConstant(table.PKColumn)
	stru.Id = prop
	props = append(props, prop)	

	for _, col := range table.Columns {
		prop := &pangu.Property{}
		prop.Comment = fmt.Sprintf("%s (SQL TYPE: %s)",col.Name,col.CType) 
		prop.Name = strings.Title(col.Name)
		prop.Json = "`"+strings.ToLower(col.Name)+"`"
		prop.IsIdentifier = false
		prop.Constant = table.CreateConstant(col)

		switch col.CType {
			case "INTEGER" :
				prop.GType = "int64"
				props = append(props, prop)	
				break
			case "TEXT" :
				prop.GType = "string"
				props = append(props, prop)	
				uprops = append(uprops, prop)	
				break
		}
	}
	

	stru.Properties = props
	stru.UpdateProperties = uprops

	return &stru
}