# Open Source Project Security Baseline

The Open Source Project Security Baseline (OSPS Baseline) is designed to act as a minimum definition of requirements for a project relative to it's maturity level.

All definitions are maintained in YAML format for tandem machine and human readibility.

## Baseline Structure

Each entry has the following values:

- **ID**:
  - Entries 1-39 are reserved for maturity level 1
  - Entries 40-69 are reserved for maturity level 2
  - Entries 70-99 are reserved for maturity level 3
- **Maturity Level**:
  - Level 1: for any code or non-code project with any number of maintainers or users
  - Level 2: for any code project that has at least 2 maintainers and a small number of consistent users
  - Level 3: for any code project that has a large number of consistent users
- **Category**:
  - Access Control
  - Build & Release
  - Documentation
  - Quality
  - Legal
- **Criteria**: 
  - A concise statement of the requirement
  - Contains `MUST` or `MUST NOT` and is written in present tense
  - The term before MUST/NOT is the _subject_ of the requirement
  - erms following MUST/NOT describes the _required behavior_
- **Objective**:
  - A concise statement of the goal of the requirement
  - Written in present tense and describes the desired outcome
- **Implementation**: 
  - A concise description of how to meet the requirement
  - Written in present tense and describes the steps to take to meet the requirement
  - May outline recommendations, examples, or best practices

## Contribution, Governance, & Security

Contributions are always welcome via pull request or GitHub Discussions. Refer to the governance documentation for information about [how the project operates] and [how to report security-related issues].

### Antitrust Policy Notice

Linux Foundation meetings involve participation by industry competitors, and it is the intention of the Linux Foundation to conduct all of its activities in accordance with applicable antitrust and competition laws. It is therefore extremely important that attendees adhere to meeting agendas, and be aware of, and not participate in, any activities that are prohibited under applicable US state, federal or foreign antitrust and competition laws.

Examples of types of actions that are prohibited at Linux Foundation meetings and in connection with Linux Foundation activities are described in the Linux Foundation Antitrust Policy available at http://www.linuxfoundation.org/antitrust-policy. If you have questions about these matters, please contact your company counsel, or if you are a member of the Linux Foundation, feel free to contact Andrew Updegrove of the firm of Gesmer Updegrove LLP, which provides legal counsel to the Linux Foundation.

[how the project operates]: governance/GOVERNANCE.md
[how to report security-related issues]: governance/SECURITY.md