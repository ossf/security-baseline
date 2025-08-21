// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"errors"
	"fmt"
	"slices"
)

// Check verifies the data parsed for consistency and completeness
func (l *Loader) Validate() error {
	var entryIDs []string
	errs := []error{}
	for _, category := range l.Catalog.ControlFamilies {
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
		}
	}

	return errors.Join(errs...)
}
