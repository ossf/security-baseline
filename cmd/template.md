# Open Source Project Security Baseline

* Contents
{:toc}

## Overview

The Open Source Project Security (OSPS) Baseline is a set of security criteria that projects should meet to be considered secure.
The criteria are organized by maturity level and category.
In the detailed subsections you will find the criteria, objectives, and implementation notes.

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Criteria Overview

| ID  | Maturity Level | Category | Criteria |
| --- | -------------- | -------- | -------- |

{{- range .Criteria }}
| [{{ .ID }}]({{ .ID | asLink }}) | {{ .MaturityLevel }} | {{ .Category }} | {{ .CriteriaText | collapseNewlines | addLinks }} |
{{- end }}

## Criteria Details

{{- range .Criteria }}

### {{ .ID }}

**Criteria:**

{{ .CriteriaText | addLinks }}
**Objective:**

{{ .Objective | addLinks}}
**Implementation:**

{{ .Implementation | addLinks }}
**Control Mappings:**
{{ if .ControlMappings }}
{{ .ControlMappings }}
{{ else }}
_No control mappings identified._
{{- end }}

**Security Insights Value:**
{{ if .SecurityInsightsValue }}
{{ .SecurityInsightsValue }}
{{ else }}
_No security insights identified._
{{- end }}

**Scorecard Probe(s):**
{{ if .ScorecardProbe }}
{{- range .ScorecardProbe }}
- {{ . }}
{{- end }}
{{- else }}
_No scorecard probe identified._
{{- end }}

---

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
