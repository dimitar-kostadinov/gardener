# managed by gardener-node-agent
{{- if .server }}
server = {{ .server | quote }}
{{ end }}
# [host.{{ .localIPURL | quote }}]
[host."http://localhost:5500"]
  capabilities = ["pull","resolve"]
{{- range .hostConfigs }}
[host.{{ .hostURL | quote }}]
  capabilities = {{ .capabilities | toJson }}
  {{- if .ca }}
  ca = {{ .ca | toJson }}
  {{- end }}
{{ end }}
