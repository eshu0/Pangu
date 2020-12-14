package generator

import (
	"strings"

	anl "github.com/eshu0/pangu/pkg/analysers"
)

type Model struct {
	pinterface.ICodeGen
	Table *anl.Table
}

func (cg *Model) GetFileName() string {
	name := strings.ToLower(cg.Table.Name)

	// crude way to take items and make it singular
	if last := len(name) - 1; last >= 0 && name[last] == 's' {
		name = name[:last]
	}

	return name
}
