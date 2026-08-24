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

func TestBuildReverseCrosswalk(t *testing.T) {
	b := setupFakeBaseline()
	result := buildReverseCrosswalk(b)

	if len(result) != 2 {
		t.Errorf("expected 2 frameworks, got %d", len(result))
	}

	nist, ok := result["NIST SP 800-53"]
	if !ok {
		t.Fatal("expected NIST SP 800-53 framework in result")
	}

	ac2Controls, ok := nist["AC-2"]
	if !ok {
		t.Fatal("expected AC-2 requirement under NIST SP 800-53")
	}
	if len(ac2Controls) != 2 {
		t.Errorf("expected 2 controls for AC-2, got %d", len(ac2Controls))
	}
	if ac2Controls[0] != "OSPS-AC-01" || ac2Controls[1] != "OSPS-GV-01" {
		t.Errorf("unexpected controls for AC-2: %v", ac2Controls)
	}

	cis, ok := result["CIS Controls v8"]
	if !ok {
		t.Fatal("expected CIS Controls v8 framework in result")
	}
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

func TestGenerateReverseCrosswalk_TableFormat(t *testing.T) {
	b := setupFakeBaseline()
	g := NewGenerator()

	var out strings.Builder
	if err := g.ExportReverseCrosswalk(b, &out); err != nil {
		t.Fatal(err)
	}

	result := out.String()

	if !strings.Contains(result, "| Framework | Requirement | OSPS Controls |") {
		t.Error("missing table header")
	}
	if !strings.Contains(result, "NIST SP 800-53") {
		t.Error("missing NIST framework in output")
	}
	if !strings.Contains(result, "CIS Controls v8") {
		t.Error("missing CIS framework in output")
	}
	if !strings.Contains(result, "OSPS-AC-01, OSPS-GV-01") {
		t.Error("expected sorted control IDs joined with comma")
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
