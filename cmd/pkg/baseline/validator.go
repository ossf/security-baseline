// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ossf/security-baseline/pkg/types"
)

func NewValidator() *Validator {
	return &Validator{}
}

type Validator struct{}

// Check verifies the data parsed for consistency and completeness
func (v *Validator) Check(b *types.Baseline) error {
	var entryIDs []string
	assessmentIDs := []string{}
	errs := []error{}
	for _, category := range b.Catalog.ControlFamilies {
		for _, entry := range category.Controls {
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
				if slices.Contains(assessment.Applicability, "retired") {
					if len(assessment.Applicability) > 1 {
						errs = append(errs, fmt.Errorf("assessment %s in control %s is marked as 'retired' but also has other applicability levels", assessment.Id, entry.Id))
					}
					if !strings.HasPrefix(assessment.Text, "Retired in") {
						errs = append(errs, fmt.Errorf("assessment %s in control %s is marked as 'retired' but does not start with 'Retired in'", assessment.Id, entry.Id))
					}
				} else {
					baseMaturity := -1
					for i, maturity := range assessment.Applicability {
						if maturity[:len(maturity)-1] != "Maturity Level " {
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
				}

				assessmentIDs = append(assessmentIDs, assessment.Id)
			}
		}
	}

	return errors.Join(errs...)
}
