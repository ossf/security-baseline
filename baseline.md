<!-- markdownlint-disable -->

# Open Source Project Security Baseline

## Overview

The Basleine is a set of security criteria that projects should meet to be considered secure. The criteria are organized by maturity level and category. In the detailed subsections you will find the criteria, objectives, and implementation notes.

This document was automatically generated from the [Open Source Project Security Baseline YAML](https://github.com/ossf/security-baselines/blob/main/baselines/ossf-security-baseline.yaml) file.

---

## Baseline Criteria

| ID  | Maturity Level | Category | Criteria |
| --- | -------------- | -------- | -------- |
| OSPS-01 | 1 | Access Control | The project's version control system MUST require multi-factor authentication for collaborators modifying the project repository settings or accessing sensitive data.  |
| OSPS-02 | 1 | Access Control | The project's version control system MUST restrict collaborator permissions to the lowest available privileges by default.  |
| OSPS-03 | 1 | Access Control | The project's version control system MUST prevent unintentional direct commits against the primary branch.  |
| OSPS-04 | 1 | Access Control | The project's version control system MUST prevent unintentional deletion of the primary branch.  |
| OSPS-05 | 1 | Build & Release | The project's build and release pipelines MUST NOT execute arbitrary code that is input from outside of the build script.  |
| OSPS-06 | 1 | Build & Release | All releases and released software assets MUST be assigned a unique version identifier.  |
| OSPS-07 | 1 | Build & Release | Any websites, API responses or other services related to the project development and release MUST be delivered using SSH, HTTPS or other encrypted channels.  |
| OSPS-09 | 1 | Documentation | The project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles.  |
| OSPS-10 | 1 | Documentation | The project documentation MUST include an explanation of the contribution process.  |
| OSPS-11 | 1 | Documentation | The project documentation MUST provide user guides for all basic functionality.  |
| OSPS-12 | 0 | Quality | The project's source code MUST be publicly readable and have a static URL.  |
| OSPS-13 | 1 | Quality | The version control system MUST contain a publicly readable record of all changes made, who made the changes, and when the changes were made.  |
| OSPS-14 | 1 | Legal | The version control system MUST require all code contributors to assert that they are legally authorized to commit the associated contributions on every commit.  |
| OSPS-15 | 1 | Legal | The license for the source code MUST be written in a standardized format approved by the OSI or FSF.  |
| OSPS-16 | 1 | Legal | The license for the source code MUST be maintained in a standard location within the project's repository.  |
| OSPS-17 | 1 | Legal | The license for the released software assets MUST be written in a standardized format approved by the OSI or FSF, if different from the source code license.  |
| OSPS-40 | 2 | Access Control | The project's permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated.  |
| OSPS-41 | 2 | Access Control | The project documentation MUST have a policy that code contributors are reviewed prior to granting escalated permissions to sensitive resources.  |
| OSPS-42 | 2 | Build & Release | All released software assets MUST be created with consistent, automated build and release pipelines.  |
| OSPS-43 | 2 | Build & Release | All build and release pipelines MUST use standardized tooling to ingest dependencies at build time.  |
| OSPS-44 | 2 | Build & Release | All releases MUST include a descriptive log of functional and security modifications.  |
| OSPS-45 | 2 | Documentation | The project documentation MUST include a policy for coordinated vulnerability reporting, with a clear timeframe for response.  |
| OSPS-46 | 2 | Documentation | The project documentation MUST include a mechanism for reporting defects.  |
| OSPS-47 | 2 | Documentation | The project documentation MUST include a guide for code contributors that includes requirements for acceptable contributions.  |
| OSPS-48 | 2 | Documentation | The project documentation MUST provide design documentation demonstrating all actions and actors within the system.  |
| OSPS-49 | 2 | Quality | All software assets MUST be released with a machine-readable list of all direct and transitive dependencies with their associated version identifier.  |
| OSPS-50 | 2 | Quality | Any automated status checks for commits MUST pass or require manual intervention prior to merge.  |
| OSPS-51 | 2 | Quality | Any additional code repositories produced by the project MUST enforce security requirements as applicable to the status and intent of the respective codebase.  |
| OSPS-52 | 2 | Quality | The version control system MUST NOT contain generated executable artifacts.  |
| OSPS-70 | 3 | Access Control | The project's version control system MUST require multi-factor authentication that does not include SMS for users when modifying the project repository settings or accessing sensitive data.  |
| OSPS-71 | 3 | Build & Release | The project documentation MUST include a policy to address applicable Software Composition Analysis results prior to any release.  |
| OSPS-72 | 3 | Documentation | The project documentation MUST define a cadence in which known vulnerabilities are evaluated, and exploitable vulnerabilities are either fixed or verified as unexploitable.  |
| OSPS-73 | 3 | Documentation | The project documentation MUST include descriptions of all input and output interfaces of the released software assets.  |

---

### OSPS-01 - Access Control

**Criteria:**

The project's version control system MUST
require multi-factor authentication for
collaborators modifying the project
repository settings or accessing sensitive
data.

**Objective:**

Protect against unauthorized access to
sensitive areas of the project's repository,
reducing the risk of account compromise or
insider threats.

**Implementation:**

Require multi-factor authentication for the
project's version control system, requiring
collaborators to provide a second form of
authentication when accessing sensitive data
or modifying repository settings.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-02 - Access Control

**Criteria:**

The project's version control system MUST
restrict collaborator permissions to the
lowest available privileges by default.

**Objective:**

Reduce the risk of unauthorized access to
the project's repository by limiting the
permissions granted to collaborators.

**Implementation:**

Configure the project's version control
system to assign the lowest available
permissions to collaborators by default when
added, granting additional permissions only
when necessary.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-03 - Access Control

**Criteria:**

The project's version control system MUST
prevent unintentional direct commits against
the primary branch.

**Objective:**

Reduce the risk of accidental changes to the
primary branch of the project's repository,
ensuring that due diligence is done before
commits are merged.

**Implementation:**

Set branch protection on the primary branch
in the project's version control system
requiring changes to be made through
pull/merge requests or other review
mechanisms.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-04 - Access Control

**Criteria:**

The project's version control system MUST
prevent unintentional deletion of the
primary branch.

**Objective:**

Protect the primary branch of the project's
repository from accidental deletion,
ensuring that the project's history and
codebase are preserved.

**Implementation:**

Set branch protection on the primary branch
in the project's version control system to
prevent deletion.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-05 - Build & Release

**Criteria:**

The project's build and release pipelines
MUST NOT execute arbitrary code that is
input from outside of the build script.

**Objective:**

Reduce the risk of code injection or other
security vulnerabilities in the project's
build and release processes by restricting
the execution of external code.

**Implementation:**

Ensure that the project's build and release
pipelines do not execute arbitrary code
provided from external sources.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-06 - Build & Release

**Criteria:**

All releases and released software assets
MUST be assigned a unique version
identifier.

**Objective:**

Ensure that each software asset produced by
the  project is uniquely identified,
enabling users to track changes and updates
to the project over time.

**Implementation:**

Assign a unique version identifier to each
release and associated software asset
produced by the project, following a
consistent naming convention or numbering
scheme.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-07 - Build & Release

**Criteria:**

Any websites, API responses or other
services related to the project development
and release MUST be delivered using SSH,
HTTPS or other encrypted channels.

**Objective:**

Protect the confidentiality and integrity
of data transmitted between the project's
services and users, reducing the risk of
eavesdropping or data tampering.

**Implementation:**

Configure the project's websites, API
responses, and other services to use
encrypted channels such as SSH or HTTPS for
data transmission.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-09 - Documentation

**Criteria:**

The project MUST have one or more mechanisms
for public discussions about proposed
changes and usage obstacles.

**Objective:**

Encourages open communication and
collaboration within the project community,
enabling users to provide feedback and
discuss proposed changes or usage
challenges.

**Implementation:**

Establish one or more mechanisms for public 
discussions within the project, such as
mailing lists, instant messaging, or issue
trackers, to facilitate open communication
and feedback.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-10 - Documentation

**Criteria:**

The project documentation MUST include an
explanation of the contribution process.

**Objective:**

Provide guidance to new contributors on how
to participate in the project, outlining the
steps required to submit changes or
enhancements to the project's codebase.

**Implementation:**

Create a CONTRIBUTING.md or CONTRIBUTING/
directory to outline the contribution
process including the steps for submitting
changes, and engaging with the project
maintainers.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-11 - Documentation

**Criteria:**

The project documentation MUST provide user
guides for all basic functionality.

**Objective:**

Ensure that users have a clear and
comprehensive understanding of the project's
current features in order to prevent damage
from misuse or misconfiguration.

**Implementation:**

Create user guides or documentation for all
basic functionality of the project,
explaining how to install, configure, and
use the project's features. If there are any
known dangerous or destructive actions
available, include highly-visible warnings.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-12 - Quality

**Criteria:**

The project's source code MUST be publicly
readable and have a static URL.

**Objective:**

Enable users to access and review the
project's source code and history, promoting
transparency and collaboration within the
project community.

**Implementation:**

Use a common VCS such as GitHub, GitLab, or
Bitbucket. Ensure the repository is publicly
readable. Avoid duplication or mirroring of
repositories unless highly visible
documentation clarifies the primary source.
Avoid frequent changes to the repository
that would impact the repository URL.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-13 - Quality

**Criteria:**

The version control system MUST contain a
publicly readable record of all changes
made, who made the changes, and when the
changes were made.

**Objective:**

Provide transparency and accountability for
changes made to the project's codebase,
enabling users to track modifications and
understand the history of the project.

**Implementation:**

Use a common VCS such as GitHub, GitLab, or
Bitbucket to maintain a publicly readable
commit history. Avoid squashing or rewriting
commits in a way that would obscure the
author of any commits.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-14 - Legal

**Criteria:**

The version control system MUST require all
code contributors to assert that they are
legally authorized to commit the associated
contributions on every commit.

**Objective:**

Ensure that code contributors are aware of
and acknowledge their legal responsibility
for the contributions they make to the
project, reducing the risk of intellectual
property disputes.

**Implementation:**

Include a DCO or CLA in the project's
repository, requiring code contributors to
assert that they are legally authorized to
commit the associated contributions on every
commit. Use a status check to ensure the
assertion is made.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-15 - Legal

**Criteria:**

The license for the source code MUST be
written in a standardized format approved by
the OSI or FSF.

**Objective:**

Ensure that the project's source code is
distributed under a recognized and legally
enforceable license, providing clarity on
how the code can be used and shared by
others.

**Implementation:**

Add a LICENSE file to the project's repo
that is written in a standardized format
approved by the Open Source Initiative (OSI)
or Free Software Foundation (FSF), ensuring
that the terms are clear and enforceable.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-16 - Legal

**Criteria:**

The license for the source code MUST be
maintained in a standard location within
the project's repository.

**Objective:**

Ensure that the project's source code is
distributed with the appropriate license
terms, making it clear to users and
contributors how the code can be used and
shared.

**Implementation:**

Include the project's source code license in
the project's LICENSE file or LICENSE/
directory to provide visibility and clarity
on the licensing terms.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-17 - Legal

**Criteria:**

The license for the released software assets
MUST be written in a standardized format
approved by the OSI or FSF, if different
from the source code license.

**Objective:**

Ensure that the project's released software
assets are distributed under a recognized
and legally enforceable license, separate
from the source code license if necessary,
providing clarity on how the software can be
used and shared.

**Implementation:**

Choose a license for the project's released
software assets that is written in a
standardized format approved by the Open
Source Initiative (OSI) or Free Software
Foundation (FSF).

Only necessary if license is different
from the source code license.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-40 - Access Control

**Criteria:**

The project's permissions in CI/CD pipelines
MUST be configured to the lowest available
privileges except when explicitly elevated.

**Objective:**

Reduce the risk of unauthorized access to
the project's build and release processes by
limiting the permissions granted to steps
within the CI/CD pipelines. Allows for
quicker adoption of third-party services
that do not require elevated permissions.

**Implementation:**

Configure the project's CI/CD pipelines to
assign the lowest available permissions to
users and services by default, elevating
permissions only when necessary for specific
tasks. In some version control systems, this
may be possible at the organizational or
repository level. If not, set permissions at
the top level of the pipeline.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-41 - Access Control

**Criteria:**

The project documentation MUST have a policy
that code contributors are reviewed prior to
granting escalated permissions to sensitive
resources.

**Objective:**

Ensure that code contributors are vetted and
reviewed before being granted elevated
permissions to sensitive resources within
the project, reducing the risk of
unauthorized access or misuse.

**Implementation:**

Publish an enforceable policy in the project
documentation that requires code
contributors to be reviewed and approved
before being granted escalated permissions
to sensitive resources, such as merge
approval or access to secrets.

It is recommended that, vetting includes
establishing a justifiable lineage of
identity such as confirming the
contributor's association with a known
trusted organization.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-42 - Build & Release

**Criteria:**

All released software assets MUST be created
with consistent, automated build and release
pipelines.

**Objective:**

Ensure that the project's software assets
are built and released using consistent and
automated processes, reducing the risk of
errors or inconsistencies in the deployment
and distribution of the software.

**Implementation:**

Implement reproducible build and release
pipelines for all software assets produced
by the project. VCS-integrated pipelines are
recommended to ensure consistency and
automation in the build and release
processes.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-43 - Build & Release

**Criteria:**

All build and release pipelines MUST use
standardized tooling to ingest dependencies
at build time.

**Objective:**

Ensure that the project's build and release
pipelines use standardized tools and
processes to manage dependencies, reducing
the risk of compatibility issues or security
vulnerabilities in the software.

**Implementation:**

Use a common tooling for your ecosystem,
such as package managers or dependency
management tools to ingest dependencies at
build time. This may include using a
dependency file, lock file, or manifest to
specify the required dependencies, which are
then pulled in by the build system.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-44 - Build & Release

**Criteria:**

All releases MUST include a descriptive log
of functional and security modifications.

**Objective:**

Provide transparency and accountability for
changes made to the project's software
releases, enabling users to understand the
modifications and improvements included in
each release.

**Implementation:**

Ensure that all releases include a
descriptive change log. 

It is recommended to ensure that the change
log is human-readable and includes details
beyond commit messages, such as descriptions 
of the security impact or relevance to
different use cases.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-45 - Documentation

**Criteria:**

The project documentation MUST include a
policy for coordinated vulnerability
reporting, with a clear timeframe for
response.

**Objective:**

Establish a process for reporting and
addressing vulnerabilities in the project,
ensuring that security issues are handled
promptly and transparently.

**Implementation:**

Create a SECURITY.md file at the root of the
directory, outlining the project's policy
for coordinated vulnerability reporting.
Include a method for reporting
vulnerabilities. Set expectations for the
how the project will respond and address
reported issues.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-46 - Documentation

**Criteria:**

The project documentation MUST include a
mechanism for reporting defects.

**Objective:**

Enable users and contributors to report
defects or issues with the released software
assets, facilitating communication and
collaboration on defect fixes and
improvements.

**Implementation:**

It is recommended that projects use their
VCS default issue tracker. If an extarnal
source is used, ensure that the project
documentation and contributing guide clearly
and visibly explain how to use the reporting
system.

It is recommended that project documentation
also sets expectations for how defects will
be triaged and resolved.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-47 - Documentation

**Criteria:**

The project documentation MUST include a
guide for code contributors that includes
requirements for acceptable contributions.

**Objective:**

Provide guidance to code contributors on how
to submit changes and enhancements to the
project's codebase, outlining the standards
and expectations for acceptable
contributions.

**Implementation:**

Extend the CONTRIBUTING.md or CONTRIBUTING/
contents in the project documentation to
outline the requirements for acceptable
contributions, including coding standards,
testing requirements, and submission
guidelines for code contributors.

It is recommended that this guide is the
source of truth for both contributors and
approvers.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-48 - Documentation

**Criteria:**

The project documentation MUST provide
design documentation demonstrating all
actions and actors within the system.

**Objective:**

Provide an overview of the project's design
and architecture, illustrating the
interactions and components of the system to
help contributors and security reviewers
understand the internal logic of the
released software assets.

**Implementation:**

Include designs in the project documentation
that explains the actions and actors. Actors
include any subsystem or entity that can
influence another segment in the system.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-49 - Quality

**Criteria:**

All software assets MUST be released with a
machine-readable list of all direct and
transitive dependencies with their
associated version identifier.

**Objective:**

Provide transparency and accountability for
the project's dependencies, enabling users
and contributors to understand the
software's dependencies and versions.

**Implementation:**

This may take the form of a software bill of
materials (SBOM) or a dependency file that
lists all direct and transitive dependencies
such as requirements.txt, package.json, or
Gemfile.lock, or go.sum.

It is recommended to use a CycloneDX or SPDX
file that is auto-generated at build time by
a tool that has been vetted for accuracy.
This enables users to ingest this data in a
standardized approach alongside other
projects in their environment.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-50 - Quality

**Criteria:**

Any automated status checks for commits MUST
pass or require manual intervention prior to
merge.

**Objective:**

Ensure that the project's approvers do not
become accustomed to tolerating failing
status checks, even if arbitrary, because it
increases the risk of overlooking security
vulnerabilities or defects identified by
automated checks.

**Implementation:**

Configure the project's version control
system to require that all automated status
checks pass or require manual intervention
before a commit can be merged into the
primary branch.

It is recommended that any non-blocking 
status checks are NOT configured as a pass
or fail requirement that approvers may be
tempted to bypass.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-51 - Quality

**Criteria:**

Any additional code repositories produced by
the project MUST enforce security
requirements as applicable to the status and
intent of the respective codebase.

**Objective:**

Ensure that additional code repositories or
subprojects produced by the project are held
to a standard that clear and appropriate
for that codebase.

**Implementation:**

The parent project should maintain a list of
any codebases that are considered
subprojects or additional repositories.
Collaborators on those repositories should
identify the proper maturity level and apply
the Open Source Project Security Baseline to
the codebase. Any subproject or repository
which is compiled into the primary project
must be held to the same standard as the
primary project. Others may be held to a
lower standard if they have lower levels of
adoption or are not intended for general
use.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-52 - Quality

**Criteria:**

The version control system MUST NOT contain
generated executable artifacts.

**Objective:**

Reduce the risk of including generated
executable artifacts in the project's
version control system, ensuring that only
source code and necessary files are stored
in the repository.

**Implementation:**

Remove generated executable artifacts
in the project's version control system.

It is recommended that any scenario where
a binary artifact appears critical to a
process such as testing, it should be stored
in a separate repository and fetched during
a specific well-documented pipeline step.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-70 - Access Control

**Criteria:**

The project's version control system MUST
require multi-factor authentication that
does not include SMS for users when
modifying the project repository settings or
accessing sensitive data.

**Objective:**

Ensure that multi-factor authentication
does not allow SMS as a factor, because
SMS lacks encryption and may be vulnerable
to attacks via Signaling System 7,
social engineering, or SIM-swapping.

**Implementation:**

Require multi-factor authentication methods
that do not include SMS for users. Common
alternatives include hardware tokens, mobile
authenticator apps, or biometric
authentication.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-71 - Build & Release

**Criteria:**

The project documentation MUST include a
policy to address applicable Software
Composition Analysis results prior to any
release.

**Objective:**

Ensure that potential vulnerabilities or
licensing issues identified by SCA tools
are addressed before software releases,
reducing the risk of shipping insecure or
non-compliant software.

**Implementation:**

Document a policy in the project to address
applicable Software Composition Analysis 
results before any release, and add status
checks that verify compliance with that 
policy prior to release.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-72 - Documentation

**Criteria:**

The project documentation MUST define a
cadence in which known vulnerabilities are
evaluated, and exploitable vulnerabilities
are either fixed or verified as
unexploitable.

**Objective:**

Establish a process for evaluating and
addressing known vulnerabilities, then
communicate this process to users and
contributors alike.

**Implementation:**

Define a policy in the project
documentation for evaluating known
vulnerabilities, fixing exploitable
vulnerabilities, and verifying unexploitable
vulnerabilities.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

### OSPS-73 - Documentation

**Criteria:**

The project documentation MUST include
descriptions of all input and output
interfaces of the released software assets.

**Objective:**

Provide users and developers with an
understanding of how to interact with the
project's software and integrate it with
other systems, enabling them to use the
software effectively.

**Implementation:**

Document all input and output interfaces of
the released software assets, explaining how
users can interact with the software and
what data is expected or produced.

**Control Mappings:**  
  
_No control mappings identified._

**Security Insights Value:**  
  
_No security insights identified._

**Scorecard Probe:**  
  
_No scorecard probe identified._

---

## Acknowledgments

This baseline was created by community leaders from across the Linux Foundation, including:

- OpenJS Foundation (OpenJS)
- Open Source Security Foundation (OpenSSF)
- Cloud Native Computing Foundation (CNCF)
- Fintech Open Source Foundation (FINOS)

<!-- markdownlint-enable -->
