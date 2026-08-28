// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/ossf/security-baseline/pkg/types"
)

func NewGenerator() *Generator {
	return &Generator{}
}

type Generator struct {
	// ArtifactVersion pins grc.store mapping document links to a published
	// version (the release tag). Empty links to the latest published version.
	ArtifactVersion string
}

// ExportMarkdown runs the baseline data through the markdown template
func (g *Generator) ExportMarkdown(b *types.Baseline, templatePath, path string) error {
	// Read the markdown template from the external file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading markdown template: %w", err)
	}

	// Open or create the output file
	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0o750)); err != nil {
		return fmt.Errorf("error creating output directory %s: %w", filepath.Dir(path), err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", path, err)
	}
	defer outputFile.Close() //nolint:errcheck

	// Create and parse the template
	tmpl, err := template.New("baseline").Funcs(template.FuncMap{
		// Template function to remove newlines and collapse text
		"collapseNewlines": collapseNewlines,
		"addLinks": func(s string) string {
			return addLinksTemplateFunction(b.Lexicon, s)
		},
		"asLink":           asLinkTemplateFunction,
		"maxLevel":         maxLevel,
		"controlsForGroup": controlsForGroup,
		"isRetired":        isRetired,
		"applicabilityTitle": func(id string) string {
			return applicabilityTitle(b.Catalog.Metadata.ApplicabilityGroups, id)
		},
		"relationsForControl": func(controlID string) []FrameworkRelation {
			return relationsForControl(b.Mappings, controlID)
		},
		// grc.store page of the mapping document targeting a framework
		// reference id, or "" when no mapping document covers it.
		"mappingDocURL": func(refID string) string {
			for i := range b.Mappings {
				if b.Mappings[i].TargetReference.ReferenceId == refID {
					return g.artifactURL(b.Mappings[i].Metadata.Id)
				}
			}
			return ""
		},
	}).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template and write to the output file
	if err := tmpl.Execute(outputFile, b); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
