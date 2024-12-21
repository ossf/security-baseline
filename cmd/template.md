# Open Source Project Security Baseline

* Contents
{:toc}

## Overview

The Open Source Project Security (OSPS) Baseline is a set of security criterion that projects should meet to be considered secure.
The criterion are organized by maturity level and category.
In the detailed subsections you will find the criterion, rationales, and details notes.

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Criterion Overview

| ID  | Maturity Level | Category | Criterion |
| --- | -------------- | -------- | -------- |

{{- range .Criterion }}
| [{{ .ID }}]({{ .ID | asLink }}) | {{ .MaturityLevel }} | {{ .Category }} | {{ .CriterionText | collapseNewlines | addLinks }} |
{{- end }}

## Criterion Details

{{- range .Criterion }}

### {{ .ID }}

**Criterion:**

{{ .CriterionText | addLinks }}
**Rationale:**

{{ .Rationale | addLinks}}
**Details:**

{{ .Details | addLinks }}
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
