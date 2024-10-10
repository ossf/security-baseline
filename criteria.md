# Open Source Project Security Baseline

This document contains the criteria for the baseline. Where applicable, terms should be interpreted according to the definitions in the [OSPS Lexicon](./lexicon.md)

## Criteria

### Maturity Level 1 (Sandbox)
 
| Category | Criteria Statement |
|---|---|
| Access Control | The project's [version control system] MUST require [multi-factor authentication] for users when modifying the project [repository] settings or accessing sensitive data. |
| Access Control | The project's [version control system] MUST restrict [collaborator] permissions to the lowest available privileges by default. |
| Access Control | The project's [version control system] MUST prevent accidental direct [commit] against the [primary branch]. |
| Access Control | The project's [version control system] MUST prevent accidental deletion of the [primary branch]. |
| Build & Release | The project's [build and release pipelines] MUST NOT execute [arbitrary code] that is input from outside of the build script. |
| Build & Release | All [releases] and [released software assets] MUST be assigned a standardized [unique version identifier]. |
| Build & Release | Any websites related to the project development and release MUST be delivered using SSH or HTTPS. |
| Documentation | The project MUST have one or more mechanisms for public discussions about proposed [changes] and usage obstacles. |
| Documentation | The [project documentation] MUST provide user guides for all basic functionality. |
| Documentation | The [project documentation] MUST include an explanation of the contribution process. |
| Quality | The [version control system] MUST be publicly readable and have a static URL. |
| Quality | The [version control system] MUST contain a publicly readable record of all [changes] made, who made the [changes], and when the [changes] were made. |
| Legal | The [version control system] MUST require all [code contributors] to assert that they are legally authorized to [commit] the associated contributions on every [commit]. |
| Legal | The license for the source code MUST be written in a standardized format approved by the OSI or FSF. |
| Legal | The license for the source code MUST be maintained in a standard location within the project's [repository]. |
| Legal | The license for the [released software assets] MUST be written in a standardized format approved by the OSI or FSF, if different from the source code license. |

### Maturity Level 2 (Incubating)

| Category | Criteria Statement |
|---|---|
| Access Control | The project's permissions in [CI/CD pipelines] MUST be configured to the lowest available privileges except when explicitly elevated. |
| Access Control | The [project documentation] MUST have a policy that [code contributors] are reviewed prior to granting escalated permissions to sensitive resources. |
| Build & Release | All [released software assets] MUST be created with consistent, automated [build and release pipeline]. |
| Build & Release | All [build and release pipelines] MUST use standardized tooling to ingest dependencies at build time. |
| Build & Release | All [releases] MUST include a descriptive log of functional and security modifications. |
| Documentation | The [project documentation] MUST include a policy for coordinated [vulnerability reporting], with a clear timeframe for response. |
| Documentation | The [project documentation] MUST include a mechanism for reporting [defects]. |
| Documentation | The [project documentation] MUST include a guide for [code contributors] and approvers that includes the requirements for acceptable contributions.|
| Documentation | The [project documentation] MUST provide design documentation demonstrating all actions and actors within the system. |
| Quality | All [released software assets] MUST be [released] with a machine-readable list of all direct and transitive dependencies with their associated version identifier. |
| Quality | Any automated [status checks] for [commits]  MUST pass or require manual intervention prior to merge. |
| Quality | Any additional code [reponsitories] produced by the project MUST enforce security requirements as applicable to the status and intent of the [codebase]. |
| Quality | The [version control system] MUST NOT contain generated executable artifacts. |

### Maturity Level 3 (Graduated)

| Category | Criteria Statement |
|---|---|
| Access Control | The project's [version control system] MUST require [multi-factor authentication] that does not include SMS for users when modifying the project [repository] settings or accessing sensitive data. |
| Build & Release | The [project documentation] MUST include a policy to address applicable [Software Composition Analysis] results prior to any [release]. |
| Documentation | The [project documentation] MUST define a cadence in which [known vulnerabilities] are evaluated, and [exploitable vulnerabilities] are either fixed or verified as unexploitable. |
| Documentation | The [project documentation] MUST include descriptions of all input and output interfaces of the [released software assets]. |

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
[known vulnerabilities]: ./lexicon.md#known-vulnerabilities
[multi-factor authentication]: ./lexicon.md#multi-factor-authentication-mfa
[primary branch]: ./lexicon.md#primary-branch
[project documentation]: ./lexicon.md#project-documentation
[released]: ./lexicon.md#release
[released software assets]: ./lexicon.md#released-software-assets
[releases]: ./lexicon.md#release
[repository]: ./lexicon.md#repository
[reponsitories]: ./lexicon.md#repository
[Software Composition Analysis]: ./lexicon.md#software-composition-analysis-sca
[status checks]: ./lexicon.md#status-checks
[unique version identifier]: ./lexicon.md#unique-version-identifier
[version control system]: ./lexicon.md#version-control-system
[vulnerability reporting]: ./lexicon.md#vulnerability-reporting
