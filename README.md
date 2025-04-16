# Open Source Project Security Baseline

The Open Source Project Security Baseline (OSPS Baseline) is designed to act as a minimum definition of requirements for a project relative to its maturity level.

All definitions are maintained in YAML format for tandem machine and human readability.

## Baseline Structure

Each entry has the following values:

- **ID**:
  - Entries are of the form OSPS-_Category_-_Index_ where
    - *Category* is a two-letter abbreviated form of the categories listed below
    - *Index* is a sequentially-assigned two-digit number. Numbers are unique within a category but not between categories
- **Maturity Level**:
  - Level 1: for any code or non-code project with any number of maintainers or users
  - Level 2: for any code project that has at least 2 maintainers and a small number of consistent users
  - Level 3: for any code project that has a large number of consistent users
- **Family** (see corresponding yaml files for descriptions):
  - [Access Control](baseline/OSPS-AC.yaml)
  - [Build & Release](baseline/OSPS-BR.yaml)
  - [Documentation](baseline/OSPS-DO.yaml)
  - [Governance](baseline/OSPS-GV.yaml)
  - [Legal](baseline/OSPS-LE.yaml)
  - [Quality](baseline/OSPS-QA.yaml)
  - [Security Assessment](baseline/OSPS-SA.yaml)
  - [Vulnerability Management](baseline/OSPS-VM.yaml)
- **Title**:
  - A concise statement of the requirement
  - Contains `MUST` or `MUST NOT` and is written in present tense
  - The term before MUST/NOT is the _subject_ of the requirement
  - Terms following MUST/NOT describe the _required behavior_
- **Objective**:
  - A concise statement of the goal of the requirement
  - Written in present tense and describes the desired outcome
- **Assessment requirement(s)**:
  - A concise description of how to meet the requirement
  - Written in present tense and describes the steps to take to meet the requirement
  - May outline recommendations, examples, or best practices

## Contribution, Governance, & Security

Contributions are always welcome via pull request or as issues, and can also be discussed on the [`#sig-security-baseline` channel on OpenSSF Slack](https://openssf.slack.com/archives/C07DC6TT2QY). Refer to the governance documentation for information about [how the project operates] and [how to report security-related issues].

You can also join the biweekly meeting.
See the [OpenSSF calendar](https://calendar.google.com/calendar/u/0?cid=czYzdm9lZmhwNWk5cGZsdGI1cTY3bmdwZXNAZ3JvdXAuY2FsZW5kYXIuZ29vZ2xlLmNvbQ) for details.

### Antitrust Policy Notice

Linux Foundation meetings involve participation by industry competitors, and it is the intention of the Linux Foundation to conduct all of its activities in accordance with applicable antitrust and competition laws. It is therefore extremely important that attendees adhere to meeting agendas, and be aware of, and not participate in, any activities that are prohibited under applicable US state, federal or foreign antitrust and competition laws.

Examples of types of actions that are prohibited at Linux Foundation meetings and in connection with Linux Foundation activities are described in the Linux Foundation Antitrust Policy available at http://www.linuxfoundation.org/antitrust-policy. If you have questions about these matters, please contact your company counsel, or if you are a member of the Linux Foundation, feel free to contact Andrew Updegrove of the firm of Gesmer Updegrove LLP, which provides legal counsel to the Linux Foundation.

[how the project operates]: governance/GOVERNANCE.md
[how to report security-related issues]: governance/SECURITY.md
