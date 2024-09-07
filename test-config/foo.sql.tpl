SELECT 
	*
FROM 
	{{.TableName}}
WHERE
	{{- if .Filters }}
		{{- range $index, $filter := .Filters }}
			{{- "\n\t" -}}
			{{- if ne $index 0 }}
				{{- "AND " -}}
			{{- end }}
			{{- $filter -}}
		{{- end }}
	{{- end }}