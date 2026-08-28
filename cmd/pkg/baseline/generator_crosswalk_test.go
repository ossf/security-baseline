// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"strings"
	"testing"

	"github.com/gemaraproj/go-gemara"

	"github.com/ossf/security-baseline/pkg/types"
)

func setupFakeBaseline() *types.Baseline {
	return &types.Baseline{
		Catalog: gemara.ControlCatalog{
			Controls: []gemara.Control{
				{Id: "OSPS-AC-01"},
				{Id: "OSPS-GV-01"},
			},
		},
		Mappings: []gemara.MappingDocument{
			{
				Metadata: gemara.Metadata{
					Id: "osps-baseline-to-nist",
					MappingReferences: []gemara.MappingReference{
						{Id: "osps-baseline", Title: "OSPS Baseline", Version: "devel"},
						{Id: "NIST SP 800-53", Title: "NIST Special Publication 800-53", Version: "r5", Url: "https://example.com/800-53"},
					},
				},
				SourceReference: gemara.TypedMapping{ReferenceId: "osps-baseline", EntryType: gemara.EntryTypeControl},
				TargetReference: gemara.TypedMapping{ReferenceId: "NIST SP 800-53", EntryType: gemara.EntryTypeGuideline},
				Mappings: []gemara.Mapping{
					{
						Id:           "M-OSPS-AC-01-NIST",
						Source:       "OSPS-AC-01",
						Relationship: gemara.RelRelatesTo,
						Targets: []gemara.MappingTarget{
							{EntryId: "AC-2"},
							{EntryId: "AC-3"},
						},
					},
					{
						Id:           "M-OSPS-GV-01-NIST",
						Source:       "OSPS-GV-01",
						Relationship: gemara.RelRelatesTo,
						Targets: []gemara.MappingTarget{
							{EntryId: "AC-2"},
						},
					},
				},
			},
			{
				Metadata:        gemara.Metadata{Id: "osps-baseline-to-cis"},
				SourceReference: gemara.TypedMapping{ReferenceId: "osps-baseline", EntryType: gemara.EntryTypeControl},
				TargetReference: gemara.TypedMapping{ReferenceId: "CIS Controls v8", EntryType: gemara.EntryTypeGuideline},
				Mappings: []gemara.Mapping{
					{
						Id:           "M-OSPS-GV-01-CIS",
						Source:       "OSPS-GV-01",
						Relationship: gemara.RelRelatesTo,
						Targets: []gemara.MappingTarget{
							{EntryId: "1.1"},
						},
					},
				},
			},
		},
	}
}

func TestReverseMappings(t *testing.T) {
	b := setupFakeBaseline()

	nist := reverseMappings(&b.Mappings[0])
	ac2, ok := nist["AC-2"]
	if !ok {
		t.Fatal("expected AC-2 requirement under NIST SP 800-53")
	}
	if len(ac2) != 2 || ac2[0] != "OSPS-AC-01" || ac2[1] != "OSPS-GV-01" {
		t.Errorf("unexpected controls for AC-2: %v", ac2)
	}

	// Duplicate source->target pairs must collapse to one entry.
	b.Mappings[0].Mappings = append(b.Mappings[0].Mappings, gemara.Mapping{
		Id:           "M-OSPS-AC-01-NIST-dup",
		Source:       "OSPS-AC-01",
		Relationship: gemara.RelRelatesTo,
		Targets:      []gemara.MappingTarget{{EntryId: "AC-2"}, {EntryId: ""}},
	})
	nist = reverseMappings(&b.Mappings[0])
	if got := nist["AC-2"]; len(got) != 2 {
		t.Errorf("expected duplicates deduplicated for AC-2, got %v", got)
	}

	cis := reverseMappings(&b.Mappings[1])
	if len(cis["1.1"]) != 1 || cis["1.1"][0] != "OSPS-GV-01" {
		t.Errorf("unexpected controls for CIS 1.1: %v", cis["1.1"])
	}
}

func TestGenerateReverseCrosswalk_Deterministic(t *testing.T) {
	b := setupFakeBaseline()
	g := NewGenerator()

	var b1, b2 strings.Builder

	if err := g.ExportReverseCrosswalk(b, &b1); err != nil {
		t.Fatal(err)
	}
	if err := g.ExportReverseCrosswalk(b, &b2); err != nil {
		t.Fatal(err)
	}
	if b1.String() != b2.String() {
		t.Error("output is not deterministic")
	}
}

