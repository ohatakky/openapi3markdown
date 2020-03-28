package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ohatakky/openapi3markdown/template_md"
)

const (
	dir    = "template_md/md/"
	outDir = "Docs/"
)

var (
	swagger  *openapi3.Swagger
	Template *template.Template
	files    []string = []string{
		dir + "task.md",
		dir + "header.md",
		dir + "schema.md",
		dir + "enum.md",
	}
)

func init() {
	s, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	swagger, err = openapi3.NewSwaggerLoader().LoadSwaggerFromData(s)
	if err != nil {
		panic(err)
	}
	Template, err = template.New("task.md").ParseFiles(files...)
	if err != nil {
		panic(err)
	}
}

func main() {
	schemas := swagger.Components.Schemas
	tMap := make(map[string]*template_md.TaskTemplate, len(swagger.Components.Schemas))
	for task, schema := range schemas {
		h := template_md.NewHeaderTemplate(schema.Value.Description)
		tMap[task] = template_md.NewTaskTemplate()
		tMap[task].SetHeader(h)
		s := template_md.NewSchemaTemplate(task)
		recrusive(schema.Value.Properties, schema.Value.Required, s, tMap[task])
	}

	err := write(tMap)
	if err != nil {
		log.Fatal(err)
	}
}

func recrusive(properties map[string]*openapi3.SchemaRef, required []string, s *template_md.SchemaTemplate, t *template_md.TaskTemplate) {
	for key, property := range properties {
		if property.Value.Items != nil {
			ps := property.Value.Items.Value.Properties
			if len(ps) == 0 {
				s.Set(key, fmt.Sprintf("[%s]", property.Value.Items.Value.Type), property.Value.Description, contains(required, key))
			} else {
				s.Set(key, fmt.Sprintf("[[%s]](#%s)", key, key), property.Value.Description, contains(required, key))
				ss := template_md.NewSchemaTemplate(key)
				recrusive(ps, property.Value.Items.Value.Required, ss, t)
			}
		} else {
			if property.Value.Enum != nil {
				e := template_md.NewEnumTemplate(key)
				for _, v := range property.Value.Enum {
					e.Set(fmt.Sprintf("%s", v), "")
				}
				t.SetEnum(e)
				s.Set(fmt.Sprintf("[%s](#%s)", key, key), fmt.Sprintf("%s", property.Value.Type), property.Value.Description, contains(required, key))
			} else {
				s.Set(key, fmt.Sprintf("%s", property.Value.Type), property.Value.Description, contains(required, key))
			}
		}
	}
	t.SetSchema(s)
}

func write(tMap map[string]*template_md.TaskTemplate) error {
	if err := os.Mkdir(outDir, 0777); err != nil {
		return err
	}
	for k, v := range tMap {
		file, err := os.Create(fmt.Sprintf("%s/%s.md", outDir, k))
		if err != nil {
			return err
		}
		defer file.Close()
		err = v.Exec(Template, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
