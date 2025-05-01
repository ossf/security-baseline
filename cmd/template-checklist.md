# OSPS Baseline checklist

## Level 1
{{ range .Categories }}
{{- range .Controls }}
{{- range .Requirements }}
{{- $req := . }}
{{- if maxLevel .Applicability 1 }}
- [ ] **{{ $req.ID }}**: {{ $req.Text | collapseNewlines }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

## Level 2
{{ range .Categories }}
{{- range .Controls }}
{{- range .Requirements }}
{{- $req := . }}
{{- if maxLevel .Applicability 2 }}
- [ ] **{{ $req.ID }}**: {{ $req.Text | collapseNewlines }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

## Level 3
{{ range .Categories }}
{{- range .Controls }}
{{- range .Requirements }}
{{- $req := . }}
{{- if maxLevel .Applicability 3 }}
- [ ] **{{ $req.ID }}**: {{ $req.Text | collapseNewlines }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}