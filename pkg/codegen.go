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

// solution to having data changes
type Datats struct {
	Templates      []*CodeGen
	Database       *anl.Database
	TargetRepoHost string
	RepoName       string
}

type CodeGen struct {
	PackageName           string
	StorageHandlerName    string
	StorageControllerName string
	TableConstant         *pangudata.Constant
	IdConstant            *pangudata.Constant
	Constants             []*pangudata.Constant
	Table                 *anl.Table
	StructDetails         *pangudata.StructDetails
	InsertDBColumns       string
	UpdateDBColumns       string
	InsertGo              string
	UpdateGo              string
	SelectDBColumns       string
	ParametersColumns     string
	CreateTableSQL        string
	ScanRow               string
	Database              *anl.Database
	TargetRepoHost        string
	RepoName              string
	Filename              string
}

func (cs *CodeGen) getHandlersName() string {
	return strings.ToLower(cs.Table.Name)
}

func (cs *CodeGen) getDataName() string {
	name := strings.ToLower(cs.Table.Name)

	// crude way to take items and make it singular
	if last := len(name) - 1; last >= 0 && name[last] == 's' {
		name = name[:last]
	}

	return name
}

func GenerateFile(dbstruct *anl.DatabaseStructure, slog sli.ISimpleLogger, usetablename bool, repohost string, reponame string) []*CodeGen {

	var temps []*CodeGen

	for _, tbl := range dbstruct.Tables {

		cs := CodeGen{PackageName: "handlers", Table: tbl, StorageHandlerName: strings.Title(tbl.Name + "Handler"), StorageControllerName: strings.Title(tbl.Name + "Controller"), Database: dbstruct.Database, TargetRepoHost: repohost, RepoName: reponame}
		cs.StructDetails = tbl.CreateStructDetails()
		consts, idconst := tbl.CreateConstants()
		if !usetablename {
			cs.Filename = cs.getHandlersName()
		} else {
			cs.Filename = cs.getDataName()
		}
		cs.Constants = consts
		cs.IdConstant = idconst
		cs.CreateTableSQL = strings.Replace(tbl.Sql, "CREATE TABLE", "CREATE TABLE IF NOT EXISTS", -1)

		cnst := &pangudata.Constant{}
		cnst.Comment = fmt.Sprintf("%s", tbl.Name)
		cnst.Name = strings.ToLower(tbl.Name) + strings.Title("TName")
		cnst.Value = tbl.TableName
		cs.TableConstant = cnst

		insertdbcolumns := ""
		updatedbcolumns := ""

		goselect := ""
		goinsert := ""
		goupdate := ""

		parameterscolumns := ""
		for i := 0; i < len(cs.StructDetails.Properties); i++ {
			if i == 0 {
				goselect = fmt.Sprintf("&%s", cs.StructDetails.Properties[i].Name)
			} else {
				goselect = fmt.Sprintf("%s,&%s", goselect, cs.StructDetails.Properties[i].Name)
			}
		}

		startedadd := false
		for i := 0; i < len(cs.StructDetails.Properties); i++ {
			if i == 0 || !startedadd {
				if !cs.StructDetails.Properties[i].IsIdentifier {
					goinsert = fmt.Sprintf("data.%s", cs.StructDetails.Properties[i].Name)
					startedadd = true
				}
			} else {
				if !cs.StructDetails.Properties[i].IsIdentifier {
					goinsert = fmt.Sprintf("%s,data.%s", goinsert, cs.StructDetails.Properties[i].Name)
				}
			}
		}

		// update is different as we want to add the indentifier at the end
		goupdate = goinsert
		for i := 0; i < len(cs.StructDetails.Properties); i++ {
			if cs.StructDetails.Properties[i].IsIdentifier {
				goupdate = fmt.Sprintf("%s,data.%s", goupdate, cs.StructDetails.Properties[i].Name)
			}
		}

		for j := 0; j < len(cs.Constants); j++ {
			if j == 0 {
				insertdbcolumns = fmt.Sprintf("%s", "+ \"[\"+"+cs.Constants[j].Name+"+\"]\" + ")
				updatedbcolumns = fmt.Sprintf("%s", "+ \"[\"+"+cs.Constants[j].Name+"+\"] = ? \" + ")

				parameterscolumns = "?"

			} else {
				insertdbcolumns = insertdbcolumns + fmt.Sprintf("%s", " \",[\"+"+cs.Constants[j].Name+"+\"]\" +")
				updatedbcolumns = updatedbcolumns + fmt.Sprintf("%s", " \",[\"+"+cs.Constants[j].Name+"+\"] = ? \" +")
				parameterscolumns = parameterscolumns + ",?"
			}
		}

		cs.InsertDBColumns = insertdbcolumns
		cs.UpdateDBColumns = updatedbcolumns
		cs.InsertGo = goinsert
		cs.UpdateGo = goupdate

		cs.SelectDBColumns = fmt.Sprintf("%s", "+ \"[\"+"+cs.IdConstant.Name+"+\"],\" ") + insertdbcolumns
		cs.ParametersColumns = parameterscolumns
		cs.ScanRow = goselect

		temps = append(temps, &cs)

	}

	return temps
}

func (cs *CodeGen) CheckCreatePath(slog sli.ISimpleLogger, path string, panicif bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if panicif {
			panic(path + " not found!")
		} else {
			os.Mkdir(path, 0777)
			fmt.Println("Created: " + path)
		}

	} else {
		fmt.Println("Exists: " + path)
	}
}

func (cs *CodeGen) CreateAndExecute(slog sli.ISimpleLogger, filename string, templ *template.Template, data interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		slog.LogError("CreateAndExecute", fmt.Sprintf("Cannot create file%s", err.Error()))
		return
	}

	err = templ.Execute(file, data)
	if err != nil {
		fmt.Println("executing template:", err)
	}

	file.Close()
}

func (cs *CodeGen) CreateTemplate(filepath string, name string) *template.Template {
	b1, err1 := ioutil.ReadFile(filepath) // just pass the file name
	if err1 != nil {
		fmt.Print(err1)
		return nil
	}
	str1 := string(b1) // convert content to a 'string'

	// Create a new template and parse the letter into it.
	return template.Must(template.New(name).Parse(str1))
}
