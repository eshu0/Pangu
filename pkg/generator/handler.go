package generator

import (
	"strings"
)

type Handler struct {
	//pinterface.ICodeGen
	//Table *anl.Table
	CodeGen
}

func (hndlr *Handler) GetFileName() string {
	return strings.ToLower(hndlr.Table.Name)
}
