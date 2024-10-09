### Maturity Level 1 (Sandbox)
 
| Category | Criteria Statement |
|---|---|
| Access Control | The project's version control system MUST require multi-factor authentication (MFA) for users when modifying the project repository settings or accessing sensitive data. |
| Access Control | The project's version control system MUST restrict collaborator permissions to the lowest available privileges by default. |
| Access Control | The project's version control system MUST prevent accidental direct commits against the primary branch. |
| Access Control | The project's version control system MUST prevent accidental deletion of the primary branch. |
| Build & Release | The project's build and release pipelines MUST NOT execute arbitrary code that is input from outside of the build script. |
| Build & Release | All software releases MUST be assigned a standardized unique version identifier. |
| Build & Release | Any websites related to the project development and release MUST be delivered using SSH or HTTPS. |
| Documentation | The project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles. |
| Documentation | The project documentation MUST provide user guides for all basic functionality. |
| Documentation | The project documentation MUST include an explanation of the contribution process. |
| Quality | The version control system MUST be publicly readable and have a static URL. |
| Quality | The version control system MUST contain a publicly readable record of all changes made, who made the changes, and when the changes were made. |
| Legal | The version control system MUST require all committers to assert that they are legally authorized to commit the associated contributions on every commit. |
| Legal | The license for the source code MUST be written in a standardized format approved by the OSI or FSF. |
| Legal | The license for the source code MUST be maintained in a standard location within the version control system repository. |
| Legal | The license for the released assets MUST be written in a standardized format approved by the OSI or FSF, if different from the source code license. |

### Maturity Level 2 (Incubating)

| Category | Criteria Statement |
|---|---|
| Access Control | The project's permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated. |
| Access Control | The project's documentation MUST have a policy that contributors are reviewed prior to granting escalated permissions to sensitive resources. |
| Build & Release | All released software assets MUST be created with consistent, automated build and release processes. |
| Build & Release | All build and release pipelines MUST use standardized tooling to ingest dependencies at build time. |
| Build & Release | All software releases MUST include a descriptive log of functional and security modifications. |
| Documentation | The project documentation MUST include a policy for coordinated vulnerability reporting, with a clear timeframe for response. |
| Documentation | The project documentation MUST include a mechanism for reporting defects. |
| Documentation | The project documentation MUST include a guide for contributors and approvers that includes the requirements for acceptable contributions.|
| Documentation | The project documentation MUST provide design documentation demonstrating all actions and actors within the system. |
| Quality | All released software assets MUST be associated with a machine-readable list of all direct and transitive dependencies with their associated version identifier. |
| Quality | Any automated commit status checks MUST pass or require manual intervention prior to merge. |
| Quality | All codebases within the project MUST enforce security requirements as applicable to the status and intent of the codebase. |
| Quality | The version control system MUST NOT contain generated executable artifacts. |

### Maturity Level 3 (Graduated)

| Category | Criteria Statement |
|---|---|
| Access Control | The project's version control system MUST require multi-factor authentication (MFA) that does not include SMS for users when modifying the project repository settings or accessing sensitive data. |
| Build & Release | All software releases MUST incorporate a policy to address applicable Software Composition Analysis results. |
| Documentation | The project documentation MUST define a cadence in which known vulnerabilities are evaluated, and exploitable vulnerabilities are either fixed or verified as unexploitable. |
| Documentation | The project documentation MUST provide reference documentation that describes the external application programmer and human interface (both input and output) of the software. |
