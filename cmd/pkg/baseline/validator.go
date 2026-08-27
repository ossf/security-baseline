// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gemaraproj/go-gemara"

	"github.com/ossf/security-baseline/pkg/types"
)

func NewValidator() *Validator {
	return &Validator{}
}

type Validator struct{}

// Check verifies the data parsed for consistency and completeness
func (v *Validator) Check(b *types.Baseline) error {
	entryIDs := make([]string, 0, len(b.Catalog.Controls))
	assessmentIDs := []string{}
	errs := []error{}
	for i := range b.Catalog.Controls {
		entry := &b.Catalog.Controls[i]
		if slices.Contains(entryIDs, entry.Id) {
			errs = append(errs, fmt.Errorf("duplicate ID for 'control' for %s", entry.Id))
		}
		if entry.Id == "" {
			errs = append(errs, fmt.Errorf("missing ID for 'control' %s", entry.Id))
		}
		if entry.Title == "" {
			errs = append(errs, fmt.Errorf("missing 'control' text for %s", entry.Id))
		}
		entryIDs = append(entryIDs, entry.Id)

		for _, assessment := range entry.AssessmentRequirements {
			if slices.Contains(assessmentIDs, assessment.Id) {
				errs = append(errs, fmt.Errorf("duplicate ID for 'assessment' for %s", assessment.Id))
			}
			if assessment.Id == "" {
				errs = append(errs, fmt.Errorf("missing ID for 'assessment' for %s", entry.Id))
			}

			if len(assessment.Applicability) == 0 {
				errs = append(errs, fmt.Errorf("missing 'applicability' for assessment %s in control %s", assessment.Id, entry.Id))
			}
			if assessment.State == gemara.LifecycleRetired {
				if !strings.HasPrefix(assessment.Text, "Retired in") {
					errs = append(errs, fmt.Errorf("assessment %s in control %s has state 'Retired' but text does not start with 'Retired in'", assessment.Id, entry.Id))
				}
			}
			baseMaturity := -1
			for i, maturity := range assessment.Applicability {
				if maturity[:len(maturity)-1] != "maturity-" {
					errs = append(errs, fmt.Errorf("invalid maturity level '%s' for assessment %s in control %s", maturity, assessment.Id, entry.Id))
					continue
				}
				var err error
				if i == 0 {
					baseMaturity, err = strconv.Atoi(maturity[len(maturity)-1:])
					if err != nil {
						errs = append(errs, fmt.Errorf("invalid maturity level '%s' for assessment %s in control %s", maturity, assessment.Id, entry.Id))
						continue
					}
				} else {
					maturityInt, err := strconv.Atoi(maturity[len(maturity)-1:])
					if err != nil {
						errs = append(errs, fmt.Errorf("invalid maturity level '%s' for assessment %s in control %s", maturity, assessment.Id, entry.Id))
						continue
					}
					if maturityInt != baseMaturity+i {
						errs = append(errs, fmt.Errorf("applicability entry %d for assessment %s in control %s was %q, expected %d", i+1, assessment.Id, entry.Id, maturity, baseMaturity+i))
					}
				}
			}

			assessmentIDs = append(assessmentIDs, assessment.Id)
		}
	}

	referenceIDs := make([]string, 0, len(b.Catalog.Metadata.MappingReferences))
	for _, ref := range b.Catalog.Metadata.MappingReferences {
		referenceIDs = append(referenceIDs, ref.Id)
	}

	targetedIDs := make([]string, 0, len(b.Mappings))
	for i := range b.Mappings {
		doc := &b.Mappings[i]
		// The mapping documents live outside the catalog, so nothing but this
		// check keeps a mapping's source pointing at a control that actually
		// exists.
		for _, m := range doc.Mappings {
			if !slices.Contains(entryIDs, m.Source) {
				errs = append(errs, fmt.Errorf("mapping %s targets unknown control %q", m.Id, m.Source))
			}
		}
		// The rendered document links each framework relation to the row for
		// this ID in the External Frameworks table, so an ID that is not
		// declared in the catalog metadata renders as a dead anchor. The empty
		// ID is deliberately not exempt: the renderer skips it in silence,
		// dropping the document from both the relations and the crosswalk
		// without failing anything.
		fw := doc.TargetReference.ReferenceId
		if !slices.Contains(referenceIDs, fw) {
			errs = append(errs, fmt.Errorf("mapping document %q targets reference %q, which is not declared in metadata mapping-references", doc.Metadata.Id, fw))
		}
		targetedIDs = append(targetedIDs, fw)
	}

	// The reverse of the check above: a reference declared in metadata with no
	// mapping document behind it renders a table row nothing can link to.
	for _, id := range referenceIDs {
		if !slices.Contains(targetedIDs, id) {
			errs = append(errs, fmt.Errorf("mapping-reference %q is declared in metadata but no mapping document targets it", id))
		}
	}

	// addLinks resolves a name to the first lexicon entry declaring it, and
	// asLink folds a term into its anchor, so two names differing only by case
	// silently share one destination. Synonyms take part in that resolution
	// too, so comparing terms alone misses the collisions that actually
	// mislink. A synonym repeating its own entry's term is harmless -- addLinks
	// skips already-wrapped text -- so only cross-entry collisions are errors.
	declaredBy := make(map[string]int, len(b.Lexicon))
	for i, entry := range b.Lexicon {
		names := append([]string{entry.Term}, entry.Synonyms...)
		for _, name := range names {
			key := strings.ToLower(strings.TrimSpace(name))
			if key == "" {
				continue
			}
			// Ownership is tracked by entry index, not by term: two entries
			// sharing a term must still collide with each other.
			if owner, ok := declaredBy[key]; ok && owner != i {
				errs = append(errs, fmt.Errorf("lexicon name %q in entry %q collides with entry %q", name, entry.Term, b.Lexicon[owner].Term))
				continue
			}
			declaredBy[key] = i
		}
	}

	return errors.Join(errs...)
}
