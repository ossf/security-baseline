# Security Baseline Project Governance

As a developing project, OSPS Baseline aims to have a quick development cycle where decisions and community issues are resolved promptly while capturing the input of interested stakeholders.

OSPS Baseline has no formal collegiate body in charge of steering. **Decisions are guided by the consensus of community members who have achieved maintainer status.**

While maintainer consensus shall be the process for decision making, all issues and proposals shall be governed by the project's [guiding principles].

## OpenSSF Security Baseline Special Interest Group

The Open Source Project Security Baseline (OSPS Baseline) is produced by the OpenSSF's Baseline SIG, part of the [BEST Working Group](https://github.com/ossf/wg-best-practices-os-developers).

Refer to the [OpenSSF Community Calendar](https://openssf.org/getinvolved/) for SIG meeting times, meeting notes, and links to past recordings.

- **SIG Lead:** Ben Cotton (@funnnelfiasco)

## Guiding Governance Principles

Any issues or proposals brought to the project's maintainers shall be framed in the OSPS Baseline guiding principles. Proposals not adhering to said principles shall not be considered for consensus.

### Favor Simplicity

The goal of OSPS Baseline is to create a minimal and efficient standard that can be quickly ingested by any project. Simple is better.

### Ensure Stability

Any enhancements to the OSPS Baseline and its delivery tooling must not cause breaking changes prior to a scheduled release.

### Cautious Incremental Improvement

New entries must be added with caution, and breaking changes should be extremely rare, even on a scheduled release. Incremental development may enter the repository between releases only if it is mapped to an open GitHub Issue.

## Maintainer Consensus

To reach a decision on an issue or proposal, the proponents must seek maintainer consensus. In the context of this document, "maintainer consensus" means collecting approvals from at least 51% of the current maintainer body, with enough time for all maintainers to review (usually 2 business days), and without a dissenting maintainer opinion.

## Review Controls for OSPS Baseline Repository

Any changes intended to be merged in the OSPS Baseline repository shall meet the following minimal criteria:

- Commits must be signed off.
- Pull requests must be approved by at least two of the project's maintainers.

Any repository under the OSPS Baseline organization may impose additional requirements to approve pull requests as long as these minimal requirements are met.

## Maintainer Status

Any community member may be considered as a candidate for maintainer status under the following conditions:

- A [sponsoring committee] may nominate a community member at any time.
- Any community member may self-nominate as a maintainer candidate after actively contributing to OSPS Baseline on at least a monthly basis for six months or more.

Nomination shall be in the form of a pull request to update the OSPS Baseline's [MAINTAINERS.md].

After the nomination is filed and deemed valid, [maintainer consensus] may be sought. Upon achieving consensus, the PR may be merged to confirm the new maintainer.

### Sponsoring Committees

To nominate a community member as a maintainer candidate, a group of maintainers may file a nomination. The committee shall meet the following criteria to be qualified to file the nomination:

- The number of members in the committee shall not be less than two (2).
- Whenever the number of organizations with maintainers in the project is more than two (2), committee members shall be from different organizations.

### Continued Maintainer Status

Once confirmed as a maintainer, continuation is contingent on regular activity and adherence to the [OpenSSF Code of Conduct](https://openssf.org/community/code-of-conduct/).

### Emeritus Maintainers

Emeritus maintainers will be listed in a separate section on the OSPS Baseline's [MAINTAINERS.md].

A maintainer who is not currently active on the project may be given Emeritus status by default after six months of no activity, such as pull request interactions or GitHub Issue interactions. A maintainer may also assign themselves Emeritus status through a pull request.

A maintainer may become active from Emeritus status through [maintainer consensus] and a corresponding pull request.

## Revisions to the Governance Model

The project's governance model shall be revisited every six months to address the needs of the community, as the project recognizes that communities need to steer themselves according to their size, members, and other factors that shape their complexity.

At any point, a OSPS Baseline Enhancement Proposal may be opened to redefine the project's governance. To be accepted, governing model proposals shall be approved by a qualified majority consisting of a minimum of 66% favorable votes of all active maintainers.

[MAINTAINERS.md]: /MAINTAINERS.md
[Maintainer Consensus]: #maintainer-consensus
[Sponsoring Committee]: #sponsoring-committees
[guiding principles]: #guiding-governance-principles
