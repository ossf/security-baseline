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

// buildReverseCrosswalk iterates through the Gemara catalog and
// inverts the Baseline -> External mapping into External -> Baseline.
func buildReverseCrosswalk(b *types.Baseline) reverseCrosswalkMap {
	result := make(reverseCrosswalkMap)
	for i := range b.Catalog.Controls {
		control := &b.Catalog.Controls[i]
		controlID := control.Id // FIXED: 'Id' instead of 'ID'
		// A guideline is a MultiEntryMapping
		for _, guideline := range control.Guidelines {
			// The Framework name is usually stored in ReferenceId
			framework := guideline.ReferenceId
			// We have to loop through the specific requirements under this framework
			for _, entry := range guideline.Entries {
				reqID := entry.ReferenceId // e.g. "CC6.1", "AC-2"
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
