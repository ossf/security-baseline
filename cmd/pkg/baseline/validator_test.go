// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"strings"
	"testing"

	"github.com/gemaraproj/go-gemara"

	"github.com/ossf/security-baseline/pkg/types"
)

const (
	testControlID = "OSPS-AC-01"
	testSourceRef = "osps-baseline"
	testTerm      = "Repository"
	wantCollision = `collides with entry "Repository"`
)

// validBaseline is the smallest baseline that passes Check: one control, one
// declared mapping-reference, and one mapping document targeting it.
func validBaseline() *types.Baseline {
	return &types.Baseline{
		Catalog: gemara.ControlCatalog{
			Metadata: gemara.Metadata{
				MappingReferences: []gemara.MappingReference{
					{Id: "CSF", Title: "Cybersecurity Framework", Version: "2.0"},
				},
			},
			Controls: []gemara.Control{
				{Id: testControlID, Title: "A control"},
			},
		},
		Mappings: []gemara.MappingDocument{
			{
				Metadata:        gemara.Metadata{Id: "osps-to-csf"},
				SourceReference: gemara.TypedMapping{ReferenceId: testSourceRef, EntryType: gemara.EntryTypeControl},
				TargetReference: gemara.TypedMapping{ReferenceId: "CSF", EntryType: gemara.EntryTypeGuideline},
				Mappings: []gemara.Mapping{
					{Id: "M-1", Source: testControlID, Relationship: gemara.RelRelatesTo},
				},
			},
		},
		Lexicon: []types.LexiconEntry{
			{Term: testTerm, Synonyms: []string{"Repo"}},
		},
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*types.Baseline)
		wantErr string // substring; empty means Check must pass
	}{
		{
			name:   "valid baseline passes",
			mutate: func(*types.Baseline) {},
		},
		{
			name: "mapping document targeting an undeclared reference",
			mutate: func(b *types.Baseline) {
				b.Mappings[0].TargetReference.ReferenceId = "NOPE"
			},
			wantErr: `targets reference "NOPE"`,
		},
		{
			// Regression guard: an empty reference-id used to be exempt, and
			// the renderer drops such a document without a word.
			name: "mapping document with an empty reference id",
			mutate: func(b *types.Baseline) {
				b.Mappings[0].TargetReference.ReferenceId = ""
			},
			wantErr: `targets reference ""`,
		},
		{
			name: "declared reference with no mapping document",
			mutate: func(b *types.Baseline) {
				b.Catalog.Metadata.MappingReferences = append(b.Catalog.Metadata.MappingReferences,
					gemara.MappingReference{Id: "ORPHAN", Title: "Orphan", Version: "1"})
			},
			wantErr: `mapping-reference "ORPHAN" is declared in metadata but no mapping document targets it`,
		},
		{
			name: "duplicate lexicon term",
			mutate: func(b *types.Baseline) {
				b.Lexicon = append(b.Lexicon, types.LexiconEntry{Term: testTerm})
			},
			wantErr: wantCollision,
		},
		{
			// asLink case-folds a term into its anchor, so these two share one
			// destination even though the strings differ.
			name: "lexicon terms differing only by case",
			mutate: func(b *types.Baseline) {
				b.Lexicon = append(b.Lexicon, types.LexiconEntry{Term: "repository"})
			},
			wantErr: wantCollision,
		},
		{
			name: "synonym colliding with another entry's term",
			mutate: func(b *types.Baseline) {
				b.Lexicon = append(b.Lexicon, types.LexiconEntry{
					Term:     "Source Control",
					Synonyms: []string{testTerm},
				})
			},
			wantErr: wantCollision,
		},
		{
			// Harmless: addLinks skips text it has already wrapped.
			name: "synonym repeating its own term is allowed",
			mutate: func(b *types.Baseline) {
				b.Lexicon[0].Synonyms = append(b.Lexicon[0].Synonyms, testTerm)
			},
		},
		{
			name: "mapping targeting an unknown control",
			mutate: func(b *types.Baseline) {
				b.Mappings[0].Mappings[0].Source = "OSPS-XX-99"
			},
			wantErr: `targets unknown control "OSPS-XX-99"`,
		},
	}

	v := NewValidator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := validBaseline()
			tt.mutate(b)

			err := v.Check(b)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected an error containing %q, got none", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected an error containing %q, got: %v", tt.wantErr, err)
			}
		})
	}
}
