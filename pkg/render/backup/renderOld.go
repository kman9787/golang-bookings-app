package renderOld

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const templateDir = "./templates"
const templateLayout = "base.layout.tmpl"

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles(templateDir+"/"+tmpl, templateDir+"/"+templateLayout)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error parsing template: ", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")
		err = createTemplate(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf(templateDir+"/%s", t),
		templateDir + "/" + templateLayout,
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
