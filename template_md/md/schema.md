{{ define "schema" }}
### {{ .Name }}
| name | type | required | description |
| :--- | :--- | :--- | :--- |
{{ range .Schemas -}}
| {{ .Name }} | {{ .Type }} | {{ .Required }} | {{ .Description }} |
{{ end }}
{{ end }}
