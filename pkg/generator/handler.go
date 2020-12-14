package generator

import (
	"strings"

	anl "github.com/eshu0/pangu/pkg/analysers"
)

type Handler struct {
	pinterface.ICodeGen
	Table *anl.Table
}

func (hndlr *Handler) GetFileName() string {
	return strings.ToLower(hndlr.Table.Name)
}
