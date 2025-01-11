# Open Source Project Security Baseline

* Contents
{:toc}

## Overview

The Open Source Project Security (OSPS) Baseline is a set of security criteria that projects should meet to be considered secure.
The criteria are organized by maturity level and category.
In the detailed subsections you will find the criterion, rationale, and details notes.

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Criteria Overview

* [Level 1](#level-1): for any code or non-code project with any number of maintainers or users
* [Level 2](#level-2): for any code project that has at least 2 maintainers and a small number of consistent users
* [Level 3](#level-3): for any code project that has a large number of consistent users


### Level 1
{{ range .Categories }}
{{- range .Criteria }}
{{- if eq .MaturityLevel 1}}
**[{{ .ID }}]({{ .ID | asLink }})**: {{ .CriterionText | addLinks }}
{{- end }}
{{- end }}
{{- end }}

### Level 2
{{ range .Categories }}
{{- range .Criteria }}
{{- if eq .MaturityLevel 2}}
**[{{ .ID }}]({{ .ID | asLink }})**: {{ .CriterionText | addLinks }}
{{- end }}
{{- end }}
{{- end }}

### Level 3
{{ range .Categories }}
{{- range .Criteria }}
{{- if eq .MaturityLevel 3}}
**[{{ .ID }}]({{ .ID | asLink }})**: {{ .CriterionText | addLinks }}
{{- end }}
{{- end }}
{{- end }}

{{ range .Categories }}

## {{ .CategoryName }}

{{ .Description }}


{{- range .Criteria }}

### {{ .ID }}

**Criterion:** {{ .CriterionText | addLinks }}

**Maturity Level:** {{ .MaturityLevel }}

**Rationale:** {{ .Rationale | addLinks}}

**Details:** {{ .Details | addLinks }}
{{ if .ControlMappings }}
**Control Mappings:**
{{ range .ControlMappings }}
- {{ . }}
{{- end }}
{{- end }}
{{ if .SecurityInsightsValue }}
**Security Insights Value:** {{ .SecurityInsightsValue }}
{{- end }}

---

{{- end }}
{{- end }}


## Lexicon
{{ range .Lexicon }}
### {{ .Term }}

{{ .Definition }}

{{- end }}
---

## Acknowledgments

This baseline was created by community leaders from across the Linux Foundation, including:

- OpenJS Foundation (OpenJS)
- Open Source Security Foundation (OpenSSF)
- Cloud Native Computing Foundation (CNCF)
- Fintech Open Source Foundation (FINOS)
- [OSPS Baseline contributors](https://github.com/ossf/security-baseline/graphs/contributors)
{{ range .Lexicon }}
[{{ .Term }}]: {{ .Term | asLink }}
{{- if .Synonyms }}
{{- $term := .Term }}
{{- range .Synonyms }}
[{{.}}]: {{ $term | asLink }}
{{- end }}
{{- end }}
{{- end }}
