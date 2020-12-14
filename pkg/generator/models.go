package generator

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func GenerateModels(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Model {
	var temps []*Model
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := genCG("models", tbl, dbstruct.Database, true, repohost, reponame)
		temps = append(temps, cg)
	}
	return temps
}
