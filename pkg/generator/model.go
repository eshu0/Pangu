package generator

import (
	"strings"

	pangu "github.com/eshu0/pangu/pkg"
)

type Model struct {
	pangu.CodeGen
}

func (cg *Model) GetFileName() string {
	name := strings.ToLower(cg.Table.Name)

	// crude way to take items and make it singular
	if last := len(name) - 1; last >= 0 && name[last] == 's' {
		name = name[:last]
	}

	return name
}
