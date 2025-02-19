# OSPS Baseline Maintenance Process

* Normal text fixes to the controls will be accepted via pull request and reviewed by the baseline project maintainers.
Allowed changes are corrections to spelling/typos, grammar corrections, or enhancements to the supplementary text supporting the controls, including: Objective, Implementation, and Control Mappings.
At least two project maintainers must review and approve these changes.
* Substantive changes to controls, including changes to text that alters the originally stated meaning, new controls proposals, or removal of controls will be documented in GitHub PR(s) and reviewed regularly by the Baseline project maintainers for inclusion in the next release.
These changes may reflect changes to global cybersecurity regulations and frameworks or changes in norms around application/project security practices.
Any such substantive changes must be approved by a majority of the project's maintainers.
* As appropriate, but at least annually, the Baseline project maintainers will evaluate the set of controls and, if necessary, publish a new version of the Baseline.
Previous versions of the Baseline will remain available, but are stable and not subject to change, except for minor changes to fix technical or typographic errors.
* Any changes to the Baseline will be reflected within the Compliance Matrix, with new requirements flagged where the Baseline Controls are appropriate.
* Versions will follow a calendar-based identification system, using the `YYYY-MM-DD` format.
* Downstream stakeholders will be notified via the project's mailing list on the changes and updates.

## Identifiers

* Identifiers for retired controls MUST NOT be reused.
Retired identifiers will remain in the source yaml files, clearly marked.
* Substantial changes to the meaning of a control will be treated as a new control, resulting in a new identifier.
Minor changes, including a change in level, between Baseline versions will not result in a new identifier.
* The numeric portion of identifiers are assigned sequentially per category.
Within a category, identifiers do not carry additional meaning.
