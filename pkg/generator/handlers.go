package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	anl "github.com/eshu0/pangu/pkg/analysers"
	"github.com/eshu0/pangu/pkg"
)

type Handler struct {
	*pangu.CodeGen
}

func (cg *Handler) () GetFileName string {
	return strings.ToLower(cg.Table.Name)
}


func GenerateHandlers(dbstruct *anl.DatabaseStructure, repohost string, reponame string) []*Handler {
	var temps []*Handler
	//Database Tables
	for _, tbl := range dbstruct.Tables {
		cg := genCG("handlers", tbl, dbstruct.Database, false, repohost, reponame)
		temps = append(temps, cg)
	}
	return temps
}
