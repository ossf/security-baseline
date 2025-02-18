// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"log"
	"slices"

	"github.com/ossf/security-baseline/pkg/types"
)

type Validator struct {
}

// Check verifies the data parsed for consistency and completeness
func Check(b *types.Baseline) error {
	var entryIDs []string
	var failed bool
	for _, category := range b.Categories {
		for _, entry := range category.Criteria {
			if slices.Contains(entryIDs, entry.ID) {
				failed = true
				log.Printf("duplicate ID for 'criterion' for %s", entry.ID)
			}
			if entry.ID == "" {
				failed = true
				log.Printf("missing ID for 'criterion' %s", entry.ID)
			}
			if entry.CriterionText == "" {
				failed = true
				log.Printf("missing 'criterion' text for %s", entry.ID)
			}
			// For after all fields are populated:
			// if entry.Rationale == "" {
			// 	failed = true
			// 	log.Printf("missing 'rationale' for %s", entry.ID)
			// }
			// if entry.Details == "" {
			// 	failed = true
			// 	log.Printf("missing 'details' for %s", entry.ID)
			// }
			entryIDs = append(entryIDs, entry.ID)
		}
	}
	if failed {
		return fmt.Errorf("error validating baseline")
	}
	return nil
}