func TestGenerateReverseCrosswalk_PageFormat(t *testing.T) {
	b := setupFakeBaseline()
	g := NewGenerator()
	g.ArtifactVersion = "v2026.08.26"

	var out strings.Builder
	if err := g.ExportReverseCrosswalk(b, &out); err != nil {
		t.Fatal(err)
	}

	result := out.String()

	// One section per framework, titled from the doc's own MappingReference
	// when it has one, and the bare reference id otherwise.
	if !strings.Contains(result, "## NIST Special Publication 800-53") {
		t.Error("missing resolved NIST framework section")
	}
	if !strings.Contains(result, "Version r5 · [Framework](https://example.com/800-53)") {
		t.Error("missing framework version/url meta line")
	}
	if !strings.Contains(result, "## CIS Controls v8") {
		t.Error("missing fallback CIS framework section")
	}
	if !strings.Contains(result, "| Requirement | OSPS Baseline Controls |") {
		t.Error("missing table header")
	}
	if !strings.Contains(result, "| AC-2 | OSPS-AC-01, OSPS-GV-01 |") {
		t.Error("expected sorted control IDs joined with comma")
	}
	// Mapping document links pin to the artifact version when one is set.
	if !strings.Contains(result, "https://grc.store/openssf/osps-baseline-to-nist/versions/v2026.08.26") {
		t.Error("missing pinned grc.store mapping document link")
	}

	var unpinned strings.Builder
	if err := NewGenerator().ExportReverseCrosswalk(b, &unpinned); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(unpinned.String(), "(https://grc.store/openssf/osps-baseline-to-nist)") {
		t.Error("missing unpinned grc.store mapping document link")
	}
}

func TestRelationsForControl(t *testing.T) {
	b := setupFakeBaseline()

	// OSPS-GV-01 is mapped by both documents, so it exercises aggregation
	// across frameworks and the alphabetical sort of the result.
	gv := relationsForControl(b.Mappings, "OSPS-GV-01")
	if len(gv) != 2 {
		t.Fatalf("expected 2 frameworks for OSPS-GV-01, got %d", len(gv))
	}
	if gv[0].Framework != "CIS Controls v8" || gv[1].Framework != "NIST SP 800-53" {
		t.Errorf("frameworks not sorted by name: %q, %q", gv[0].Framework, gv[1].Framework)
	}
	if len(gv[1].Entries) != 1 || gv[1].Entries[0] != "AC-2" {
		t.Errorf("unexpected NIST entries for OSPS-GV-01: %v", gv[1].Entries)
	}

	ac := relationsForControl(b.Mappings, "OSPS-AC-01")
	if len(ac) != 1 {
		t.Fatalf("expected 1 framework for OSPS-AC-01, got %d", len(ac))
	}
	if len(ac[0].Entries) != 2 || ac[0].Entries[0] != "AC-2" || ac[0].Entries[1] != "AC-3" {
		t.Errorf("entry order not preserved: %v", ac[0].Entries)
	}

	if got := relationsForControl(b.Mappings, "OSPS-ZZ-99"); got != nil {
		t.Errorf("expected nil for an unmapped control, got %v", got)
	}

	// A target with no entry-id must be skipped rather than rendered blank,
	// which is what produced the empty crosswalk row for OSPS-DO-07.
	b.Mappings[1].Mappings[0].Targets = append(b.Mappings[1].Mappings[0].Targets, gemara.MappingTarget{})
	cis := relationsForControl(b.Mappings, "OSPS-GV-01")[0]
	if len(cis.Entries) != 1 {
		t.Errorf("empty entry-id was not skipped: %v", cis.Entries)
	}
}

func TestArtifactURL(t *testing.T) {
	g := NewGenerator()
	if got := g.artifactURL(""); got != "" {
		t.Errorf("empty artifact id must not link to the bare namespace, got %q", got)
	}
	if got := g.artifactURL("osps-baseline"); got != "https://grc.store/openssf/osps-baseline" {
		t.Errorf("unexpected unpinned URL: %q", got)
	}
	g.ArtifactVersion = "v1"
	if got := g.artifactURL("osps-baseline"); got != "https://grc.store/openssf/osps-baseline/versions/v1" {
		t.Errorf("unexpected pinned URL: %q", got)
	}
}
