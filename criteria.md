# Open Source Project Security Baseline

This document contains the criteria for the baseline. Where applicable, terms must be interpreted according to the definitions in the [OSPS Lexicon](./lexicon.md)

## Criteria

### Maturity Level 1 (Sandbox)
 
| ID | Category | Criteria Statement |
|---|---|---|
| OSPS-01 | Access Control | The project's [version control system] MUST require [multi-factor authentication] for users when modifying the project [repository] settings or accessing sensitive data. |
| OSPS-02 | Access Control | The project's [version control system] MUST restrict [collaborator] permissions to the lowest available privileges by default. |
| OSPS-03 | Access Control | The project's [version control system] MUST prevent accidental direct [commits] against the [primary branch]. |
| OSPS-04 | Access Control | The project's [version control system] MUST prevent accidental deletion of the [primary branch]. |
| OSPS-05 | Build & Release | The project's [build and release pipelines] MUST NOT execute [arbitrary code] that is input from outside of the build script. |
| OSPS-06 | Build & Release | All [releases] and [released software assets] MUST be assigned a [unique version identifier]. |
| OSPS-07 | Build & Release | Any websites, API responses or other services related to the project development and release MUST be delivered using SSH,HTTPS or other encrypted channels8. |
| OSPS-09 | Documentation | The project MUST have one or more mechanisms for public discussions about proposed [changes] and usage obstacles. |
| OSPS-10 | Documentation | The [project documentation] MUST include an explanation of the contribution process. |
| OSPS-11 | Documentation | The [project documentation] MUST provide user guides for all basic functionality. |
| OSPS-12 | Quality | The [version control system] MUST be publicly readable and have a static URL. |
| OSPS-13 | Quality | The [version control system] MUST contain a publicly readable record of all [changes] made, who made the [changes], and when the [changes] were made. |
| OSPS-14 | Legal | The [version control system] MUST require all [code contributors] to assert that they are legally authorized to [commit] the associated contributions on every [commit]. |
| OSPS-15 | Legal | The [license] for the source code MUST be written in a standardized format approved by the [OSI] or [FSF]. |
| OSPS-16 | Legal | The [license] for the source code MUST be maintained in a standard location within the project's [repository]. |
| OSPS-17 | Legal | The [license] for the [released software assets] MUST be written in a standardized format approved by the [OSI] or [FSF], if different from the source code [license]. |

### Maturity Level 2 (Incubating)

| ID | Category | Criteria Statement |
|---|---|---|
| OSPS-40 | Access Control | The project's permissions in [CI/CD pipelines] MUST be configured to the lowest available privileges except when explicitly elevated. |
| OSPS-41 | Access Control | The [project documentation] MUST have a policy that [code contributors] are reviewed prior to granting escalated permissions to sensitive resources. |
| OSPS-42 | Build & Release | All [released software assets] MUST be created with consistent, automated [build and release pipelines]. |
| OSPS-43 | Build & Release | All [build and release pipelines] MUST use standardized tooling to ingest dependencies at build time. |
| OSPS-44 | Build & Release | All [releases] MUST include a descriptive log of functional and security modifications. |
| OSPS-45 | Documentation | The [project documentation] MUST include a policy for coordinated [vulnerability reporting], with a clear timeframe for response. |
| OSPS-46 | Documentation | The [project documentation] MUST include a mechanism for reporting [defects]. |
| OSPS-47 | Documentation | The [project documentation] MUST include a guide for [code contributors] that includes the requirements for acceptable contributions. |
| OSPS-48 | Documentation | The [project documentation] MUST provide design documentation demonstrating all actions and actors within the system. |
| OSPS-49 | Quality | All [software assets] MUST be [released] with a machine-readable list of all direct and transitive dependencies with their associated version identifier. |
| OSPS-50 | Quality | Any automated [status checks] for [commits]  MUST pass or require manual intervention prior to merge. |
| OSPS-51 | Quality | Any additional code [repositories] produced by the project MUST enforce security requirements as applicable to the status and intent of the respective [codebase]. |
| OSPS-52 | Quality | The [version control system] MUST NOT contain generated executable artifacts. |

### Maturity Level 3 (Graduated)

| ID | Category | Criteria Statement |
|---|---|---|
| OSPS-70 | Access Control | The project's [version control system] MUST require [multi-factor authentication] that does not include SMS for users when modifying the project [repository] settings or accessing sensitive data. |
| OSPS-71 | Build & Release | The [project documentation] MUST include a policy to address applicable [Software Composition Analysis] results prior to any [release]. |
| OSPS-72 | Documentation | The [project documentation] MUST define a cadence in which [known vulnerabilities] are evaluated, and [exploitable vulnerabilities] are either fixed or verified as unexploitable. |
| OSPS-73 | Documentation | The [project documentation] MUST include descriptions of all input and output interfaces of the [released software assets]. |

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
