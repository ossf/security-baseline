# Open Source Project Security Baseline

* Contents
{:toc}

## Overview

The Open Source Project Security (OSPS) Baseline is a set of security criteria that projects should meet to be considered secure.
The criteria are organized by maturity level and category.
In the detailed subsections you will find the criterion, rationale, and details notes.

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Criteria Overview

* [Level 1](#level-1): for any code or non-code project with any number of maintainers or users
* [Level 2](#level-2): for any code project that has at least 2 maintainers and a small number of consistent users
* [Level 3](#level-3): for any code project that has a large number of consistent users


### Level 1

**[OSPS-AC-01](#osps-ac-01)**: The project's [version control system] MUST
require [multi-factor authentication] for
[collaborators] modifying the project
[repository] settings or accessing sensitive
data.

**[OSPS-AC-02](#osps-ac-02)**: The project's [version control system] MUST
restrict [collaborator] permissions to the
lowest available privileges by default.

**[OSPS-AC-03](#osps-ac-03)**: The project's [version control system] MUST
prevent unintentional direct [commits] against
the [primary branch].

**[OSPS-AC-04](#osps-ac-04)**: The project's [version control system] MUST
prevent unintentional deletion of the
[primary branch].

**[OSPS-BR-01](#osps-br-01)**: The project's [build and release pipelines]
MUST NOT execute [arbitrary code] that is
input from outside of the build script.

**[OSPS-BR-03](#osps-br-03)**: Any websites, API responses or other
services involved in the project development
and [release] MUST be delivered using SSH,
HTTPS or other encrypted channels.

**[OSPS-DO-13](#osps-do-13)**: The [project documentation] MUST include a
descriptive statement about the scope and
duration of support.

**[OSPS-GV-02](#osps-gv-02)**: The project MUST have one or more mechanisms
for public discussions about proposed
[changes] and usage obstacles.

**[OSPS-GV-03](#osps-gv-03)**: The [project documentation] MUST include an
explanation of the contribution process.

**[OSPS-LE-02](#osps-le-02)**: The [license] for the source code MUST
meet the OSI Open Source Definition or
the FSF Free Software Definition.

**[OSPS-LE-03](#osps-le-03)**: The [license] for the source code MUST be
maintained in a standard location within
the project's [repository].

**[OSPS-LE-04](#osps-le-04)**: The [license] for the [released software assets]
MUST meet the OSI Open Source Definition or
the FSF Free Software Definition.

**[OSPS-QA-01](#osps-qa-01)**: The project's source code MUST be publicly
readable and have a static URL.

**[OSPS-QA-02](#osps-qa-02)**: The [version control system] MUST contain a
publicly readable record of all [changes]
made, who made the [changes], and when the
[changes] were made.

**[OSPS-SA-02](#osps-sa-02)**: The [project documentation] MUST include
descriptions of all external input and output
interfaces of the [released software assets].


### Level 2

**[OSPS-AC-05](#osps-ac-05)**: The project's permissions in [CI/CD pipelines]
MUST be configured to the lowest available
privileges except when explicitly elevated.

**[OSPS-BR-02](#osps-br-02)**: All [releases] and [released software assets]
MUST be assigned a unique version
identifier for each [release] intended to be
used by users.

**[OSPS-BR-04](#osps-br-04)**: All [released software assets] MUST be created
with consistent, automated build and [release]
pipelines.

**[OSPS-BR-05](#osps-br-05)**: All [build and release pipelines] MUST use
standardized tooling where available to
ingest dependencies at build time.

**[OSPS-BR-06](#osps-br-06)**: All [releases] MUST provide a descriptive log
of functional and security modifications.

**[OSPS-BR-08](#osps-br-08)**: All [released software assets] MUST be signed
or accounted for in a signed manifest
including each asset's cryptographic hashes.

**[OSPS-DO-03](#osps-do-03)**: The [project documentation] MUST provide user
guides for all basic functionality.

**[OSPS-DO-05](#osps-do-05)**: The [project documentation] MUST include a
mechanism for reporting [defects].

**[OSPS-DO-12](#osps-do-12)**: The [project documentation] MUST contain
instructions to verify the integrity 
and authenticity of the [release] assets,
including the expected identity of the person
or process authoring the software [release].

**[OSPS-DO-15](#osps-do-15)**: The [project documentation] MUST include a
description of how the project selects,
obtains, and tracks its dependencies.

**[OSPS-GV-01](#osps-gv-01)**: The [project documentation] MUST include the
Roles and Responsibilities for members of the
project.

**[OSPS-GV-04](#osps-gv-04)**: The [project documentation] MUST include a
guide for code [contributors] that includes
requirements for acceptable contributions.

**[OSPS-GV-05](#osps-gv-05)**: The [project documentation] MUST have a policy
that code [contributors] are reviewed prior to
granting escalated permissions to sensitive
resources.

**[OSPS-LE-01](#osps-le-01)**: The [version control system] MUST require all
code [contributors] to assert that they are
legally authorized to [commit] the associated
contributions on every [commit].

**[OSPS-QA-03](#osps-qa-03)**: All [released software assets] MUST be
delivered with a machine-readable list of
all direct and transitive internal software
dependencies with their associated version
identifiers.

**[OSPS-QA-04](#osps-qa-04)**: Any automated [status checks] for [commits] MUST
pass or require manual acknowledgement prior to
merge.

**[OSPS-QA-06](#osps-qa-06)**: The [version control system] MUST NOT contain
generated executable artifacts.

**[OSPS-SA-01](#osps-sa-01)**: The [project documentation] MUST provide
design documentation demonstrating all
actions and actors within the system.

**[OSPS-SA-04](#osps-sa-04)**: The project MUST perform a security
assessment to understand the most likely and
impactful potential security problems that
could occur within the software.

**[OSPS-VM-03](#osps-vm-03)**: The [project documentation] MUST include a
policy for coordinated vulnerability
reporting, with a clear timeframe for
response.


### Level 3

**[OSPS-AC-07](#osps-ac-07)**: The project's [version control system] MUST
require [multi-factor authentication] that
does not include SMS for users when
modifying the project [repository] settings or
accessing sensitive data.

**[OSPS-DO-14](#osps-do-14)**: The [project documentation] MUST provide a
descriptive statement when [releases] or
versions are no longer supported and that
will no longer receive security updates.

**[OSPS-QA-05](#osps-qa-05)**: Any additional [subproject] code [repositories]
produced by the project and compiled into a
[release] MUST enforce security requirements as
applicable to the status and intent of the
respective [codebase].

**[OSPS-QA-08](#osps-qa-08)**: The project MUST use at least one automated
test suite for the source code [repository]
which clearly documents when and how tests
are run.

**[OSPS-QA-09](#osps-qa-09)**: The [project documentation] MUST include a
policy that all major [changes] to the
software produced by the project should
add or update tests of the functionality
in an [automated test suite].

**[OSPS-QA-10](#osps-qa-10)**: The project's [version control system] MUST
require at least one non-author approval of
[changes] before merging into the [release] or
[primary branch].

**[OSPS-SA-03](#osps-sa-03)**: The project MUST perform a threat modeling and
attack surface analysis to understand and
protect against attacks on critical code
paths, functions, and interactions within
the system.

**[OSPS-VM-01](#osps-vm-01)**: The [project documentation] MUST include a
policy that defines a threshold for remediation
of [SCA] findings related to vulnerabilities and
[licenses].

**[OSPS-VM-02](#osps-vm-02)**: The [project documentation] MUST include a
policy to address [SCA] violations prior to any
[release].

**[OSPS-VM-04](#osps-vm-04)**: All proposed [changes] to the project's
[codebase] must be automatically evaluated 
against a documented policy for known
vulnerabilities and blocked in the
event of violations except when declared
and supressed as non exploitable.


## Categories



### Access Control

Access Control criteria focus on the mechanisms and
policies that control access to the project's version
control system and CI/CD pipelines. These criteria help
ensure that only authorized users can access sensitive
data, modify repository settings, or execute build and
release processes.


### OSPS-AC-01

**Criterion:**

The project's [version control system] MUST
require [multi-factor authentication] for
[collaborators] modifying the project
[repository] settings or accessing sensitive
data.

**Maturity Level:**
1

**Category:**


**Rationale:**

Protect against unauthorized access to
sensitive areas of the project's [repository],
reducing the risk of account compromise or
insider threats. Passwords are often easily captured,
either during communication or on servers,
so depending solely on passwords is a weaker
form of authentication.

**Details:**

Require [multi-factor authentication] ([MFA]) for the
project's [version control system], requiring
[collaborators] to provide a second form of
authentication when accessing sensitive data
or modifying [repository] settings.
Passkeys are acceptable for this criterion.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-AC-02

**Criterion:**

The project's [version control system] MUST
restrict [collaborator] permissions to the
lowest available privileges by default.

**Maturity Level:**
1

**Category:**


**Rationale:**

Reduce the risk of unauthorized access to
the project's [repository] by limiting the
permissions granted to [collaborators].

**Details:**

Configure the project's version control
system to assign the lowest available
permissions to [collaborators] by default when
added, granting additional permissions only
when necessary.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-AC-03

**Criterion:**

The project's [version control system] MUST
prevent unintentional direct [commits] against
the [primary branch].

**Maturity Level:**
1

**Category:**


**Rationale:**

Reduce the risk of accidental [changes] to the
[primary branch] of the project's [repository],
ensuring that due diligence is done before
[commits] are merged.

**Details:**

Set branch protection on the [primary branch]
in the project's [version control system]
requiring [changes] to be made through
pull/merge requests or other review
mechanisms.
Alternatively, use a decentralized approach
(like the Linux kernel's) where [changes] are
first proposed in another [repository], and
merging [changes] into the primary [repository]
requires a specific separate act.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-AC-04

**Criterion:**

The project's [version control system] MUST
prevent unintentional deletion of the
[primary branch].

**Maturity Level:**
1

**Category:**


**Rationale:**

Protect the [primary branch] of the project's
[repository] from accidental deletion,
ensuring that the project's history and
[codebase] are preserved.

**Details:**

Set branch protection on the [primary branch]
in the project's [version control system] to
prevent deletion.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-AC-05

**Criterion:**

The project's permissions in [CI/CD pipelines]
MUST be configured to the lowest available
privileges except when explicitly elevated.

**Maturity Level:**
2

**Category:**


**Rationale:**

Reduce the risk of unauthorized access to
the project's build and [release] processes by
limiting the permissions granted to steps
within the [CI/CD pipelines].

**Details:**

Configure the project's [CI/CD pipelines] to
assign the lowest available permissions to
users and services by default, elevating
permissions only when necessary for specific
tasks. In some [version control systems], this
may be possible at the organizational or
[repository] level. If not, set permissions at
the top level of the pipeline.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-AC-07

**Criterion:**

The project's [version control system] MUST
require [multi-factor authentication] that
does not include SMS for users when
modifying the project [repository] settings or
accessing sensitive data.

**Maturity Level:**
3

**Category:**


**Rationale:**

Ensure that [multi-factor authentication]
does not allow SMS as a factor, because
SMS lacks encryption and may be vulnerable
to attacks via Signaling System 7,
social engineering, or SIM-swapping.

**Details:**

Require [multi-factor authentication] methods
that do not include SMS for users. Common
alternatives include hardware tokens, mobile
authenticator apps, or biometric
authentication.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Build and Release

Build and Release criteria focus on the processes and
tools used to compile, package, and distribute the
project's software. These criteria help ensure that the
project's build and release pipelines are secure,
consistent, and reliable, reducing the risk of
vulnerabilities or errors in the software distribution
process.


### OSPS-BR-01

**Criterion:**

The project's [build and release pipelines]
MUST NOT execute [arbitrary code] that is
input from outside of the build script.

**Maturity Level:**
1

**Category:**


**Rationale:**

Reduce the risk of code injection or other
security vulnerabilities in the project's
build and [release] processes by restricting
the execution of external code.

**Details:**

Ensure that the project's build and [release]
pipelines do not execute [arbitrary code]
provided from external sources.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-BR-02

**Criterion:**

All [releases] and [released software assets]
MUST be assigned a unique version
identifier for each [release] intended to be
used by users.

**Maturity Level:**
2

**Category:**


**Rationale:**

Ensure that each software asset produced by
the  project is uniquely identified,
enabling users to track [changes] and updates
to the project over time.

**Details:**

Assign a unique [version identifier] to each
[release] and associated software asset
produced by the project, following a
consistent naming convention or numbering
scheme.
Examples include SemVer, CalVer, or
git [commit] id.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-BR-03

**Criterion:**

Any websites, API responses or other
services involved in the project development
and [release] MUST be delivered using SSH,
HTTPS or other encrypted channels.

**Maturity Level:**
1

**Category:**


**Rationale:**

Protect the confidentiality and integrity
of data transmitted between the project's
services and users, reducing the risk of
eavesdropping or data tampering.

**Details:**

Configure the project's websites, API
responses, and other services to use
encrypted channels such as SSH or HTTPS for
data transmission.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-BR-04

**Criterion:**

All [released software assets] MUST be created
with consistent, automated build and [release]
pipelines.

**Maturity Level:**
2

**Category:**


**Rationale:**

Ensure that the project's software assets
are built and released using consistent and
automated processes, reducing the risk of
errors or inconsistencies in the deployment
and distribution of the software.

**Details:**

[VCS]-integrated pipelines are
recommended to ensure consistency and
automation in the build and [release]
processes.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-BR-05

**Criterion:**

All [build and release pipelines] MUST use
standardized tooling where available to
ingest dependencies at build time.

**Maturity Level:**
2

**Category:**


**Rationale:**

Ensure that the project's build and [release]
pipelines use standardized tools and
processes to manage dependencies, reducing
the risk of compatibility issues or security
vulnerabilities in the software.

**Details:**

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

---

### OSPS-BR-06

**Criterion:**

All [releases] MUST provide a descriptive log
of functional and security modifications.

**Maturity Level:**
2

**Category:**


**Rationale:**

Provide transparency and accountability for
[changes] made to the project's software
[releases], enabling users to understand the
modifications and improvements included in
each [release].

**Details:**

Ensure that all [releases] include a
descriptive [change] log. 

It is recommended to ensure that the [change]
log is human-readable and includes details
beyond [commit] messages, such as descriptions 
of the security impact or relevance to
different use cases.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-BR-08

**Criterion:**

All [released software assets] MUST be signed
or accounted for in a signed manifest
including each asset's cryptographic hashes.

**Maturity Level:**
2

**Category:**


**Rationale:**

Provide users with a mechanism to verify the
authenticity and integrity of released 
software assets, reducing the risk of
tampering or unauthorized modifications.

**Details:**

Sign all [released software assets] at build
time with a cryptographic signature or
attestations, such  as GPG or PGP signature,
Sigstore signatures, SLSA [provenance], or SLSA
VSAs. Include the cryptographic hashes of
each asset in a signed manifest or
metadata file.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Documentation

Documentation criteria focus on the information
provided to users, contributors, and maintainers
of the project. These criteria help ensure that
the project's documentation is comprehensive,
accurate, and up-to-date, enabling users to
understand the project's features and functionality.


### OSPS-DO-03

**Criterion:**

The [project documentation] MUST provide user
guides for all basic functionality.

**Maturity Level:**
2

**Category:**
Documentation

**Rationale:**

Ensure that users have a clear and
comprehensive understanding of the project's
current features in order to prevent damage
from misuse or misconfiguration.

**Details:**

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

---

### OSPS-DO-05

**Criterion:**

The [project documentation] MUST include a
mechanism for reporting [defects].

**Maturity Level:**
2

**Category:**
Documentation

**Rationale:**

Enable users and [contributors] to report
[defects] or issues with the released software
assets, facilitating communication and
collaboration on [defect] fixes and
improvements.

**Details:**

It is recommended that projects use their
[VCS] default issue tracker. If an extarnal
source is used, ensure that the project
documentation and contributing guide clearly
and visibly explain how to use the reporting
system.

It is recommended that [project documentation]
also sets expectations for how [defects] will
be triaged and resolved.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-DO-12

**Criterion:**

The [project documentation] MUST contain
instructions to verify the integrity 
and authenticity of the [release] assets,
including the expected identity of the person
or process authoring the software [release].

**Maturity Level:**
2

**Category:**
Documentation

**Rationale:**

Enable users to verify the authenticity and
integrity of the project's released software
assets, reducing the risk of using tampered
or unauthorized versions of the software.

**Details:**

Instructions in the project should contain
information about the technology used, the
commands to run, and the expected output. The 
expected identity may be in the form of key
IDs used to sign, issuer and identity from a
sigstore certificate, or other similar forms.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-DO-13

**Criterion:**

The [project documentation] MUST include a
descriptive statement about the scope and
duration of support.

**Maturity Level:**
1

**Category:**
Documentation

**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-DO-14

**Criterion:**

The [project documentation] MUST provide a
descriptive statement when [releases] or
versions are no longer supported and that
will no longer receive security updates.

**Maturity Level:**
3

**Category:**
Documentation

**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-DO-15

**Criterion:**

The [project documentation] MUST include a
description of how the project selects,
obtains, and tracks its dependencies.

**Maturity Level:**
2

**Category:**
Vulnerability Management

**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Governance

Governance criteria focus on the policies and
procedures that guide the project's decision-making
and community interactions. These criteria help ensure
that the project is well positioned to respond to
both threats and opportunities.


### OSPS-GV-01

**Criterion:**

The [project documentation] MUST include the
Roles and Responsibilities for members of the
project.

**Maturity Level:**
2

**Category:**


**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-GV-02

**Criterion:**

The project MUST have one or more mechanisms
for public discussions about proposed
[changes] and usage obstacles.

**Maturity Level:**
1

**Category:**


**Rationale:**

Encourages open communication and
collaboration within the project community,
enabling users to provide feedback and
discuss proposed [changes] or usage
challenges.

**Details:**

Establish one or more mechanisms for public 
discussions within the project, such as
mailing lists, instant messaging, or issue
trackers, to facilitate open communication
and feedback.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-GV-03

**Criterion:**

The [project documentation] MUST include an
explanation of the contribution process.

**Maturity Level:**
1

**Category:**


**Rationale:**

Provide guidance to new [contributors] on how
to participate in the project, outlining the
steps required to submit [changes] or
enhancements to the project's [codebase].

**Details:**

Create a CONTRIBUTING.md or CONTRIBUTING/
directory to outline the contribution
process including the steps for submitting
[changes], and engaging with the project
maintainers.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-GV-04

**Criterion:**

The [project documentation] MUST include a
guide for code [contributors] that includes
requirements for acceptable contributions.

**Maturity Level:**
2

**Category:**


**Rationale:**

Provide guidance to code [contributors] on how
to submit [changes] and enhancements to the
project's [codebase], outlining the standards
and expectations for acceptable
contributions.

**Details:**

Extend the CONTRIBUTING.md or CONTRIBUTING/
contents in the [project documentation] to
outline the requirements for acceptable
contributions, including coding standards,
testing requirements, and submission
guidelines for code [contributors].

It is recommended that this guide is the
source of truth for both [contributors] and
approvers.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-GV-05

**Criterion:**

The [project documentation] MUST have a policy
that code [contributors] are reviewed prior to
granting escalated permissions to sensitive
resources.

**Maturity Level:**
2

**Category:**


**Rationale:**

Ensure that code [contributors] are vetted and
reviewed before being granted elevated
permissions to sensitive resources within
the project, reducing the risk of
unauthorized access or misuse.

**Details:**

Publish an enforceable policy in the project
documentation that requires code
[contributors] to be reviewed and approved
before being granted escalated permissions
to sensitive resources, such as merge
approval or access to secrets.

It is recommended that vetting includes
establishing a justifiable lineage of
identity such as confirming the
[contributor]'s association with a known
trusted organization.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Legal

Legal criteria focus on the policies and
procedures that govern the project's licensing
and intellectual property. These criteria help
ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software license,
reducing the risk of intellectual property
disputes or licensing violations.


### OSPS-LE-01

**Criterion:**

The [version control system] MUST require all
code [contributors] to assert that they are
legally authorized to [commit] the associated
contributions on every [commit].

**Maturity Level:**
2

**Category:**


**Rationale:**

Ensure that code [contributors] are aware of
and acknowledge their legal responsibility
for the contributions they make to the
project, reducing the risk of intellectual
property disputes against the project.

**Details:**

Include a DCO or CLA in the project's
[repository], requiring code [contributors] to
assert that they are legally authorized to
[commit] the associated contributions on every
[commit]. Use a [status check] to ensure the
assertion is made.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-LE-02

**Criterion:**

The [license] for the source code MUST
meet the OSI Open Source Definition or
the FSF Free Software Definition.

**Maturity Level:**
1

**Category:**


**Rationale:**

Ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software [license],
providing clarity on how the code
can be used and shared by others.

**Details:**

Add a [LICENSE] file to the project's [repo]
with a [license] that is
an approved [license] by the
Open Source Initiative (OSI), or
a free [license] as approved by the
Free Software Foundation (FSF).
Examples of such [licenses] include the
MIT, BSD 2-clause, BSD 3-clause revised,
Apache 2.0, Lesser GNU General Public
[License] (LGPL), and the
GNU General Public [License] (GPL).
Releasing to the public domain (e.g., CC0)
meets this criterion if there are no
other encumbrances (e.g., patents).

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-LE-03

**Criterion:**

The [license] for the source code MUST be
maintained in a standard location within
the project's [repository].

**Maturity Level:**
1

**Category:**


**Rationale:**

Ensure that the project's source code is
distributed with the appropriate [license]
terms, making it clear to users and
[contributors] how the code can be used and
shared.

**Details:**

Include the project's source code [license] in
the project's [LICENSE] file, COPYING file,
or [LICENSE]/
directory to provide visibility and clarity
on the licensing terms. The filename MAY
have an extension.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-LE-04

**Criterion:**

The [license] for the [released software assets]
MUST meet the OSI Open Source Definition or
the FSF Free Software Definition.

**Maturity Level:**
1

**Category:**


**Rationale:**

Ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software [license],
providing clarity on how the code
can be used and shared by others.

**Details:**

If a different [license] is included with
[released software assets], ensure it is
an approved [license] by the
Open Source Initiative (OSI), or
a free [license] as approved by the
Free Software Foundation (FSF).
Examples of such [licenses] include the
MIT, BSD 2-clause, BSD 3-clause revised,
Apache 2.0, Lesser GNU General Public
[License] (LGPL), and the
GNU General Public [License] (GPL).
Note that the [license] for the released
software assets may be different than the
source code.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Quality

Quality criteria focus on the processes and
practices used to ensure the quality and
reliability of the project's source code and
software assets. These criteria help ensure
that the project's source code is well
maintained, secure, and reliable, reducing the
risk of defects or vulnerabilities in the
software.


### OSPS-QA-01

**Criterion:**

The project's source code MUST be publicly
readable and have a static URL.

**Maturity Level:**
1

**Category:**


**Rationale:**

Enable users to access and review the
project's source code and history, promoting
transparency and collaboration within the
project community.

**Details:**

Use a common [VCS] such as GitHub, GitLab, or
Bitbucket. Ensure the [repository] is publicly
readable. Avoid duplication or mirroring of
[repositories] unless highly visible
documentation clarifies the primary source.
Avoid frequent [changes] to the [repository]
that would impact the [repository] URL.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-02

**Criterion:**

The [version control system] MUST contain a
publicly readable record of all [changes]
made, who made the [changes], and when the
[changes] were made.

**Maturity Level:**
1

**Category:**


**Rationale:**

Provide transparency and accountability for
[changes] made to the project's [codebase],
enabling users to track modifications and
understand the history of the project.

**Details:**

Use a common [VCS] such as GitHub, GitLab, or
Bitbucket to maintain a publicly readable
[commit] history. Avoid squashing or rewriting
[commits] in a way that would obscure the
author of any [commits].

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-03

**Criterion:**

All [released software assets] MUST be
delivered with a machine-readable list of
all direct and transitive internal software
dependencies with their associated version
identifiers.

**Maturity Level:**
2

**Category:**


**Rationale:**

Provide transparency and accountability for
the project's dependencies, enabling users
and [contributors] to understand the
software's dependencies and versions.

**Details:**

This may take the form of a software bill of
materials (SBOM) or a dependency file that
lists all direct and transitive dependencies
such as package.json, Gemfile.lock, or
go.sum.

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

---

### OSPS-QA-04

**Criterion:**

Any automated [status checks] for [commits] MUST
pass or require manual acknowledgement prior to
merge.

**Maturity Level:**
2

**Category:**


**Rationale:**

Ensure that the project's approvers do not
become accustomed to tolerating failing
[status checks], even if arbitrary, because it
increases the risk of overlooking security
vulnerabilities or [defects] identified by
automated checks.

**Details:**

Configure the project's version control
system to require that all automated status
checks pass or require manual acknowledgement
before a [commit] can be merged into the
[primary branch].

It is recommended that any optional
[status checks] are NOT configured as a pass
or fail requirement that approvers may be
tempted to bypass.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-05

**Criterion:**

Any additional [subproject] code [repositories]
produced by the project and compiled into a
[release] MUST enforce security requirements as
applicable to the status and intent of the
respective [codebase].

**Maturity Level:**
3

**Category:**


**Rationale:**

Ensure that additional code [repositories] or
[subprojects] produced by the project are held
to a standard that clear and appropriate
for that [codebase].

**Details:**

The parent project should maintain a list of
any [codebases] that are considered
[subprojects] or additional [repositories].
[Collaborators] on those [repositories] should
identify the proper maturity level and apply
the Open Source Project Security Baseline to
the [codebase]. Any [subproject] or [repository]
from the project which is compiled into the
primary project must be held to the same 
standard as the primary project. Others may
be held to a lower standard if they have
lower levels of adoption or are not intended
for general use.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-06

**Criterion:**

The [version control system] MUST NOT contain
generated executable artifacts.

**Maturity Level:**
2

**Category:**


**Rationale:**

Reduce the risk of including generated
executable artifacts in the project's
[version control system], ensuring that only
source code and necessary files are stored
in the [repository].

**Details:**

Remove generated executable artifacts
in the project's [version control system].

It is recommended that any scenario where a
generated executable artifact appears
critical to a process such as testing, it
should be instead be generated at build time
or stored separately and fetched during a
specific well-documented pipeline step.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-08

**Criterion:**

The project MUST use at least one automated
test suite for the source code [repository]
which clearly documents when and how tests
are run.

**Maturity Level:**
3

**Category:**


**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-09

**Criterion:**

The [project documentation] MUST include a
policy that all major [changes] to the
software produced by the project should
add or update tests of the functionality
in an [automated test suite].

**Maturity Level:**
3

**Category:**


**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-QA-10

**Criterion:**

The project's [version control system] MUST
require at least one non-author approval of
[changes] before merging into the [release] or
[primary branch].

**Maturity Level:**
3

**Category:**
Governance

**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Security Assessment

Security Assessment criteria encourage practices that
help ensure that the project is well positioned
to identify and address security vulnerabilities
and threats in the software.


### OSPS-SA-01

**Criterion:**

The [project documentation] MUST provide
design documentation demonstrating all
actions and actors within the system.

**Maturity Level:**
2

**Category:**


**Rationale:**

Provide an overview of the project's design
and architecture, illustrating the
interactions and components of the system to
help [contributors] and security reviewers
understand the internal logic of the
[released software assets].

**Details:**

Include designs in the [project documentation]
that explains the actions and actors. Actors
include any subsystem or entity that can
influence another segment in the system.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-SA-02

**Criterion:**

The [project documentation] MUST include
descriptions of all external input and output
interfaces of the [released software assets].

**Maturity Level:**
1

**Category:**


**Rationale:**

Provide users and developers with an
understanding of how to interact with the
project's software and integrate it with
other systems, enabling them to use the
software effectively.

**Details:**

Document all input and output interfaces of
the [released software assets], explaining how
users can interact with the software and
what data is expected or produced.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-SA-04

**Criterion:**

The project MUST perform a security
assessment to understand the most likely and
impactful potential security problems that
could occur within the software.

**Maturity Level:**
2

**Category:**


**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-SA-03

**Criterion:**

The project MUST perform a threat modeling and
attack surface analysis to understand and
protect against attacks on critical code
paths, functions, and interactions within
the system.

**Maturity Level:**
3

**Category:**


**Rationale:**


**Details:**


**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### Vulnerability Management

Vulnerability Management criteria focus on the
processes and practices used to identify and
address security vulnerabilities in the project's
software dependencies. These criteria help ensure
that the project is well positioned to respond to
security threats and vulnerabilities in the software.


### OSPS-VM-01

**Criterion:**

The [project documentation] MUST include a
policy that defines a threshold for remediation
of [SCA] findings related to vulnerabilities and
[licenses].

**Maturity Level:**
3

**Category:**


**Rationale:**

Ensure that the project clearly communicates
the threshold for remediation of [SCA] findings,
including vulnerabilities and [license] issues
in software dependencies.

**Details:**

Document a policy in the project that
defines a threshold for remediation of [SCA]
findings related to vulnerabilities and
[licenses]. Include the process for
identifying, prioritizing, and remediating
these findings.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-VM-02

**Criterion:**

The [project documentation] MUST include a
policy to address [SCA] violations prior to any
[release].

**Maturity Level:**
3

**Category:**


**Rationale:**

Ensure that violations of your [SCA] policy
are addressed before software [releases],
reducing the risk of shipping insecure or
non-compliant software.

**Details:**

Document a policy in the project to address
applicable [Software Composition Analysis] 
results before any [release], and add status
checks that verify compliance with that 
policy prior to [release].

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-VM-03

**Criterion:**

The [project documentation] MUST include a
policy for coordinated vulnerability
reporting, with a clear timeframe for
response.

**Maturity Level:**
2

**Category:**


**Rationale:**

Establish a process for reporting and
addressing vulnerabilities in the project,
ensuring that security issues are handled
promptly and transparently.

**Details:**

Create a SECURITY.md file at the root of the
directory, outlining the project's policy
for coordinated [vulnerability reporting].
Include a method for reporting
vulnerabilities. Set expectations for the
how the project will respond and address
reported issues.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---

### OSPS-VM-04

**Criterion:**

All proposed [changes] to the project's
[codebase] must be automatically evaluated 
against a documented policy for known
vulnerabilities and blocked in the
event of violations except when declared
and supressed as non exploitable.

**Maturity Level:**
3

**Category:**


**Rationale:**

Identify and address [defects] and security weaknesses 
in the project's [codebase] early in the 
development process, reducing the risk of 
shipping insecure software.

**Details:**

Create a [status check] in the project's
[version control system] that runs a Static
Application Security Testing (SAST) tool on
all [changes] to the [codebase]. Require that the
[status check] passes before [changes] can be
merged.

**Control Mappings:**

_No control mappings identified._

**Security Insights Value:**

_No security insights identified._

---


## Lexicon

### Arbitrary Code

Code provided by an external source that is
executed by a system without validation or
restriction.

### Automated Test Suite

A collection of pre-written test cases that, when invoked,
execute the software to verify that actual results are expected results
without requiring manual intervention.
An automated test suite must return an overall "pass" or "fail" result,
and is often implemented using a test framework.
Common ways to invoke automated tests include `make check`, `make test`, `npm test`, and `cargo test` manually or as part of a Continuous Integration workflow.

### Build and Release Pipeline

A series of automated processes that compile
and deploy software. Similar to the generic
term CI/CD Pipelines, but this term excludes
some pipelines, such as pre-merge status
checks.

### Change

Any alteration of the project's codebase,
CI/CD Pipelines, or documentation. This may
include addition, deletion, or modification
of content.

### CI/CD Pipeline

Automated pipelines for Continuous
Integration and Continuous Delivery.
Responsible for building, testing, and
delivering changes. These pipelines integrate
contributions frequently, enabling rapid and
reliable software delivery. CI focuses on
testing and building code, while CD delivers
software to location such as a package
registry.

In the context of the Open Source Project
Security Baseline, CD refers only to
continuous delivery, not to continuous
deployment, as sometimes used elsewhere.

### Contributor

Entities who commit code or documentation to
the project. Code contributors include
collaborators or external participants who
submit changes.

In the context of the Open Source Project
Security Baseline, code contributors does not
address non-code contributions such as
designing, triaging, reviewing, or testing.

### Codebase

The collection of source code and related
assets that make up the project. The codebase
includes all files necessary to build and
test the software. Lives in the repository,
sometimes alongside documentation and CI/CD
pipelines. The contents of the codebase are
the primary deliverable in a release.

### Collaborator

A user with a role on the project's version
control system who can approve changes or
manage the repository settings. Collaborators
may have varying permission levels based on
their role in the project. This does not
include contributors whose changes only
originate through a request from a repository
fork.

### Commit

A record of a single change submitted to the
version control system. Each commit includes
details such as the modifications made, the
contributor who made them, and the timestamp
of the change.

### Defect

Errors or flaws in the software that cause it
to produce incorrect or unintended results, 
or to behave in an unintended way. Defects
can include bugs, vulnerabilities, or other
issues that impact the software's
functionality or security. Defects may have
originally been intentional, but a change in
environment or understanding has made them
undesirable.

### Exploitable Vulnerabilities

Defects in the software that can be leveraged
by attackers to gain unauthorized access,
execute arbitrary code, or cause other
undesired outcomes.

### License

A legal document that defines the terms under
which the software can be used, modified, and
distributed. May be stored at the top level
of the repository in a file named `LICENSE`
or within files in a directory named
`LICENSE/`. The license applies to repository
contents and any released software assets,
unless otherwise stated.

### Known Vulnerabilities

Publicly acknowledged exploitable
vulnerabilities that have been identified
within the software. These vulnerabilities
often have associated advisories, patches, or
recommended mitigations.

All proposed changes to the project's
codebase must be automatically evaluated 
against a documented policy for known
vulnerabilities and blocked in the
event of violations.

### Multi-factor Authentication

An authentication method that requires two or
more verification factors (e.g., a password
and a token) to gain access to a resource.
This method strengthens security by requiring
multiple forms of identification.

### Primary Branch

The main development branch in the version
control system, representing the latest
stable codebase. Releases are typically made
from this branch. Commonly named `main` or
`master`. In some situations where branches
are not featured, a repository with forked
repositories will have the original repo
acting as an equivalent to the primary
branch.

### Project Documentation

Written materials related to the project,
such as user guides, developer guides, and
contribution guidelines. These documents help
users and developers understand, contribute
to, and interact with the software. At
release time, this may include provenance
information, licensing details, and other
metadata.

### Software Provenance

Information about the origin and history of
the released software assets. This may
include details about its development,
dependencies, vulnerabilities, contributors,
and licensing.

### Release

- _(verb)_ The process of making a version
controlled bundle of assets available to
users, such as through a package registry.
- _(noun)_ A version-controlled bundle of
code, documentation, and other assets made
available to users. A release often includes
release notes that describe the changes.

### Released Software Asset

Deliverables provided to users as part of a
release. These assets can include binaries,
libraries, or containers.

### Repository

A storage location managed by a version
control system where the project's code,
documentation, and other resources are
stored. It tracks changes, manages
collaborator permissions, and includes
configuration options such as branch
protection and access controls.

### Software Composition Analysis

The process of identifying and cataloging all
components and dependencies in a software
codebase. SCA is essential for managing
security vulnerabilities and ensuring
compliance with organizational policies.

### Status Check

Automated tests or validations that run on
commits before they are merged. Status checks
ensure that any changes meet the project's
quality and security standards.

### Subproject

A codebase that is part of the project but
maintained in a separate repository.
Subprojects may be compiled into the primary
project or used as standalone components.

### Version Identifier

A label assigned to a specific release of the
software, such as `v1.2.3`. Commonly
recommended formats are Semantic Versioning
or Calendar Versioning.

### Version Control System

A tool that tracks changes to files over time
and facilitates collaboration among
contributors. Examples of version control
systems include Git, Subversion, and
Mercurial.

### Vulnerability Reporting

The act of identifying and documenting
exploitable vulnerabilities in released
software assets. This may include privately
or openly reporting vulnerabilities to
maintainers, security teams, or the public,
as well as tracking the resolution of these
vulnerabilities.

---

## Acknowledgments

This baseline was created by community leaders from across the Linux Foundation, including:

- OpenJS Foundation (OpenJS)
- Open Source Security Foundation (OpenSSF)
- Cloud Native Computing Foundation (CNCF)
- Fintech Open Source Foundation (FINOS)
- [OSPS Baseline contributors](https://github.com/ossf/security-baseline/graphs/contributors)

[Arbitrary Code]: #arbitrary-code
[Automated Test Suite]: #automated-test-suite
[Build and Release Pipeline]: #build-and-release-pipeline
[build and release pipelines]: #build-and-release-pipeline
[Change]: #change
[changes]: #change
[CI/CD Pipeline]: #ci/cd-pipeline
[CI/CD pipelines]: #ci/cd-pipeline
[Contributor]: #contributor
[contributors]: #contributor
[Codebase]: #codebase
[codebases]: #codebase
[Collaborator]: #collaborator
[collaborators]: #collaborator
[Commit]: #commit
[commits]: #commit
[Defect]: #defect
[defects]: #defect
[Exploitable Vulnerabilities]: #exploitable-vulnerabilities
[Exploitable Vulnerability]: #exploitable-vulnerabilities
[License]: #license
[licenses]: #license
[Known Vulnerabilities]: #known-vulnerabilities
[Known Vulnerability]: #known-vulnerabilities
[Multi-factor Authentication]: #multi-factor-authentication
[MFA]: #multi-factor-authentication
[Primary Branch]: #primary-branch
[Project Documentation]: #project-documentation
[Software Provenance]: #software-provenance
[Provenance]: #software-provenance
[Release]: #release
[releases]: #release
[Released Software Asset]: #released-software-asset
[released software assets]: #released-software-asset
[Repository]: #repository
[Repo]: #repository
[Repositories]: #repository
[Software Composition Analysis]: #software-composition-analysis
[SCA]: #software-composition-analysis
[Status Check]: #status-check
[status checks]: #status-check
[Subproject]: #subproject
[subprojects]: #subproject
[Version Identifier]: #version-identifier
[Version Control System]: #version-control-system
[VCS]: #version-control-system
[version control systems]: #version-control-system
[Vulnerability Reporting]: #vulnerability-reporting
[Coordinated Vulnerability Reporting]: #vulnerability-reporting
