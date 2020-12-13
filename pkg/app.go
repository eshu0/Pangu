package pangu

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	anl "github.com/eshu0/pangu/pkg/analysers"
	"github.com/eshu0/pangu/pkg/generator"

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

func (pa *PanguApp) Parse(dbname string, odir string, tdir string) {

	if tdir == "" {
		tdir = "./templates/"
	}

	dbfolder := strings.Replace(filepath.Base(dbname), filepath.Ext(dbname), "", -1)

	outputdir := odir + strings.ToLower(dbfolder) //strings.Title(dbfolder)
	fmt.Println("Outputting to: " + outputdir)

	pkgdir := outputdir + "/pkg"
	fmt.Println("Package directory is: " + outputdir)

	// packages
	datastoredir := pkgdir + "/datastore/"
	handlerdir := pkgdir + "/handlers/"
	modelsdir := pkgdir + "/models/"
	pkgrestdir := pkgdir + "/REST/"
	controllersdir := pkgrestdir + "controllers/"

	// apps
	cmddir := outputdir + "/cmd/"
	appdir := cmddir + "testapp/"
	restdir := cmddir + "restserver/"

	pa.CheckCreatePath(dbname, true)
	pa.CheckCreatePath(odir, false)
	pa.CheckCreatePath(pkgdir, false)
	pa.CheckCreatePath(cmddir, false)

	pa.CheckCreatePath(pkgrestdir, false)
	pa.CheckCreatePath(outputdir, false)
	pa.CheckCreatePath(datastoredir, false)
	pa.CheckCreatePath(handlerdir, false)
	pa.CheckCreatePath(modelsdir, false)
	pa.CheckCreatePath(appdir, false)
	pa.CheckCreatePath(controllersdir, false)
	pa.CheckCreatePath(restdir, false)

	fds := &anl.DatabaseAnalyser{}
	fds.Filename = dbname
	fds.Create()

	dbstruct := fds.GetDatabaseStructure()

	handlersTemplate := pa.CreateTemplate(tdir+"handlers.txt", "code")
	modelsTemplate := pa.CreateTemplate(tdir+"models.txt", "data")
	datastoreTemplate := pa.CreateTemplate(tdir+"datastore.txt", "dl")
	controllersTemplate := pa.CreateTemplate(tdir+"controllers.txt", "control")

	MainTemplate := pa.CreateTemplate(tdir+"apps/testapp.txt", "main")
	RESTServerTemplate := pa.CreateTemplate(tdir+"apps/restserver.txt", "control")

	rpfolder := strings.Replace(filepath.Base(dbstruct.Database.Filename), filepath.Ext(dbstruct.Database.Filename), "", -1)
	reponame := strings.ToLower(rpfolder)
	fullreponame := "eshu0/" + reponame
	targetrepohost := "github.com"

	// Execute the template for each recipient.
	ctemplates := generator.GenerateHandlers(dbstruct, targetrepohost, fullreponame)

	for _, cs := range ctemplates {
		pa.CreateAndExecute(handlerdir+cs.Filename+".go", handlersTemplate, cs)
		pa.CreateAndExecute(controllersdir+cs.Filename+".go", controllersTemplate, cs)
	}

	ctemplates = generator.GenerateModels(dbstruct, targetrepohost, fullreponame)

	for _, cs := range ctemplates {
		pa.CreateAndExecute(modelsdir+cs.Filename+".go", modelsTemplate, cs)
	}

	dl := Datats{Database: ctemplates[0].Database, Templates: ctemplates, TargetRepoHost: targetrepohost, RepoName: fullreponame}
	pa.CreateAndExecute(strings.ToLower(datastoredir+dl.Database.FilenameTrimmed)+".go", datastoreTemplate, dl)

	// Examples:
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
