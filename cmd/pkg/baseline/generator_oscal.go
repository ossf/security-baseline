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

	// OpenSSFNS is the OSCAL namespace URI to define the baseline names.
	OpenSSFNS = "http://baseline.openssf.org/ns/oscal"
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
					Rel:  "canonical",
				},
			},
			OscalVersion: "1.1.3",
			Published:    &n,
			Title:        "Open Source Project Security Baseline",
			Version:      VersionOSPS,
		},
	}

	catalogGroups := []oscal.Group{}

	for _, family := range b.Catalog.ControlFamilies {
		group := oscal.Group{
			Class:    "OSPS",
			Controls: nil,
			ID:       b.ControlFamilyIDs[family.Title],
			Parts: &[]oscal.Part{
				{
					Name:  "overview",
					Prose: family.Description,
				},
			},
			Title: family.Title,
		}

		controls := []oscal.Control{}
		for _, control := range family.Controls {
			// Create the new OSCAL control.
			newOscalCtl := oscal.Control{
				Class: b.ControlFamilyIDs[family.Title],
				ID:    control.Id,
				Title: strings.TrimSpace(control.Id), // For some reason, control.Title is the full description
				Links: &[]oscal.Link{
					{
						Href: fmt.Sprintf(controlHREF, VersionOSPS, strings.ToLower(control.Id)),
						Rel:  "canonical",
					},
				},
				// The main prose of the control lives in the statement part
				Parts: &[]oscal.Part{
					{
						ID:    control.Id + "_smt",
						Name:  "statement",
						Ns:    OpenSSFNS,
						Prose: control.Title,
					},
					{
						ID:    control.Id + "_obj",
						Name:  "objective",
						Ns:    OpenSSFNS,
						Prose: control.Objective,
					},
				},
			}

			items := []oscal.Part{}
			for _, ar := range control.AssessmentRequirements {
				items = append(items, oscal.Part{
					ID:    ar.Id,
					Name:  "item",
					Ns:    OpenSSFNS,
					Prose: ar.Text,
					Title: ar.Id,
					Parts: &[]oscal.Part{
						{
							ID:    ar.Id + "_obj",
							Name:  "assessment-objective",
							Prose: ar.Recommendation,
						},
					},
				})
			}

			(*newOscalCtl.Parts)[0].Parts = &items
			controls = append(controls, newOscalCtl)
		}

		group.Controls = &controls
		catalogGroups = append(catalogGroups, group)
	}
	catalog.Groups = &catalogGroups

	// Wrap the catalog to render the required "catalog" wrapper
	// in the JSON file:
	wrapper := struct {
		Catalog oscal.Catalog `json:"catalog"`
	}{
		Catalog: catalog,
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(wrapper); err != nil {
		return fmt.Errorf("encoding oscal json data: %w", err)
	}
	return nil
}
