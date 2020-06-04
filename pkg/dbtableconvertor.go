
package pangu

import (
	"fmt"
	"strings"
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func (table *anl.Table) CreateConstants() ([]*Constant, *Constant) {
	var cnsts []*Constant
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


func (table *anl.Table) CreateConstant(col *anl.Column)  *Constant {

	cnst := &Constant{}
	cnst.Comment = fmt.Sprintf("%s",col.Name) 
	cnst.Name = strings.ToLower(table.Name) + strings.Title(col.Name + "CName")
	cnst.Value = col.Name

	return cnst
}


func (table *anl.Table) CreateStructDetails() *StructDetails {
	stru := StructDetails { Name: strings.Title(table.Name+"DBStruct") }
	stru.Comment = fmt.Sprintf("Built from: %s",table.Name) 

	var props []*Property
	var uprops []*Property

	prop := &Property{}
	prop.Comment = fmt.Sprintf("%s (SQL TYPE: %s)",table.PKColumn.Name,table.PKColumn.CType) 
	prop.Name = strings.Title(table.PKColumn.Name)
	prop.Json = "`"+strings.ToLower(table.PKColumn.Name)+"`"
	prop.GType = "int64"
	prop.IsIdentifier = true
	prop.Constant = table.CreateConstant(table.PKColumn)
	stru.Id = prop
	props = append(props, prop)	

	for _, col := range table.Columns {
		prop := &Property{}
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