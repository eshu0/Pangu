package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	anl "github.com/eshu0/pangu/pkg/analysers"
	"github.com/eshu0/pangu/pkg/"
)

type Model struct {
	*pangu.CodeGen
}

func (cg *Model) () GetFileName string {
	name := strings.ToLower(cs.Table.Name)

	// crude way to take items and make it singular
	if last := len(name) - 1; last >= 0 && name[last] == 's' {
		name = name[:last]
	}

	return name
}

func GenerateModels(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Model {
	var temps []*Model
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := genCG("models", tbl, dbstruct.Database, true, repohost, reponame)
		temps = append(temps, cg)
	}
	return temps
}
