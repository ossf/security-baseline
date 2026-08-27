// SPDX-License-Identifier: Apache-2.0

package schema

import "github.com/gemaraproj/gemara@v1"

// A mapping-reference id is interpolated verbatim into the anchor that links a
// framework relation to its row in the External Frameworks table, so it has to
// be usable as one. A space is the trap: it renders a broken link and nothing
// else complains. The empty string is excluded by the leading character class,
// which also stops a mapping document from silently dropping out of the render.
_refID: =~"^[A-Za-z0-9][A-Za-z0-9._-]*$"

// #OSPSBaseline layers OSPS-specific constraints on top of the Gemara #ControlCatalog.
#OSPSBaseline: gemara.#ControlCatalog & {
	let _agID = =~"^maturity-"

	metadata: "applicability-groups": [{id: _agID}, ...{id: _agID}]
	metadata: "mapping-references": [...{id: _refID}]
}

// #OSPSMapping layers OSPS-specific constraints on top of the Gemara #MappingDocument.
// Every OSPS mapping document maps FROM the osps-baseline catalog OUT to an
// external framework, so the source reference is pinned to "osps-baseline".
#OSPSMapping: gemara.#MappingDocument & {
	"source-reference": "reference-id": "osps-baseline"
	"target-reference": "reference-id": _refID
}
