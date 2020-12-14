package generator

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func GenerateModels(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Model {
	var temps []*Model
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := Create("models", tbl, dbstruct.Database, true, repohost, reponame)
		mod := Model{} //cg.(*Model)
		mod.CodeGen = cg
		temps = append(temps, &mod)
	}
	return temps
}
