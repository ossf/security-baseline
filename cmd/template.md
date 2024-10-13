# Open Source Project Security Baseline

## Overview

The Basleine is a set of security criteria that projects should meet to be considered secure. The criteria are organized by maturity level and category. In the detailed subsections you will find the criteria, objectives, and implementation notes.

> [!NOTE]
> This document was automatically generated from the [Open Source Project Security Baseline YAML](https://github.com/ossf/security-baselines/blob/main/baselines/ossf-security-baseline.yaml) file.

---

## Baseline Criteria

| ID  | Maturity Level | Category | Criteria |
| --- | -------------- | -------- | -------- |

{{- range .Criteria }}
| {{ .ID }} | {{ .MaturityLevel }} | {{ .Category }} | {{ .CriteriaText | collapseNewlines | addLinks }} |
{{- end }}

---

{{- range .Criteria }}

### {{ .ID }} - {{ .Category }}

**Criteria:**

{{ .CriteriaText | addLinks}}
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

**Scorecard Probe:**  
{{ if .ScorecardProbe }}  
{{ .ScorecardProbe }}  
{{ else }}  
_No scorecard probe identified._  
{{- end }}

---

{{- end }}

## Lexicon

{{- range .Lexicon }}
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
{{ range .Lexicon }}
{{- if .Synonyms }}
[{{ .Term }}]: {{ .Term | linkPrep }}
{{- $term := .Term }}
{{- range .Synonyms }}
[{{.}}]: {{ $term | linkPrep }}
{{- end }}
{{- end }}
{{- end }}
