package main

import (
	"os"
	"flag"
	"fmt"
    "io/ioutil"
    "text/template"
	sl "github.com/eshu0/simplelogger"
	anl "github.com/eshu0/pangu/pkg/analysers"
)

func main() {

	dbname := flag.String("db", "./some.db", "Database defaults to ./some.db")

	flag.Parse()

	slog := sl.NewApplicationLogger()

	// lets open a flie log using the session
	slog.OpenAllChannels()

	fds := &anl.DatabaseAnalyser{}
	fds.Filename = *dbname
	fds.Create(slog)

	dbstruct := fds.GetDatabaseStructure()

	t := CreateTemplate("./Templates/CodeTemplate.txt","code")
	tz := CreateTemplate("./Templates/DataTemplate.txt", "data")
	t1 := CreateTemplate("./Templates/DLTemplate.txt", "dl")
	t2 := CreateTemplate("./Templates/MainTemplate.txt", "main")

	// Execute the template for each recipient.
	ctemplates := GenerateFile(t,dbstruct,slog)

	for _, cs := range ctemplates {	

		file, err := os.Create("./output/"+cs.Table.Name + ".go")
		if err != nil {
			slog.LogError("CreateCSV", fmt.Sprintf("Cannot create file%s", err.Error()))
			return
		}

		defer file.Close()

		err = t.Execute(file, cs)
		if err != nil {
			fmt.Println("executing template:", err)
		}

		file, err = os.Create("./output/"+cs.Table.Name + "Data.go")
		if err != nil {
			slog.LogError("CreateCSV", fmt.Sprintf("Cannot create file%s", err.Error()))
			return
		}

		defer file.Close()

		err = tz.Execute(file, cs)
		if err != nil {
			fmt.Println("executing template:", err)
		}
	}	

	dl := Datats{Database: ctemplates[0].Database, Templates:ctemplates}

	file1, err := os.Create("./output/"+dl.Database.FilenameTrimmed+"Datastore.go")
	if err != nil {
		slog.LogError("CreateCSV", fmt.Sprintf("Cannot create file%s", err.Error()))
		return
	}
	err = t1.Execute(file1, dl)
	if err != nil {
		fmt.Println("executing template:", err)
	}

	defer file1.Close()


	file2, err := os.Create("./output/main.go")
	if err != nil {
		slog.LogError("CreateCSV", fmt.Sprintf("Cannot create file%s", err.Error()))
		return
	}

	err = t2.Execute(file2, ctemplates)
	if err != nil {
		fmt.Println("executing template:", err)
	}

	defer file2.Close()

}

func CreateTemplate(filepath string, name string) *template.Template {
	b1, err1 := ioutil.ReadFile(filepath) // just pass the file name
    if err1 != nil {
		fmt.Print(err1)
		return nil
    }
    str1 := string(b1) // convert content to a 'string'

	// Create a new template and parse the letter into it.
	return template.Must(template.New(name).Parse(str1))
}
