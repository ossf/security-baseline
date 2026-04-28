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
				{
					Id: "OSPS-AC-01",
					Guidelines: []gemara.MultiEntryMapping{
						{
							ReferenceId: "NIST SP 800-53",
							Entries: []gemara.ArtifactMapping{
								{ReferenceId: "AC-2"},
								{ReferenceId: "AC-3"},
							},
						},
					},
				},
				{
					Id: "OSPS-GV-01",
					Guidelines: []gemara.MultiEntryMapping{
						{
							ReferenceId: "NIST SP 800-53",
							Entries: []gemara.ArtifactMapping{
								{ReferenceId: "AC-2"},
							},
						},
						{
							ReferenceId: "CIS Controls v8",
							Entries: []gemara.ArtifactMapping{
								{ReferenceId: "1.1"},
							},
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
