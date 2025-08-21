# Open Source Project Security Baseline

Version: devel

{: .warning}
Not for production use.

<!-- A button for returning to the top of the page -->
<button onclick="toTop()" id="topButton" title="Go to top"
style="display: none; position: fixed; bottom: 20px; right: 30px; border: none; background-color: CornflowerBlue; color: white; cursor: pointer; padding: 10px; border-radius: 10px; font-size: 18px;">to top</button>

<script>
let topButton = document.getElementById("topButton");
window.onscroll = function() {scrollFunction()};

function scrollFunction() {
  if (document.documentElement.scrollTop > 50 ) {
    topButton.style.display = "block";
  } else {
    topButton.style.display = "none";
  }
}

function toTop() {
  document.documentElement.scrollTop = 0;
}
</script>
<!-- That's enough button stuff for now -->

* Contents
{:toc}

## Overview

The Open Source Project Security (OSPS) Baseline is a set of security criteria that projects should meet to demonstrate a strong security posture.
The controls are organized by maturity level and category.
In the detailed subsections you will find the control, rationale, and details notes.

Where possible, we have added control mappings to external frameworks.
These are not guaranteed to be 100% matches, but instead serve as references
when working to meet the corresponding controls.

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Controls Overview

* [Level 1](#level-1): for any code or non-code project with any number of maintainers or users
* [Level 2](#level-2): for any code project that has at least 2 maintainers and a small number of consistent users
* [Level 3](#level-3): for any code project that has a large number of consistent users

### Level 1
{{ range .Catalog.ControlFamilies }}
{{- range .Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if maxLevel .Applicability 1 }}
**[{{ $req.Id }}]({{ $req.Id | asLink }})**: {{ $req.Text | addLinks }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

### Level 2
{{ range .Catalog.ControlFamilies }}
{{- range .Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if maxLevel .Applicability 2 }}
**[{{ $req.Id }}]({{ $req.Id | asLink }})**: {{ $req.Text | addLinks }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

### Level 3
{{ range .Catalog.ControlFamilies }}
{{- range .Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if maxLevel .Applicability 3 }}
**[{{ $req.Id }}]({{ $req.Id | asLink }})**: {{ $req.Text | addLinks }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}

{{ range .Catalog.ControlFamilies }}

## {{ .Title }}

{{ .Description }}

{{ range .Controls }}

### {{ .Id }} - {{ .Title | addLinks | collapseNewlines }}

{{ .Objective }}

{{ range .AssessmentRequirements }}

#### {{ .Id }}

**Requirement:** {{ .Text | addLinks | collapseNewlines }}

**Recommendation:** {{ .Recommendation }}

**Control applies to:**
{{ range .Applicability }}- {{ . }}
{{ end }}

{{ end }}

#### External Framework Mappings
{{ if  .GuidelineMappings }}
  {{ range .GuidelineMappings }}
  - **{{ .ReferenceId | addLinks }}**: {{ range $index, $entry := .Entries }}{{ if $index }}, {{ end }}{{ $entry.ReferenceId }}{{ end }}
  {{- end }}
{{ end }}

---

{{- end }}
{{- end }}

## External Frameworks

Controls within this document may map to the following external frameworks:

| ID | Title | Version | Description |
|----|-------|---------|-------------|
{{ range .Catalog.Metadata.MappingReferences -}}
| {{ .Id }} | [{{ .Title }}]({{ .Url }}) | {{ .Version }} | {{ .Description }} |
{{ end }}

---

## Lexicon
{{ range .Lexicon }}

### {{ .Term }}

{{ .Definition }}

{{ if .References }}
**References:**
{{ range .References }}
  - {{.}}
{{ end -}}
{{ end -}}
{{ end }}

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
