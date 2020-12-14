package generator

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func GenerateHandlers(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Handler {
	var temps []*Handler
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := Create("handlers", tbl, dbstruct.Database, false, repohost, reponame)
		temps = append(temps, cg)
	}
	return temps
}
