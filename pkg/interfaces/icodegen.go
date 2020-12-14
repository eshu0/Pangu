package pinterface

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

type ICodeGen interface {
	GetFileName() string
	GetTable() *anl.Table
	Create(pkgn string, tbl *anl.Table, database *anl.Database, usetablename bool, repohost string, reponame string) ICodeGen
}
