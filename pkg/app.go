package pangu

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	anl "github.com/eshu0/pangu/pkg/analysers"
	sl "github.com/eshu0/simplelogger/pkg"
)

type PanguApp struct {
	sl.AppLogger
}

func (pa *PanguApp) CheckCreatePath(path string, panicif bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if panicif {
			panic(path + " not found!")
		} else {
			os.Mkdir(path, 0777)
			pa.LogDebug("CheckCreatePath", "Created: "+path)
		}

	} else {
		pa.LogDebug("CheckCreatePath", "Exists: "+path)
	}
}

func (pa *PanguApp) CreateAndExecute(filename string, templ *template.Template, data interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		pa.LogError("CreateAndExecute", fmt.Sprintf("Cannot create file%s", err.Error()))
		return
	}

	err = templ.Execute(file, data)
	if err != nil {
		pa.LogErrorEf("CreateAndExecute", "executing template: %s", err)
	}

	file.Close()
}

func (pa *PanguApp) Parse(dbname string, odir string) {

	dbfolder := strings.Replace(filepath.Base(dbname), filepath.Ext(dbname), "", -1)

	outputdir := odir + strings.Title(dbfolder)
	fmt.Println("Outputting to: " + outputdir)

	datastoredir := outputdir + "/Datastore/"
	handlerdir := outputdir + "/Handlers/"
	modelsdir := outputdir + "/Models/"
	appdir := outputdir + "/TestApp/"
	restdir := outputdir + "/REST/"
	controllersdir := restdir + "/Controllers/"

	pa.CheckCreatePath(dbname, true)
	pa.CheckCreatePath(odir, false)
	pa.CheckCreatePath(outputdir, false)
	pa.CheckCreatePath(datastoredir, false)
	pa.CheckCreatePath(handlerdir, false)
	pa.CheckCreatePath(modelsdir, false)
	pa.CheckCreatePath(appdir, false)
	pa.CheckCreatePath(controllersdir, false)
	pa.CheckCreatePath(restdir, false)

	fds := &anl.DatabaseAnalyser{}
	fds.Filename = dbname
	fds.Create(pa)

	dbstruct := fds.GetDatabaseStructure()

	CodeTemplate := pa.CreateTemplate("./Templates/CodeTemplate.txt", "code")
	DataTemplate := pa.CreateTemplate("./Templates/DataTemplate.txt", "data")
	DLTemplate := pa.CreateTemplate("./Templates/DLTemplate.txt", "dl")
	MainTemplate := pa.CreateTemplate("./Templates/MainTemplate.txt", "main")
	ControllersTemplate := pa.CreateTemplate("./Templates/Controllers.txt", "control")
	RESTServerTemplate := pa.CreateTemplate("./Templates/RESTServer.txt", "control")

	// Execute the template for each recipient.
	ctemplates := GenerateFile(dbstruct)

	for _, cs := range ctemplates {
		pa.CreateAndExecute(handlerdir+cs.GetHandlersName()+".go", CodeTemplate, cs)
		pa.CreateAndExecute(controllersdir+cs.GetHandlersName()+".go", ControllersTemplate, cs)
		pa.CreateAndExecute(modelsdir+cs.GetDataName()+".go", DataTemplate, cs)
	}

	dl := Datats{Database: ctemplates[0].Database, Templates: ctemplates}

	pa.CreateAndExecute(datastoredir+dl.Database.FilenameTrimmed+".go", DLTemplate, dl)
	pa.CreateAndExecute(appdir+"main.go", MainTemplate, ctemplates)
	pa.CreateAndExecute(restdir+"main.go", RESTServerTemplate, ctemplates)

}

func (pa *PanguApp) CreateTemplate(filepath string, name string) *template.Template {
	b1, err1 := ioutil.ReadFile(filepath) // just pass the file name
	if err1 != nil {
		fmt.Print(err1)
		return nil
	}
	str1 := string(b1) // convert content to a 'string'

	// Create a new template and parse the letter into it.
	return template.Must(template.New(name).Parse(str1))
}
