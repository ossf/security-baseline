# Open Source Project Security Baseline

This document contains the criteria for the baseline. Where applicable, terms must be interpreted according to the definitions in the [OSPS Lexicon](./lexicon.md)

## Criteria

### Maturity Level 1 (Sandbox)
 
| ID | Category | Criteria Statement |
|---|---|---|
| mfa | Access Control | The project's [version control system] MUST require [multi-factor authentication] for users when modifying the project [repository] settings or accessing sensitive data. |
| restrict-privileges | Access Control | The project's [version control system] MUST restrict [collaborator] permissions to the lowest available privileges by default. |
| direct-commits | Access Control | The project's [version control system] MUST prevent accidental direct [commits] against the [primary branch]. |
| branch-deletion | Access Control | The project's [version control system] MUST prevent accidental deletion of the [primary branch]. |
| arbitrary-code | Build & Release | The project's [build and release pipelines] MUST NOT execute [arbitrary code] that is input from outside of the build script. |
| version-id | Build & Release | All [releases] and [released software assets] MUST be assigned a [unique version identifier]. |
| website-delivery | Build & Release | Any websites, API responses or other services related to the project development and release MUST be delivered using SSH,HTTPS or other encrypted channels. |
| public-discussions | Documentation | The project MUST have one or more mechanisms for public discussions about proposed [changes] and usage obstacles. |
| basic-docs | Documentation | The [project documentation] MUST provide user guides for all basic functionality. |
| contribution-docs | Documentation | The [project documentation] MUST include an explanation of the contribution process. |
| public-repo | Quality | The [version control system] MUST be publicly readable and have a static URL. |
| change-tracking | Quality | The [version control system] MUST contain a publicly readable record of all [changes] made, who made the [changes], and when the [changes] were made. |
| contributor-assertion | Legal | The [version control system] MUST require all [code contributors] to assert that they are legally authorized to [commit] the associated contributions on every [commit]. |
| license-standard | Legal | The [license] for the source code MUST be written in a standardized format approved by the [OSI] or [FSF]. |
| license-location | Legal | The [license] for the source code MUST be maintained in a standard location within the project's [repository]. |
| release-license | Legal | The [license] for the [released software assets] MUST be written in a standardized format approved by the [OSI] or [FSF], if different from the source code [license]. |

### Maturity Level 2 (Incubating)

| ID | Category | Criteria Statement |
|---|---|---|
| pipeline-permissions | Access Control | The project's permissions in [CI/CD pipelines] MUST be configured to the lowest available privileges except when explicitly elevated. |
| contributor-review | Access Control | The [project documentation] MUST have a policy that [code contributors] are reviewed prior to granting escalated permissions to sensitive resources. |
| automated-pipelines | Build & Release | All [released software assets] MUST be created with consistent, automated [build and release pipelines]. |
| standardized-deps | Build & Release | All [build and release pipelines] MUST use standardized tooling to ingest dependencies at build time. |
| change-log | Build & Release | All [releases] MUST include a descriptive log of functional and security modifications. |
| vuln-report | Documentation | The [project documentation] MUST include a policy for coordinated [vulnerability reporting], with a clear timeframe for response. |
| defect-reporting | Documentation | The [project documentation] MUST include a mechanism for reporting [defects]. |
| contributor-guide | Documentation | The [project documentation] MUST include a guide for [code contributors] that includes the requirements for acceptable contributions. |
| design-documentation | Documentation | The [project documentation] MUST provide design documentation demonstrating all actions and actors within the system. |
| dependency-list | Quality | All [software assets] MUST be [released] with a machine-readable list of all direct and transitive dependencies with their associated version identifier. |
| status-checks | Quality | Any automated [status checks] for [commits]  MUST pass or require manual intervention prior to merge. |
| repo-security | Quality | Any additional code [repositories] produced by the project MUST enforce security requirements as applicable to the status and intent of the respective [codebase]. |
| no-executables | Quality | The [version control system] MUST NOT contain generated executable artifacts. |

### Maturity Level 3 (Graduated)

| Category | Criteria Statement |
|---|---|
| mfa-sms | Access Control | The project's [version control system] MUST require [multi-factor authentication] that does not include SMS for users when modifying the project [repository] settings or accessing sensitive data. |
| sca-policy | Build & Release | The [project documentation] MUST include a policy to address applicable [Software Composition Analysis] results prior to any [release]. |
| vuln-evaluation | Documentation | The [project documentation] MUST define a cadence in which [known vulnerabilities] are evaluated, and [exploitable vulnerabilities] are either fixed or verified as unexploitable. |
| interface-descriptions | Documentation | The [project documentation] MUST include descriptions of all input and output interfaces of the [released software assets]. |

[arbitrary code]: ./lexicon.md#arbitrary-code
[build and release pipelines]: ./lexicon.md#build-and-release-pipeline
[build and release pipeline]: ./lexicon.md#build-and-release-pipeline
[changes]: ./lexicon.md#changes
[CI/CD pipelines]: ./lexicon.md#cicd-pipelines
[code contributors]: ./lexicon.md#code-contributors
[codebase]: ./lexicon.md#codebase
[collaborator]: ./lexicon.md#collaborator
[commit]: ./lexicon.md#commit
[commits]: ./lexicon.md#commit
[defects]: ./lexicon.md#defects
[exploitable vulnerabilities]: ./lexicon.md#exploitable-vulnerabilities
[license]: ./lexicon.md#license
[known vulnerabilities]: ./lexicon.md#known-vulnerabilities
[multi-factor authentication]: ./lexicon.md#multi-factor-authentication-mfa
[primary branch]: ./lexicon.md#primary-branch
[project documentation]: ./lexicon.md#project-documentation
[released]: ./lexicon.md#release
[released software assets]: ./lexicon.md#released-software-assets
[releases]: ./lexicon.md#release
[repository]: ./lexicon.md#repository
[repositories]: ./lexicon.md#repository
[Software Composition Analysis]: ./lexicon.md#software-composition-analysis-sca
[status checks]: ./lexicon.md#status-checks
[unique version identifier]: ./lexicon.md#unique-version-identifier
[version control system]: ./lexicon.md#version-control-system
[vulnerability reporting]: ./lexicon.md#vulnerability-reporting

[FSF]: https://www.fsf.org/
[OSI]: https://opensource.org/
