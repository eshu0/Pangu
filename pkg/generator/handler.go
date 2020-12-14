package generator

import (
	"strings"

	pangu "github.com/eshu0/pangu/pkg"
)

type Handler struct {
	pangu.CodeGen
}

func (cg *Handler) GetFileName() string {
	return strings.ToLower(cg.Table.Name)
}
