package pinterface

import (
	anl "github.com/eshu0/pangu/pkg/analysers"
)

type ICodeGen interface {
	GetFileName() string
	GetTable() *anl.Table
}
