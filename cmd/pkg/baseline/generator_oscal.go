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

	for code, cat := range b.Categories {
		group := oscal.Group{
			Class:    "OSPS",
			Controls: nil,
			ID:       code,
			Title:    cat.Description,
		}

		controls := []oscal.Control{}

		for _, control := range cat.Controls {
			parts := []oscal.Part{}
			for _, ar := range control.Requirements {
				parts = append(parts, oscal.Part{
					Class: control.ID,
					ID:    ar.ID,
					Name:  ar.ID,
					Ns:    "",
					Parts: &[]oscal.Part{
						{
							ID:    ar.ID + ".R",
							Name:  "recommendation",
							Ns:    OpenSSFNS,
							Prose: ar.Recommendation,
							Links: &[]oscal.Link{
								{
									Href: fmt.Sprintf(controlHREF, VersionOSPS, ar.ID),
									Rel:  "canonical",
								},
							},
						},
					},
					Prose: ar.Text,
					Title: "",
				})
			}

			newCtl := oscal.Control{
				Class: code,
				ID:    control.ID,
				Links: &[]oscal.Link{
					{
						Href: fmt.Sprintf(controlHREF, VersionOSPS, strings.ToLower(control.ID)),
						Rel:  "canonical",
					},
				},
				Parts: &parts,
				Title: strings.TrimSpace(control.Title),
			}
			controls = append(controls, newCtl)
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
