{{ template "header" .HeaderTemplate }}
{{ range .SchemaTemplateList -}}
{{ template "schema" . }}
{{ end }}
{{ range .EnumTemplateList -}}
{{ template "enum" . }}
{{ end }}
