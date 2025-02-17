// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	oscal "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/ossf/security-baseline/pkg/types"
)

const (
	VersionOSPS = "devel"
	controlHREF = "https://baseline.openssf.org/versions/%s#%s"
	catalogUUID = "8c222a23-fc7e-4ad8-b6dd-289014f07a9f"
)

func (g *Generator) ExportOSCAL(b *types.Baseline, w io.Writer) error {
	n := time.Now()
	catalog := oscal.Catalog{
		UUID:   catalogUUID,
		Groups: nil,
		Metadata: oscal.Metadata{
			LastModified: n,
			Links: &[]oscal.Link{
				{
					Href: fmt.Sprintf(controlHREF, VersionOSPS, ""),
				},
			},
			OscalVersion: "1.1.3",
			Published:    &n,
			Title:        "Open Source Project Security Baseline",
			Version:      VersionOSPS,
		},
	}

	catalogGroups := []oscal.Group{}

	for code, cat := range b.Categories {
		group := oscal.Group{
			Class:    "OSPS",
			Controls: nil,
			ID:       code,
			Title:    cat.Description,
		}

		controls := []oscal.Control{}

		for _, criteria := range cat.Criteria {
			newCtl := oscal.Control{
				Class: criteria.Category,
				ID:    strings.TrimPrefix(criteria.ID, "OSPS-"),
				Links: &[]oscal.Link{
					{
						Href: fmt.Sprintf(controlHREF, VersionOSPS, strings.ToLower(criteria.ID)),
						Rel:  "reference",
					},
				},
				Parts: &[]oscal.Part{
					{
						ID:    strings.TrimPrefix(criteria.ID, "OSPS-") + "_level",
						Name:  "maturity-level",
						Prose: fmt.Sprintf("%d", criteria.MaturityLevel),
					},
					{
						ID:    strings.TrimPrefix(criteria.ID, "OSPS-") + "_details",
						Name:  "details",
						Prose: strings.TrimSpace(criteria.Details),
					},
					{
						ID:    strings.TrimPrefix(criteria.ID, "OSPS-") + "_rationale",
						Name:  "rationale",
						Prose: strings.TrimSpace(criteria.Rationale),
					},
				},
				Title: strings.TrimSpace(criteria.CriterionText),
			}

			if criteria.Implementation != "" {
				pts := append(*newCtl.Parts, oscal.Part{
					ID:    strings.TrimPrefix(criteria.ID, "OSPS-") + "_implementation",
					Name:  "implementation",
					Prose: strings.TrimSpace(criteria.Implementation),
				})
				newCtl.Parts = &pts
			}

			controls = append(controls, newCtl)
		}

		group.Controls = &controls
		catalogGroups = append(catalogGroups, group)
	}
	catalog.Groups = &catalogGroups

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(catalog); err != nil {
		return fmt.Errorf("encoding oscal json data: %w", err)
	}
	return nil
}
