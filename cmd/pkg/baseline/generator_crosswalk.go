// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/ossf/security-baseline/pkg/types"
)

// reverseCrosswalkMap is a nested map:
// Framework name -> Requirement ID -> sorted list of OSPS Control IDs
type reverseCrosswalkMap map[string]map[string][]string

// buildReverseCrosswalk iterates the loaded #MappingDocument artifacts and
// inverts the Baseline -> External mapping into External -> Baseline.
// The framework name is taken from each mapping document's TargetReference.
func buildReverseCrosswalk(b *types.Baseline) reverseCrosswalkMap {
	result := make(reverseCrosswalkMap)
	for i := range b.Mappings {
		doc := &b.Mappings[i]
		framework := doc.TargetReference.ReferenceId
		if framework == "" {
			continue
		}
		for _, m := range doc.Mappings {
			controlID := m.Source
			for _, t := range m.Targets {
				reqID := t.EntryId
				if reqID == "" {
					continue
				}
				if _, ok := result[framework]; !ok {
					result[framework] = make(map[string][]string)
				}
				result[framework][reqID] = append(result[framework][reqID], controlID)
			}
		}
	}
	// Sort the control ID lists within each cell for determinism
	for framework := range result {
		for reqID := range result[framework] {
			sort.Strings(result[framework][reqID])
		}
	}
	return result
}

// ExportReverseCrosswalk writes a GitHub-flavored Markdown table
// mapping External Framework Requirements -> OSPS Baseline Controls
// to the provided writer.
func (g *Generator) ExportReverseCrosswalk(b *types.Baseline, w io.Writer) error {
	crosswalk := buildReverseCrosswalk(b)

	// Sort framework names for deterministic output
	frameworks := make([]string, 0, len(crosswalk))
	for fw := range crosswalk {
		frameworks = append(frameworks, fw)
	}
	sort.Strings(frameworks)
	// Write table header
	_, err := fmt.Fprintln(w, "| Framework | Requirement | OSPS Controls |")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, "|-----------|-------------|---------------|")
	if err != nil {
		return err
	}
	// Write rows, sorted by framework then requirement ID
	for _, fw := range frameworks {
		reqMap := crosswalk[fw]
		reqIDs := make([]string, 0, len(reqMap))
		for req := range reqMap {
			reqIDs = append(reqIDs, req)
		}
		sort.Strings(reqIDs)
		for _, req := range reqIDs {
			controls := strings.Join(reqMap[req], ", ")
			_, err := fmt.Fprintf(w, "| %s | %s | %s |\n", fw, req, controls)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
