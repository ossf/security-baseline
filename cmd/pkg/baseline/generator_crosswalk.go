// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"

	"github.com/gemaraproj/go-gemara"

	"github.com/ossf/security-baseline/pkg/types"
)

// grcStoreNamespace is the browsable page for the artifacts that
// .github/workflows/publish.yaml pushes to oci.grc.store/openssf.
const grcStoreNamespace = "https://grc.store/openssf"

// artifactURL returns the grc.store page for a published artifact id, pinned
// to the generator's artifact version when one is set.
func (g *Generator) artifactURL(id string) string {
	if g.ArtifactVersion != "" {
		return fmt.Sprintf("%s/%s/versions/%s", grcStoreNamespace, id, g.ArtifactVersion)
	}
	return grcStoreNamespace + "/" + id
}

// reverseMappings inverts one mapping document into
// framework requirement id -> sorted, deduplicated OSPS control ids.
func reverseMappings(doc *gemara.MappingDocument) map[string][]string {
	result := map[string][]string{}
	for _, m := range doc.Mappings {
		for _, t := range m.Targets {
			if t.EntryId == "" {
				continue
			}
			result[t.EntryId] = append(result[t.EntryId], m.Source)
		}
	}
	for req := range result {
		sort.Strings(result[req])
		result[req] = slices.Compact(result[req])
	}
	return result
}

// targetReference resolves a mapping document's target framework to its
// MappingReference entry, falling back to a bare id when it is not declared.
func targetReference(doc *gemara.MappingDocument) gemara.MappingReference {
	refID := doc.TargetReference.ReferenceId
	for _, ref := range doc.Metadata.MappingReferences {
		if ref.Id == refID {
			return ref
		}
	}
	return gemara.MappingReference{Id: refID, Title: refID}
}

// ExportReverseCrosswalk writes a markdown page inverting every loaded
// #MappingDocument: one section per external framework, listing each
// framework requirement alongside the OSPS Baseline controls related to it.
func (g *Generator) ExportReverseCrosswalk(b *types.Baseline, w io.Writer) error {
	var out strings.Builder

	out.WriteString(`# External Framework Crosswalk

This crosswalk inverts the OSPS Baseline mapping documents: for each external
framework requirement, it lists the Baseline controls that the maintainers
believe relate to it. These are reference relationships, not functional
connections, and progress on one does not imply progress on the other.

Every section below is generated from a machine-readable
[Gemara](https://gemara.openssf.org) mapping document published to
[grc.store](` + grcStoreNamespace + `) with each Baseline release.

* Contents
{:toc}
`)

	for i := range b.Mappings {
		doc := &b.Mappings[i]
		if doc.TargetReference.ReferenceId == "" {
			continue
		}
		ref := targetReference(doc)

		fmt.Fprintf(&out, "\n## %s\n\n", ref.Title)
		if ref.Version != "" {
			fmt.Fprintf(&out, "Version %s · ", ref.Version)
		}
		if ref.Url != "" {
			fmt.Fprintf(&out, "[Framework](%s) · ", ref.Url)
		}
		fmt.Fprintf(&out, "[Mapping document](%s)\n\n", g.artifactURL(doc.Metadata.Id))

		out.WriteString("| Requirement | OSPS Baseline Controls |\n")
		out.WriteString("|-------------|------------------------|\n")

		crosswalk := reverseMappings(doc)
		reqIDs := make([]string, 0, len(crosswalk))
		for req := range crosswalk {
			reqIDs = append(reqIDs, req)
		}
		sort.Strings(reqIDs)
		for _, req := range reqIDs {
			fmt.Fprintf(&out, "| %s | %s |\n", req, strings.Join(crosswalk[req], ", "))
		}
	}

	_, err := io.WriteString(w, out.String())
	return err
}
