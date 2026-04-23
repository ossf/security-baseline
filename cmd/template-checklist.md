# OSPS Baseline checklist, version: devel

## Level 1
{{ range .Catalog.Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if and (not (isRetired .State)) (maxLevel .Applicability 1) }}
- [ ] **{{ $req.Id }}**: {{ $req.Text | collapseNewlines }}
{{- end }}
{{- end }}
{{- end }}

## Level 2
{{ range .Catalog.Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if and (not (isRetired .State)) (maxLevel .Applicability 2) }}
- [ ] **{{ $req.Id }}**: {{ $req.Text | collapseNewlines }}
{{- end }}
{{- end }}
{{- end }}

## Level 3
{{ range .Catalog.Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if and (not (isRetired .State)) (maxLevel .Applicability 3) }}
- [ ] **{{ $req.Id }}**: {{ $req.Text | collapseNewlines }}
{{- end }}
{{- end }}
{{- end }}
