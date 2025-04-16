# Open Source Project Security Baseline

The Open Source Project Security Baseline (OSPS Baseline) is designed to act as a minimum definition of requirements for a project relative to it's maturity level.
It is maintained by the [OpenSSF Security Baseline SIG](https://github.com/ossf/security-baseline/blob/main/governance/MAINTAINERS.md) according to the [project governance documentation](https://github.com/ossf/security-baseline/blob/main/governance/GOVERNANCE.md).

For more information on the motivation and purpose, see the [FAQ](faq).

## Versions

Previous versions are presented for historical reference.
Downstream consumers of the OSPS Baseline should specify their compliance against a specific version.
Only the version labeled as "current" should be used for new compliance efforts.

* Current version: [v2025.02.25](versions/2025-02-25)
* Previous versions:
    * (none)
* [In-development version](versions/devel)

Versions are managed according to the [Baseline maintenance process](maintenance).

## Guiding principles

The OSPS Baseline controls help project maintainers understand security best practices and expectations.
Assessing a project's compliance against the controls helps maintainers and project consumers understand where the project excels at security and where it has room to improve.
Project consumers can then use the assessment results to understand how their usage of the project impacts their own security and compliance goals.
Therefore, OSPS Baseline work is:

* **Focused:** Controls only contain *MUST* entries, not *SHOULD*.
* **Realistic:** Controls are practical for project maintainers to implement at the appropriate level for their project.
* **Actionable:** Controls provide specific recommendations.
* **Meaningful:** Controls have an impact on a project's security posture.
Ineffective controls add to maintainer burden.