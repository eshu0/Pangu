package pangu

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	anl "github.com/eshu0/pangu/pkg/analysers"
	pangudata "github.com/eshu0/pangu/pkg/structures"
	sli "github.com/eshu0/simplelogger/pkg/interfaces"
)

type Handler struct {
	CodeGen
}

func (cg *Handler) () GetFileName string {
	return strings.ToLower(cg.Table.Name)
}
