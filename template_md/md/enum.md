{{ define "enum" }}
### {{ .Name }}
| value | description |
| :--- | :--- |
{{ range .Enums -}}
| {{ .Value }} | {{ .Description }} |
{{ end }}
{{ end }}
