package main

import (
	"fmt"
	"strings"
	"text/template"
	sl "github.com/eshu0/simplelogger"

)

type CodeTemplate struct {
	PackageName string
	StorageHandlerName string
	TableConstant *Constant
	IdConstant *Constant
	Constants []*Constant
	Table *Table
	StructDetails *StructDetails
	InsertDBColumns string
	UpdateDBColumns string
	InsertGo string
	UpdateGo string
	SelectDBColumns string
	ParametersColumns string
	CreateTableSQL string
	ScanRow string
	Database *Database
}


func GenerateFile(t *template.Template, dbstruct *DatabaseStructure, slog *sl.SimpleLogger) []*CodeTemplate {
	
	var temps []*CodeTemplate

	for _, tbl := range dbstruct.Tables {

		cs := CodeTemplate { PackageName: "main", Table: tbl , StorageHandlerName: strings.Title(tbl.Name+"Handler"), Database : dbstruct.Database }
		cs.StructDetails = tbl.CreateStructDetails()
		consts, idconst := tbl.CreateConstants()

		cs.Constants = consts
		cs.IdConstant = idconst
		cs.CreateTableSQL = strings.Replace(tbl.Sql, "CREATE TABLE", "CREATE TABLE IF NOT EXISTS", -1) 
		
		cnst := &Constant{}
		cnst.Comment = fmt.Sprintf("%s",tbl.Name) 
		cnst.Name = strings.ToLower(tbl.Name) + strings.Title("TName")
		cnst.Value = tbl.TableName
		cs.TableConstant = cnst

		insertdbcolumns := ""
		updatedbcolumns := ""

		goselect := ""
		goinsert := ""
		goupdate := ""

		parameterscolumns := ""
		for i := 0; i < len(cs.StructDetails.Properties); i++ {
			if i == 0 {
				goselect = fmt.Sprintf("&%s",cs.StructDetails.Properties[i].Name) 
			} else {
				goselect = fmt.Sprintf("%s,&%s",goselect,cs.StructDetails.Properties[i].Name) 
			}
		}	

		startedadd := false
		for i := 0; i < len(cs.StructDetails.Properties); i++ {
			if i == 0 || !startedadd{
				if !cs.StructDetails.Properties[i].IsIdentifier {
					goinsert = fmt.Sprintf("data.%s",cs.StructDetails.Properties[i].Name) 
					startedadd = true
				}
			} else {
				if !cs.StructDetails.Properties[i].IsIdentifier {
					goinsert = fmt.Sprintf("%s,data.%s",goinsert,cs.StructDetails.Properties[i].Name) 
				}
			}
		}	

		// update is different as we want to add the indentifier at the end
		goupdate = goinsert
		for i := 0; i < len(cs.StructDetails.Properties); i++ {
			if cs.StructDetails.Properties[i].IsIdentifier {
				goupdate = fmt.Sprintf("%s,data.%s",goupdate,cs.StructDetails.Properties[i].Name) 
			}
		}	

		for j := 0; j < len(cs.Constants); j++ {
			if j == 0 {
				insertdbcolumns = fmt.Sprintf("%s","+ \"[\"+"+cs.Constants[j].Name+"+\"]\" + ") 
				updatedbcolumns = fmt.Sprintf("%s","+ \"[\"+"+cs.Constants[j].Name+"+\"] = ? \" + ") 

				parameterscolumns = "?" 

			}else{
				insertdbcolumns = insertdbcolumns + fmt.Sprintf("%s"," \",[\"+" + cs.Constants[j].Name + "+\"]\" +") 
				updatedbcolumns  = updatedbcolumns + fmt.Sprintf("%s"," \",[\"+" + cs.Constants[j].Name + "+\"] = ? \" +") 
				parameterscolumns = parameterscolumns + ",?" 
			}
		}	

		cs.InsertDBColumns = insertdbcolumns
		cs.UpdateDBColumns = updatedbcolumns
		cs.InsertGo = goinsert
		cs.UpdateGo = goupdate

		cs.SelectDBColumns = fmt.Sprintf("%s","+ \"[\"+" + cs.IdConstant.Name + "+\"],\" ")  + insertdbcolumns
		cs.ParametersColumns = parameterscolumns
		cs.ScanRow = goselect

		temps = append(temps, &cs)	

	}

	return temps
}
