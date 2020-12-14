package generator

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func GenerateModels(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Model {
	var temps []*Model
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := create("models", tbl, dbstruct.Database, true, repohost, reponame)
		mod := Model{} //cg.(*Model)
		mod.CodeGen = cg
		mod.Filename = mod.GetFileName()

		temps = append(temps, &mod)
	}
	return temps
}
