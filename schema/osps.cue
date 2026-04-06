// SPDX-License-Identifier: Apache-2.0

package schema

import "github.com/gemaraproj/gemara@v1"

// #OSPSBaseline layers OSPS-specific constraints on top of the Gemara #ControlCatalog.
#OSPSBaseline: gemara.#ControlCatalog & {
	let _agID = =~"^maturity-"

	metadata: "applicability-groups": [{id: _agID}, ...{id: _agID}]
}
