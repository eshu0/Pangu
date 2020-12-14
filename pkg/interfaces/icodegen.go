package pinterface

type ICodeGen interface {
	GetFileName() string
	//GetTableName() string
	//Create(pkgn string, tbl *anl.Table, database *anl.Database, usetablename bool, repohost string, reponame string) ICodeGen
}
