package generator

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func GenerateHandlers(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Handler {
	var temps []*Handler
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := create("handlers", tbl, dbstruct.Database, false, repohost, reponame)
		hndl := Handler{} //cg.(*Handler)
		hndl.CodeGen = cg

		temps = append(temps, &hndl)
	}
	return temps
}
