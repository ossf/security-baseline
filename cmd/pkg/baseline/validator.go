// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"errors"
	"fmt"
	"slices"

	"github.com/ossf/security-baseline/pkg/types"
)

func NewValidator() *Validator {
	return &Validator{}
}

type Validator struct {
}

// Check verifies the data parsed for consistency and completeness
func (v *Validator) Check(b *types.Baseline) error {
	var entryIDs []string
	var errs = []error{}
	for _, category := range b.Categories {
		for _, entry := range category.Controls {
			if slices.Contains(entryIDs, entry.ID) {
				errs = append(errs, fmt.Errorf("duplicate ID for 'control' for %s", entry.ID))
			}
			if entry.ID == "" {
				errs = append(errs, fmt.Errorf("missing ID for 'control' %s", entry.ID))
			}
			if entry.Title == "" {
				errs = append(errs, fmt.Errorf("missing 'control' text for %s", entry.ID))
			}
			// For after all fields are populated:
			// if entry.Rationale == "" {
			//   errs = append(errs, fmt.Errorf("missing 'rationale' for %s", entry.ID))
			// }
			// if entry.Details == "" {
			//   errs = append(errs, fmt.Errorf("missing 'details' for %s", entry.ID))
			// }
			entryIDs = append(entryIDs, entry.ID)
		}
	}

	return errors.Join(errs...)
}
