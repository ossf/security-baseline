# Open Source Project Security Baseline

The Open Source Project Security Baseline (OSPS Baseline) is designed to act as a minimum definition of requirements for a project relative to its maturity level.

All definitions are maintained in YAML format using the [Simplified Compliance Infrastructure Layer 2 schema](https://github.com/revanite-io/sci?tab=readme-ov-file#layer-2-controls) for tandem machine and human readability.

## Baseline Structure

The baseline is a set of Control Families where each Family has a set of Controls.

- [Access Control](baseline/OSPS-AC.yaml)
- [Build & Release](baseline/OSPS-BR.yaml)
- [Documentation](baseline/OSPS-DO.yaml)
- [Governance](baseline/OSPS-GV.yaml)
- [Legal](baseline/OSPS-LE.yaml)
- [Quality](baseline/OSPS-QA.yaml)
- [Security Assessment](baseline/OSPS-SA.yaml)
- [Vulnerability Management](baseline/OSPS-VM.yaml)

Controls are identified by an ID in the format: `OSPS-${ControlFamilyAbbreviated}`.
For example, Control ID: `OSPS-AC-01` refers to the Access Control (AC) control family.

Controls are described with these attributes:

- **Title**:
  - A concise statement of the requirement
  - Contains `MUST` or `MUST NOT` and is written in present tense
  - The term before MUST/NOT is the _subject_ of the requirement
  - Terms following MUST/NOT describe the _required behavior_
- **Objective**:
  - A concise statement of the goal of the requirement
  - Written in present tense and describes the desired outcome

Each Control has 1 or more Assessment Requirements. These requirements are
identified by an ID in the format `OSPS-${ControlFamilyAbbreviated}-${IndexWithinControlFamily}`.
For example, `OSPS-AC-01.01` refers to the first
Assessment Requirement in the first Control of the Access Control Family:

> When a user attempts to access a sensitive resource in the project's
  version control system, the system MUST require the user to complete
  a multi-factor authentication process.

Each Assessment Requirement has:

- **Text**:
  - A concise statement of the requirement
  - Contains `MUST` or `MUST NOT` and is written in present tense
  - The term before `MUST/NOT` is the subject of the requirement
  - Terms following `MUST/NOT` describe the required behavior
- **Recommendation**:
  - description of how to meet the requirement
  - Written in present tense and describes the steps to take to meet the requirement
  - May provide examples and reference best practices
- **Applicability**:
  - One or more project maturity levels to which the assessment requirement applies

## baseline-compiler

The Security Baseline SIG maintains a Golang tool for validating and rendering Baseline controls. It includes the following commands:

* [compile](#compile): Compile a YAML file of security controls
* completion  Generate the autocompletion script for the specified shell
* help        Help about any command
* [oscal](#oscal): Write the Baseline definition to an OSCAL json catalog
* [validate](#validate): Validate the baseline data files

From the `cmd/` directory, run `go run . [command] [arguments]`

### compile

This command reads the YAML data describing the Baseline controls in `/baseline` and generates two optional Markdown files: a listing of all OSPS Baseline controls for use in a static site generator and a checklist of controls by level.

To produce a file for use in a static site generator (which is how we create baseline.openssf.org!), specify a file argument to `--output`.

To produce a checklist file to use as an aid to evaluating your project's conformance to OSPS Baseline, specify a file argument to `--checklist-output`.

Use the `--help` flag for more options.

### oscal

This command reads the YAML OSPS Baseline data files and translates the catalog of controls into [OSCAL](https://pages.nist.gov/OSCAL/) format. It writes the JSON output to STDOUT by default. You can specify a file with `--out`.

Use the `--help` flag for more options.

### validate

This command validates the correctness of the OSPS Baseline input.

## Contribution, Governance, & Security

Contributions are always welcome via pull request or as issues, and can also be discussed on the [`#sig-security-baseline` channel on OpenSSF Slack](https://openssf.slack.com/archives/C07DC6TT2QY). Refer to the governance documentation for information about [how the project operates] and [how to report security-related issues].

You can also join the biweekly meeting.
See the [OpenSSF calendar](https://calendar.google.com/calendar/u/0?cid=czYzdm9lZmhwNWk5cGZsdGI1cTY3bmdwZXNAZ3JvdXAuY2FsZW5kYXIuZ29vZ2xlLmNvbQ) for details.

### Antitrust Policy Notice

Linux Foundation meetings involve participation by industry competitors, and it is the intention of the Linux Foundation to conduct all of its activities in accordance with applicable antitrust and competition laws. It is therefore extremely important that attendees adhere to meeting agendas, and be aware of, and not participate in, any activities that are prohibited under applicable US state, federal or foreign antitrust and competition laws.

Examples of types of actions that are prohibited at Linux Foundation meetings and in connection with Linux Foundation activities are described in the Linux Foundation Antitrust Policy available at http://www.linuxfoundation.org/antitrust-policy. If you have questions about these matters, please contact your company counsel, or if you are a member of the Linux Foundation, feel free to contact Andrew Updegrove of the firm of Gesmer Updegrove LLP, which provides legal counsel to the Linux Foundation.

[how the project operates]: governance/GOVERNANCE.md
[how to report security-related issues]: governance/SECURITY.md
