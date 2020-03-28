package template_md

import (
	"io"
	"text/template"
)

type TaskTemplate struct {
	HeaderTemplate     *HeaderTemplate
	SchemaTemplateList *[]SchemaTemplate
	EnumTemplateList   *[]EnumTemplate
}

func NewTaskTemplate() *TaskTemplate {
	return &TaskTemplate{
		HeaderTemplate:     &HeaderTemplate{},
		SchemaTemplateList: &[]SchemaTemplate{},
		EnumTemplateList:   &[]EnumTemplate{},
	}
}

func (t *TaskTemplate) SetHeader(h *HeaderTemplate) {
	t.HeaderTemplate = h
}

func (t *TaskTemplate) SetSchema(s *SchemaTemplate) {
	*t.SchemaTemplateList = append(*t.SchemaTemplateList, *s)
}

func (t *TaskTemplate) SetEnum(e *EnumTemplate) {
	*t.EnumTemplateList = append(*t.EnumTemplateList, *e)
}

func (t *TaskTemplate) Exec(Template *template.Template, wr io.Writer) error {
	return Template.Execute(wr, t)
}

type HeaderTemplate struct {
	Description string
}

func NewHeaderTemplate(description string) *HeaderTemplate {
	return &HeaderTemplate{description}
}

type SchemaTemplate struct {
	Name    string
	Schemas *[]Schema
}

type Schema struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

func NewSchemaTemplate(name string) *SchemaTemplate {
	schema := make([]Schema, 0)
	return &SchemaTemplate{
		Name:    name,
		Schemas: &schema,
	}
}

func (s *SchemaTemplate) Set(name, typ, description string, required bool) {
	*s.Schemas = append(*s.Schemas, Schema{
		Name:        name,
		Type:        typ,
		Required:    required,
		Description: description,
	})
}

type EnumTemplate struct {
	Name  string
	Enums *[]Enum
}

type Enum struct {
	Value       string
	Description string
}

func NewEnumTemplate(name string) *EnumTemplate {
	enum := make([]Enum, 0)
	return &EnumTemplate{
		Name:  name,
		Enums: &enum,
	}
}

func (e *EnumTemplate) Set(value, description string) {
	*e.Enums = append(*e.Enums, Enum{
		Value:       value,
		Description: description,
	})
}
