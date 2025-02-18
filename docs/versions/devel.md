# Open Source Project Security Baseline

Version: devel (not for production use)

<!-- A button for returning to the top of the page -->
<button onclick="toTop()" id="topButton" title="Go to top"
style="display: none; position: fixed; bottom: 20px; right: 30px; border: none; background-color: CornflowerBlue; color: white; cursor: pointer; padding: 10px; border-radius: 10px; font-size: 18px;">to top</button> 

<script>
let topButton = document.getElementById("topButton");
window.onscroll = function() {scrollFunction()};

function scrollFunction() {
  if (document.documentElement.scrollTop > 50 ) {
    topButton.style.display = "block";
  } else {
    topButton.style.display = "none";
  }
}

function toTop() {
  document.documentElement.scrollTop = 0;
}
</script>
<!-- That's enough button stuff for now -->

* Contents
{:toc}

## Overview

The Open Source Project Security (OSPS) Baseline is a set of security criteria that projects should meet to demonstrate a strong security posture.
The criteria are organized by maturity level and category.
In the detailed subsections you will find the criterion, rationale, and details notes.


Where possible, we have added control mappings to external frameworks.
These are not guaranteed to be 100% matches, but instead serve as references
when working to meet the corresponding controls.

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
MUST NOT permit untrusted input that allows
access to privileged resources.

**[OSPS-BR-03](#osps-br-03)**: Any websites and [version control systems]
involved in the project development
MUST be delivered using SSH,
HTTPS, or other encrypted channels.

**[OSPS-BR-09](#osps-br-09)**: Any websites or other services involved in the
distribution of [released software assets] MUST
be delivered using HTTPS or other encrypted
channels.

**[OSPS-DO-03](#osps-do-03)**: The [project documentation] MUST provide user
guides for all basic functionality.

**[OSPS-DO-05](#osps-do-05)**: The [project documentation] MUST include a
mechanism for reporting [defects].

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

**[OSPS-VM-05](#osps-vm-05)**: The project publishes contacts and process
for reporting vulnerabilities.


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

**[OSPS-BR-10](#osps-br-10)**: Any websites, API responses or other
services involved in [release] pipelines MUST be
fetched using SSH, HTTPS or other encrypted
channels.

**[OSPS-DO-12](#osps-do-12)**: The [project documentation] MUST contain
instructions to verify the integrity 
and authenticity of the [release] assets,
including the expected identity of the person
or process authoring the software [release].

**[OSPS-DO-13](#osps-do-13)**: The [project documentation] MUST include a
descriptive statement about the scope and
duration of support.

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

**[OSPS-SA-02](#osps-sa-02)**: The [project documentation] MUST include
descriptions of all external software
interfaces of the [released software assets].

**[OSPS-SA-04](#osps-sa-04)**: The project MUST perform a security
assessment to understand the most likely and
impactful potential security problems that
could occur within the software.

**[OSPS-VM-03](#osps-vm-03)**: The [project documentation] MUST include a
policy for coordinated vulnerability
reporting, with a clear timeframe for
response.

**[OSPS-VM-06](#osps-vm-06)**: The project MUST provide a means for reporting 
security vulnerabilities privately to the security
contacts within the project.

**[OSPS-VM-07](#osps-vm-07)**: The project MUST publicly publish data about vulnerabilities
discovered within the [codebase].


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

**[OSPS-SA-03](#osps-sa-03)**: The project MUST perform a [threat modeling] and
[attack surface analysis] to understand and
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




## Access Control

Access Control criteria focus on the mechanisms and
policies that control access to the project's version
control system and CI/CD pipelines. These criteria help
ensure that only authorized users can access sensitive
data, modify repository settings, or execute build and
release processes.


### OSPS-AC-01

**Criterion:** The project's [version control system] MUST
require [multi-factor authentication] for
[collaborators] modifying the project
[repository] settings or accessing sensitive
data.


**Maturity Level:** 1

**Rationale:** Protect against unauthorized access to
sensitive areas of the project's [repository],
reducing the risk of account compromise or
insider threats. Passwords are often easily captured,
either during communication or on servers,
so depending solely on passwords is a weaker
form of authentication.


**Details:** Require [multi-factor authentication] ([MFA]) for the
project's [version control system], requiring
[collaborators] to provide a second form of
authentication when accessing sensitive data
or modifying [repository] settings.
Passkeys are acceptable for this criterion.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-G-1 |
| [CRA] | 1.2d, 1.2e, 1.2f |
| CSF | PR.AA-02 |
| OCRE | 486-813, 124-564, 347-352, 333-858, 152-725, 201-246 |
| SSDF | PO3.2, PS1 |



---

### OSPS-AC-02

**Criterion:** The project's [version control system] MUST
restrict [collaborator] permissions to the
lowest available privileges by default.


**Maturity Level:** 1

**Rationale:** Reduce the risk of unauthorized access to
the project's [repository] by limiting the
permissions granted to [collaborators].


**Details:** Most public [version control systems] are configured
in this manner. Ensure the project's version control
system always assigns the lowest available
permissions to [collaborators] by default when
added, granting additional permissions only
when necessary.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2f |
| CSF | PR:AA-02 |
| OCRE | 486-813, 124-564, 802-056, 368-633, 152-725 |
| SSDF | PO3.2, PS1 |



---

### OSPS-AC-03

**Criterion:** The project's [version control system] MUST
prevent unintentional direct [commits] against
the [primary branch].


**Maturity Level:** 1

**Rationale:** Reduce the risk of accidental [changes] to the
[primary branch] of the project's [repository],
ensuring that due diligence is done before
[commits] are merged.

**Implementation:** If the [VCS] is centralized,
set branch protection on the [primary branch]
in the project's [VCS]

**Details:** Set branch protection on the [primary branch]
in the project's [version control system]
requiring [changes] to be made through
pull/merge requests or other review
mechanisms.
Alternatively, use a decentralized approach
(like the Linux kernel's) where [changes] are
first proposed in another [repository], and
merging [changes] into the primary [repository]
requires a specific separate act.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2f |
| CSF | PR.AA-02 |
| OCRE | 486-813, 124-564, 152-725 |
| SSDF | PO3.2, PS1 |



---

### OSPS-AC-04

**Criterion:** The project's [version control system] MUST
prevent unintentional deletion of the
[primary branch].


**Maturity Level:** 1

**Rationale:** Protect the [primary branch] of the project's
[repository] from accidental deletion,
ensuring that the project's history and
[codebase] are preserved.


**Details:** Set branch protection on the [primary branch]
in the project's [version control system] to
prevent deletion.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2b, 1.2f |
| CSF | PR.AA-02 |
| OCRE | 486-813, 124-564,123-124, 152-725 |
| SSDF | PO3.2, PS1 |



---

### OSPS-AC-05

**Criterion:** The project's permissions in [CI/CD pipelines]
MUST be configured to the lowest available
privileges except when explicitly elevated.


**Maturity Level:** 2

**Rationale:** Reduce the risk of unauthorized access to
the project's build and [release] processes by
limiting the permissions granted to steps
within the [CI/CD pipelines].


**Details:** Configure the project's [CI/CD pipelines] to
assign the lowest available permissions to
users and services by default, elevating
permissions only when necessary for specific
tasks. In some [version control systems], this
may be possible at the organizational or
[repository] level. If not, set permissions at
the top level of the pipeline.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2d, 1.2e, 1.2f |
| CSF | PR.AA-02, PR.AA-05 |
| OCRE | 486-813, 124-564,347-507, 263-284, 123-124 |
| SSDF | PO2, PO3.2, PS1 |



---

### OSPS-AC-07

**Criterion:** The project's [version control system] MUST
require [multi-factor authentication] that
does not include SMS for users when
modifying the project [repository] settings or
accessing sensitive data.


**Maturity Level:** 3

**Rationale:** Ensure that [multi-factor authentication]
does not allow SMS as a factor, because
SMS lacks encryption and may be vulnerable
to attacks via Signaling System 7,
social engineering, or SIM-swapping.


**Details:** Require [multi-factor authentication] methods
that do not include SMS for users. Common
alternatives include hardware tokens, mobile
authenticator apps, or biometric
authentication.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-G-1 |
| [CRA] | 1.2d |
| CSF | PR.AA-02 |
| OCRE | 486-813, 124-564,333-858, 102-811, 354-752 |
| SSDF | PO3.2, PS1 |



---

## Build and Release

Build and Release criteria focus on the processes and
tools used to compile, package, and distribute the
project's software. These criteria help ensure that the
project's build and release pipelines are secure,
consistent, and reliable, reducing the risk of
vulnerabilities or errors in the software distribution
process.


### OSPS-BR-01

**Criterion:** The project's [build and release pipelines]
MUST NOT permit untrusted input that allows
access to privileged resources.


**Maturity Level:** 1

**Rationale:** Reduce the risk of code injection or other
security vulnerabilities in the project's
build and [release] by preventing untrusted input
to access privileged resources
(secret exfiltration, final [release], etc.)


**Details:** Ensure that any integration or [release] pipeline actions
that accept externally-controlled untrusted input (e.g. git
branch names) do not use input in ways that could
provide unintended access to privileged resources.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2f |
| CSF | PR.AA-02 |
| OCRE | 483-813, 124-564, 357-352 |
| SSDF | PO3.2, PS1 |



---

### OSPS-BR-02

**Criterion:** All [releases] and [released software assets]
MUST be assigned a unique version
identifier for each [release] intended to be
used by users.


**Maturity Level:** 2

**Rationale:** Ensure that each software asset produced by
the  project is uniquely identified,
enabling users to track [changes] and updates
to the project over time.


**Details:** Assign a unique [version identifier] to each
[release] and associated software asset
produced by the project, following a
consistent naming convention or numbering
scheme.
Examples include SemVer, CalVer, or
git [commit] id.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-B-5, CC-B-6, CC-B-7 |
| [CRA] | 1.2f |
| OCRE | 483-813, 124-564 |
| SSDF | PO3.2, PS1, PS2, PS3 |



---

### OSPS-BR-03

**Criterion:** Any websites and [version control systems]
involved in the project development
MUST be delivered using SSH,
HTTPS, or other encrypted channels.


**Maturity Level:** 1

**Rationale:** Protect the confidentiality and integrity
of project source code during development,
reducing the risk of eavesdropping or data
tampering.


**Details:** Configure the project's websites and version
control systems to use encrypted channels
such as SSH or HTTPS for data transmission.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-11 |
| [CRA] | 1.2d, 1.2e, 1.2f, 1.2i, 1.2j, 1.2k |
| OCRE | 483-813, 124-564, 263-184 |
| SSDF | PO3.2, PS1 |



---

### OSPS-BR-04

**Criterion:** All [released software assets] MUST be created
with consistent, automated build and [release]
pipelines.


**Maturity Level:** 2

**Rationale:** Ensure that the project's software assets
are built and released using consistent and
automated processes, reducing the risk of
errors or inconsistencies in the deployment
and distribution of the software.


**Details:** [VCS]-integrated pipelines are
recommended to ensure consistency and
automation in the build and [release]
processes.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | Q-B-7 |
| [CRA] | 1.2b, 1.2d, 1.2f, 1.2h, 1.2j |
| OCRE | 483-813, 124-564, 347-352, 263-184, 208-355 |
| SSDF | PO3.2, PS1 |


**Security Insights Value:** project-lifecycle.release-process

---

### OSPS-BR-05

**Criterion:** All [build and release pipelines] MUST use
standardized tooling where available to
ingest dependencies at build time.


**Maturity Level:** 2

**Rationale:** Ensure that the project's build and [release]
pipelines use standardized tools and
processes to manage dependencies, reducing
the risk of compatibility issues or security
vulnerabilities in the software.


**Details:** Use a common tooling for your ecosystem,
such as package managers or dependency
management tools to ingest dependencies at
build time. This may include using a
dependency file, lock file, or manifest to
specify the required dependencies, which are
then pulled in by the build system.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | Q-B-2 |
| [CRA] | 1.2b, 1.2d, 1.2f, 1.2h, 1.2j, 2.1 |
| OCRE | 483-813, 124-564, 347-352, 715-334 |
| SSDF | PO3.2, PS1 |



---

### OSPS-BR-06

**Criterion:** All [releases] MUST provide a descriptive log
of functional and security modifications.


**Maturity Level:** 2

**Rationale:** Provide transparency and accountability for
[changes] made to the project's software
[releases], enabling users to understand the
modifications and improvements included in
each [release].


**Details:** Ensure that all [releases] include a
descriptive [change] log. 

It is recommended to ensure that the [change]
log is human-readable and includes details
beyond [commit] messages, such as descriptions 
of the security impact or relevance to
different use cases.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-B-8, CC-B-9 |
| [CRA] | 1.2l, 2.2 |
| OCRE | 483-813, 124-564, 745-356 |
| SSDF | PS1, PS2, PS3, PW1.2 |



---

### OSPS-BR-08

**Criterion:** All [released software assets] MUST be signed
or accounted for in a signed manifest
including each asset's cryptographic hashes.


**Maturity Level:** 2

**Rationale:** Provide users with a mechanism to verify the
authenticity and integrity of released 
software assets, reducing the risk of
tampering or unauthorized modifications.


**Details:** Sign all [released software assets] at build
time with a cryptographic signature or
attestations, such  as GPG or PGP signature,
Sigstore signatures, SLSA [provenance], or SLSA
VSAs. Include the cryptographic hashes of
each asset in a signed manifest or
metadata file.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| SSDF | PO5.2, PS2.1, PW6.2 |


**Security Insights Value:** Signed-Releases

---

### OSPS-BR-09

**Criterion:** Any websites or other services involved in the
distribution of [released software assets] MUST
be delivered using HTTPS or other encrypted
channels.


**Maturity Level:** 1

**Rationale:** Protect the confidentiality and integrity
of [release] assets consumed by the project's
users, reducing the risk of eavesdropping or
data tampering.


**Details:** Configure the project's websites and
distribution services to use encrypted channels
such as HTTPS for data transmission.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-11 |
| [CRA] | 1.2d, 1.2e, 1.2f, 1.2i, 1.2j, 1.2k |
| OCRE | 483-813, 124-564, 263-184 |
| SSDF | PO3.2, PS1 |



---

### OSPS-BR-10

**Criterion:** Any websites, API responses or other
services involved in [release] pipelines MUST be
fetched using SSH, HTTPS or other encrypted
channels.


**Maturity Level:** 2

**Rationale:** Protect the confidentiality and integrity
of assets used in the [release] pipeline,
reducing the risk of eavesdropping or data
tampering.


**Details:** Configure the project's [release] pipeline to
only fetch data from websites, API
responses, and other services which use
encrypted channels such as SSH or HTTPS for
data transmission.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-11 |
| [CRA] | 1.2d, 1.2e, 1.2f, 1.2i, 1.2j, 1.2k |
| OCRE | 483-813, 124-564, 263-184 |
| SSDF | PO3.2, PS1 |



---

## Documentation

Documentation criteria focus on the information
provided to users, contributors, and maintainers
of the project. These criteria help ensure that
the project's documentation is comprehensive,
accurate, and up-to-date, enabling users to
understand the project's features and functionality.


### OSPS-DO-03

**Criterion:** The [project documentation] MUST provide user
guides for all basic functionality.


**Maturity Level:** 1

**Rationale:** Ensure that users have a clear and
comprehensive understanding of the project's
current features in order to prevent damage
from misuse or misconfiguration.


**Details:** Create user guides or documentation for all
basic functionality of the project,
explaining how to install, configure, and
use the project's features. If there are any
known dangerous or destructive actions
available, include highly-visible warnings.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-1, B-B-9, B-S-7, B-S-9 |
| [CRA] | 1.2b, 1.2j, 1.2k |
| CSF | GV.OC-04, GV.OC-05 |
| OC | 4.1.4 |
| OCRE | 036-275 |
| SSDF | PW1.2 |



---

### OSPS-DO-05

**Criterion:** The [project documentation] MUST include a
mechanism for reporting [defects].


**Maturity Level:** 1

**Rationale:** Enable users and [contributors] to report
[defects] or issues with the released software
assets, facilitating communication and
collaboration on [defect] fixes and
improvements.


**Details:** It is recommended that projects use their
[VCS] default issue tracker. If an extarnal
source is used, ensure that the project
documentation and contributing guide clearly
and visibly explain how to use the reporting
system.

It is recommended that [project documentation]
also sets expectations for how [defects] will
be triaged and resolved.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-3, R-B-1+, R-B-1, R-B-2, R-S-2 |
| [CRA] | 1.2c, 1.2l, 2.1, 2.2,2.5, 2.6 |
| CSF | RS.MA-02, GV.RM-05 |
| OC | 4.2.1 |
| SSDF | PW1.2, RV1.1, RV2.1, RV1.2 |



---

### OSPS-DO-12

**Criterion:** The [project documentation] MUST contain
instructions to verify the integrity 
and authenticity of the [release] assets,
including the expected identity of the person
or process authoring the software [release].


**Maturity Level:** 2

**Rationale:** Enable users to verify the authenticity and
integrity of the project's released software
assets, reducing the risk of using tampered
or unauthorized versions of the software.


**Details:** Instructions in the project should contain
information about the technology used, the
commands to run, and the expected output. The 
expected identity may be in the form of key
IDs used to sign, issuer and identity from a
sigstore certificate, or other similar forms.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-B-8 |
| [CRA] | 1.2d |
| OCRE | 171-222 |
| SSDF | PO4.2, PS.2, PS2.1, PS3.1, RV1.3 |



---

### OSPS-DO-13

**Criterion:** The [project documentation] MUST include a
descriptive statement about the scope and
duration of support.


**Maturity Level:** 2

**Rationale:** 
**Implementation:** The project should have either a "Support"
header in the README, a SUPPORT.md file
present in the [repo] root, or a SUPPORT.eox
file in the [OpenEOX format](https://github.com/OpenEoX/openeox/blob/main/schema/schema.json)
describing the scope and duration of support
for the project's [released software assets].

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | R-B-3 |
| OC | 4.1, 4.3.1 |
| SSDF | PO4.2, PS3.1, RV1.3 |



---

### OSPS-DO-14

**Criterion:** The [project documentation] MUST provide a
descriptive statement when [releases] or
versions are no longer supported and that
will no longer receive security updates.


**Maturity Level:** 3

**Rationale:** 

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2c, 2.6 |
| OC | 4.1.1, 4.3.1 |
| OCRE | 673-475, 053-751 |



---

### OSPS-DO-15

**Criterion:** The [project documentation] MUST include a
description of how the project selects,
obtains, and tracks its dependencies.


**Maturity Level:** 2

**Rationale:** 

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | A-S-1 |
| [CRA] | 2.1 |
| OCRE | 613-286, 053-751 |


**Security Insights Value:** Pinned-Dependencies

---

## Governance

Governance criteria focus on the policies and
procedures that guide the project's decision-making
and community interactions. These criteria help ensure
that the project is well positioned to respond to
both threats and opportunities.


### OSPS-GV-01

**Criterion:** The [project documentation] MUST include the
Roles and Responsibilities for members of the
project.


**Maturity Level:** 2

**Rationale:** Documenting project roles and responsibilities
helps project particpants, potential [contributors],
and downstream consumers have an accurate
understand of who is working on the project
and what areas of authority they may have.

**Implementation:** Document project participants and their roles
through such artifacts as members.md, governance.md, 
maintainers.md, or similar file within the source 
code [repository] of the project.

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-S-3, B-S-4 |
| OCRE | 013-021 |



---

### OSPS-GV-02

**Criterion:** The project MUST have one or more mechanisms
for public discussions about proposed
[changes] and usage obstacles.


**Maturity Level:** 1

**Rationale:** Encourages open communication and
collaboration within the project community,
enabling users to provide feedback and
discuss proposed [changes] or usage
challenges.


**Details:** Establish one or more mechanisms for public 
discussions within the project, such as
mailing lists, instant messaging, or issue
trackers, to facilitate open communication
and feedback.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-3, B-B-12 |
| [CRA] | 1.2l, 2.3, 2.4, 2.6 |
| CSF |  |
| OC |  |
| OCRE |  |
| SSDF | PS3, PW1.2 |



---

### OSPS-GV-03

**Criterion:** The [project documentation] MUST include an
explanation of the contribution process.


**Maturity Level:** 1

**Rationale:** Provide guidance to new [contributors] on how
to participate in the project, outlining the
steps required to submit [changes] or
enhancements to the project's [codebase].


**Details:** Create a CONTRIBUTING.md or CONTRIBUTING/
directory to outline the contribution
process including the steps for submitting
[changes], and engaging with the project
maintainers.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-4, B-S-3, B-B-4+, R-B-1, Q-G-2 |
| [CRA] | 1.2l, 2.4 |
| SSDF | PW1.2 |



---

### OSPS-GV-04

**Criterion:** The [project documentation] MUST include a
guide for code [contributors] that includes
requirements for acceptable contributions.


**Maturity Level:** 2

**Rationale:** Provide guidance to code [contributors] on how
to submit [changes] and enhancements to the
project's [codebase], outlining the standards
and expectations for acceptable
contributions.


**Details:** Extend the CONTRIBUTING.md or CONTRIBUTING/
contents in the [project documentation] to
outline the requirements for acceptable
contributions, including coding standards,
testing requirements, and submission
guidelines for code [contributors].

It is recommended that this guide is the
source of truth for both [contributors] and
approvers.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-5, B-S-3, B-B-4+, Q-G-2 |
| [CRA] | 1.2l, 2.1, 2.2, 2.5, 2.6 |
| OC | 4.1.2 |



---

### OSPS-GV-05

**Criterion:** The [project documentation] MUST have a policy
that code [contributors] are reviewed prior to
granting escalated permissions to sensitive
resources.


**Maturity Level:** 2

**Rationale:** Ensure that code [contributors] are vetted and
reviewed before being granted elevated
permissions to sensitive resources within
the project, reducing the risk of
unauthorized access or misuse.


**Details:** Publish an enforceable policy in the project
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


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2d |
| CSF | PR.AA-02, PR.AA-05 |
| OCRE | 123-124, 152-725 |
| SSDF | PO2, PO3.2 |



---

## Legal

Legal criteria focus on the policies and
procedures that govern the project's licensing
and intellectual property. These criteria help
ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software license,
reducing the risk of intellectual property
disputes or licensing violations.


### OSPS-LE-01

**Criterion:** The [version control system] MUST require all
code [contributors] to assert that they are
legally authorized to [commit] the associated
contributions on every [commit].


**Maturity Level:** 2

**Rationale:** Ensure that code [contributors] are aware of
and acknowledge their legal responsibility
for the contributions they make to the
project, reducing the risk of intellectual
property disputes against the project.


**Details:** Include a DCO or CLA in the project's
[repository], requiring code [contributors] to
assert that they are legally authorized to
[commit] the associated contributions on every
[commit]. Use a [status check] to ensure the
assertion is made.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-S-1 |
| [CRA] | 1.2b, 1.2f |
| SSDF | PO3.2, PS1, PW1.2, PW2.1 |



---

### OSPS-LE-02

**Criterion:** The [license] for the source code MUST
meet the OSI Open Source Definition or
the FSF Free Software Definition.


**Maturity Level:** 1

**Rationale:** Ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software [license],
providing clarity on how the code
can be used and shared by others.


**Details:** Add a [LICENSE] file to the project's [repo]
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


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-6, B-B-7 |
| [CRA] | 1.2b |
| CSF | GV.OC-03 |
| SSDF | PO3.2 |



---

### OSPS-LE-03

**Criterion:** The [license] for the source code MUST be
maintained in a standard location within
the project's [repository].


**Maturity Level:** 1

**Rationale:** Ensure that the project's source code is
distributed with the appropriate [license]
terms, making it clear to users and
[contributors] how the code can be used and
shared.


**Details:** Include the project's source code [license] in
the project's [LICENSE] file, COPYING file,
or [LICENSE]/
directory to provide visibility and clarity
on the licensing terms. The filename MAY
have an extension.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-8 |
| [CRA] | 1.2b |
| SSDF | PO3.2 |



---

### OSPS-LE-04

**Criterion:** The [license] for the [released software assets]
MUST meet the OSI Open Source Definition or
the FSF Free Software Definition.


**Maturity Level:** 1

**Rationale:** Ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software [license],
providing clarity on how the code
can be used and shared by others.


**Details:** If a different [license] is included with
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


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-6, B-B-7 |
| [CRA] | 1.2b |
| CSF | GV.OC-03 |
| SSDF | PO3.2 |



---

## Quality

Quality criteria focus on the processes and
practices used to ensure the quality and
reliability of the project's source code and
software assets. These criteria help ensure
that the project's source code is well
maintained, secure, and reliable, reducing the
risk of defects or vulnerabilities in the
software.


### OSPS-QA-01

**Criterion:** The project's source code MUST be publicly
readable and have a static URL.


**Maturity Level:** 1

**Rationale:** Enable users to access and review the
project's source code and history, promoting
transparency and collaboration within the
project community.


**Details:** Use a common [VCS] such as GitHub, GitLab, or
Bitbucket. Ensure the [repository] is publicly
readable. Avoid duplication or mirroring of
[repositories] unless highly visible
documentation clarifies the primary source.
Avoid frequent [changes] to the [repository]
that would impact the [repository] URL.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-B-1 |
| [CRA] | 1.2b, 1.2j |
| OCRE | 486-813, 124-564 |
| SSDF | PS1, PS2, PS3, PW1.2 |



---

### OSPS-QA-02

**Criterion:** The [version control system] MUST contain a
publicly readable record of all [changes]
made, who made the [changes], and when the
[changes] were made.


**Maturity Level:** 1

**Rationale:** Provide transparency and accountability for
[changes] made to the project's [codebase],
enabling users to track modifications and
understand the history of the project.


**Details:** Use a common [VCS] such as GitHub, GitLab, or
Bitbucket to maintain a publicly readable
[commit] history. Avoid squashing or rewriting
[commits] in a way that would obscure the
author of any [commits].


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-B-2, CC-B-3, R-B-5 |
| [CRA] | 1.2b, 1.2f, 1.2j |
| CSF | ID.AM-02, ID.RA-01, ID.RA-08 |
| OC | 4.1.4 |
| OCRE | 486-813, 124-564, 757-271 |
| SSDF | PO3.2, PS1, PS2, PS3, PW1.2, PW2.1 |



---

### OSPS-QA-03

**Criterion:** All [released software assets] MUST be
delivered with a machine-readable list of
all direct and transitive internal software
dependencies with their associated version
identifiers.


**Maturity Level:** 2

**Rationale:** Provide transparency and accountability for
the project's dependencies, enabling users
and [contributors] to understand the
software's dependencies and versions.


**Details:** This may take the form of a software bill of
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


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | Q-S-9 |
| [CRA] | 1.2b, 2.1 |
| CSF | ID.AM-02 |
| OC | 4.3.1 |
| OCRE | 486-813, 124-564, 863-521 |
| SSDF | PO4, PS1 |



---

### OSPS-QA-04

**Criterion:** Any automated [status checks] for [commits] MUST
pass or require manual acknowledgement prior to
merge.


**Maturity Level:** 2

**Rationale:** Ensure that the project's approvers do not
become accustomed to tolerating failing
[status checks], even if arbitrary, because it
increases the risk of overlooking security
vulnerabilities or [defects] identified by
automated checks.


**Details:** Configure the project's version control
system to require that all automated status
checks pass or require manual acknowledgement
before a [commit] can be merged into the
[primary branch].

It is recommended that any optional
[status checks] are NOT configured as a pass
or fail requirement that approvers may be
tempted to bypass.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2f, 1.2k |
| CSF | ID.IM-02 |
| SSDF | PO4.1, PS1 |



---

### OSPS-QA-05

**Criterion:** Any additional [subproject] code [repositories]
produced by the project and compiled into a
[release] MUST enforce security requirements as
applicable to the status and intent of the
respective [codebase].


**Maturity Level:** 3

**Rationale:** Ensure that additional code [repositories] or
[subprojects] produced by the project are held
to a standard that clear and appropriate
for that [codebase].


**Details:** The parent project should maintain a list of
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


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2b, 1.2f |
| OCRE | 486-813, 124-564 |
| SSDF | PO3.2, PO4.1, PS1 |



---

### OSPS-QA-06

**Criterion:** The [version control system] MUST NOT contain
generated executable artifacts.


**Maturity Level:** 2

**Rationale:** Reduce the risk of including generated
executable artifacts in the project's
[version control system], ensuring that only
source code and necessary files are stored
in the [repository].


**Details:** Remove generated executable artifacts
in the project's [version control system].

It is recommended that any scenario where a
generated executable artifact appears
critical to a process such as testing, it
should be instead be generated at build time
or stored separately and fetched during a
specific well-documented pipeline step.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2b |
| OCRE | 486-813, 124-564 |
| SSDF | PS1 |



---

### OSPS-QA-08

**Criterion:** The project MUST use at least one automated
test suite for the source code [repository]
which clearly documents when and how tests
are run.


**Maturity Level:** 3

**Rationale:** 

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | Q-B-4 |
| [CRA] | 2.3 |
| OC | 4.1.5 |
| OCRE | 207-435, 088-377 |
| SSDF | PW8.2 |



---

### OSPS-QA-09

**Criterion:** The [project documentation] MUST include a
policy that all major [changes] to the
software produced by the project should
add or update tests of the functionality
in an [automated test suite].


**Maturity Level:** 3

**Rationale:** 

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | Q-B-8, Q-B-9, Q-B-10, Q-S-2 |
| [CRA] | 2.3 |
| CSF | ID.IM-02 |
| OC | 4.1.5 |
| OCRE | 207-435, 088-377 |
| SSDF | PW8.2 |



---

### OSPS-QA-10

**Criterion:** The project's [version control system] MUST
require at least one non-author approval of
[changes] before merging into the [release] or
[primary branch].


**Maturity Level:** 3

**Rationale:** 

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-G-3 |



---

## Security Assessment

Security Assessment criteria encourage practices that
help ensure that the project is well positioned
to identify and address security vulnerabilities
and threats in the software.


### OSPS-SA-01

**Criterion:** The [project documentation] MUST provide
design documentation demonstrating all
actions and actors within the system.


**Maturity Level:** 2

**Rationale:** Provide an overview of the project's design
and architecture, illustrating the
interactions and components of the system to
help [contributors] and security reviewers
understand the internal logic of the
[released software assets].


**Details:** Include designs in the [project documentation]
that explains the actions and actors. Actors
include any subsystem or entity that can
influence another segment in the system.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-1, B-S-7, B-S-8 |
| [CRA] | 1.2a, 1.2b |
| CSF | ID.AM-02 |
| OCRE | 155-155, 326-704, 068-102, 036-275, 162-655 |
| SSDF | PO.1, PO.2, PO3.2 |



---

### OSPS-SA-02

**Criterion:** The [project documentation] MUST include
descriptions of all external software
interfaces of the [released software assets].


**Maturity Level:** 2

**Rationale:** Provide users and developers with an
understanding of how to interact with the
project's software and integrate it with
other systems, enabling them to use the
software effectively.

**Implementation:** The project README should include a link to
the API documentation, often under a header
such as "Usage", "API Reference", or "Docs".
Documentation badges should also be considered
when recognized (e.g. readthedocs, godoc, etc.).

**Details:** Document all software interfaces (APIs) of
the [released software assets], explaining how
users can interact with the software and
what data is expected or produced.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-B-10, B-S-7 |
| [CRA] | 1.2a, 1.2b |
| CSF | GV.OC-05, ID.AM-01 |
| OC | 4.1.4 |
| OCRE | 155-155, 068-102, 072-713, 820-878 |
| SSDF | PW1.2 |



---

### OSPS-SA-03

**Criterion:** The project MUST perform a [threat modeling] and
[attack surface analysis] to understand and
protect against attacks on critical code
paths, functions, and interactions within
the system.


**Maturity Level:** 3

**Rationale:** Provides project maintainers an understanding of how
the software can be misused or broken allows them
to plan mitigations to close off the potential of 
those threats from occuring.

Providing downstream consumers with this threat model
can assist them in understanding the security capabilities
or gaps that exist within the project and allows them 
to apply their own additional controls and mitigations.

**Implementation:** [Threat modeling] is an activity where the project
looks at the [codebase], associated processes and 
infrastructure, interfaces, key components and 
"thinks like a hacker" and brainstorms how the 
system be be broken or compromised.

Each identified threat is listed out so the project
can then think about how to proactively avoid or 
close off any gaps/vulnerabilities that could arise.

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-S-8 |
| [CRA] | 1.2j, 1.2k |
| CSF | ID.RA-01, ID.RA-04, ID.RA-05, DE.AE-07 |
| OC | 4.1.5 |
| OCRE | 068-102, 154-031, 888-770 |
| SSDF | PO5.1, PW1.1 |



---

### OSPS-SA-04

**Criterion:** The project MUST perform a security
assessment to understand the most likely and
impactful potential security problems that
could occur within the software.


**Maturity Level:** 2

**Rationale:** Performing a security assessment informs both project 
members as well as downstream consumers that the 
project understands what problems could arise within
the software.

Understanding what threats could be realized helps
the project manage and address risk.  This information
is useful to downstream consumers to demonstrate
the security accumen and practices of the project.

**Implementation:** Security assessments can take many forms in a spectrum
ranging from informal to highly rigourous.  At its most 
basic, a security assessment lists potential threats, 
estimates the likelihood/probability of those threats
occuring and the the possible consquences/impact if those
threats do occur.

Software that is deployed into highly-regulated spaces
would benefit from more formal risk management practices
such as [NIST SP 800-171](https://csrc.nist.gov/projects/risk-management/about-rmf), 
[ENISA's Risk Management Framework](https://www.enisa.europa.eu/topics/risk-management), 
or a methodolgy such as [OpenFAIR](https://www.opengroup.org/open-fair).

**Details:** 

| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-W-8, S-G-1 |
| [CRA] | 1.1, 2.2 |
| CSF | ID.RA-04, ID.RA-05, DE.AE-07 |
| OC | 4.1.5 |
| OCRE | 068-102, 307-242, 660-867 |
| SSDF | PO5.1, PW1.1 |



---

## Vulnerability Management

Vulnerability Management criteria focus on the
processes and practices used to identify and
address security vulnerabilities in the project's
software dependencies. These criteria help ensure
that the project is well positioned to respond to
security threats and vulnerabilities in the software.


### OSPS-VM-01

**Criterion:** The [project documentation] MUST include a
policy that defines a threshold for remediation
of [SCA] findings related to vulnerabilities and
[licenses].


**Maturity Level:** 3

**Rationale:** Ensure that the project clearly communicates
the threshold for remediation of [SCA] findings,
including vulnerabilities and [license] issues
in software dependencies.


**Details:** Document a policy in the project that
defines a threshold for remediation of [SCA]
findings related to vulnerabilities and
[licenses]. Include the process for
identifying, prioritizing, and remediating
these findings.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | Q-B-12,  Q-S-9, S-B-14, S-B-15, A-B-3, A-B-8 |
| [CRA] | 1.2a, 1.2b, 1.2c, 2.1, 2.2, 2.3 |
| CSF | GV.RM-05, GV.RM-06, GV.PO-01, GV.PO-02, ID.RA-01, ID.RA-08, ID.IM-02 |
| OC | 4.1.5, 4.2.1, 4.3.2 |
| OCRE | 124-564, 832-555, 611-158, 207-435, 088-377 |
| SSDF | PO.4, PW1.2, PW8.1, RV2.1, RV 2.2 |



---

### OSPS-VM-02

**Criterion:** The [project documentation] MUST include a
policy to address [SCA] violations prior to any
[release].


**Maturity Level:** 3

**Rationale:** Ensure that violations of your [SCA] policy
are addressed before software [releases],
reducing the risk of shipping insecure or
non-compliant software.


**Details:** Document a policy in the project to address
applicable [Software Composition Analysis] 
results before any [release], and add status
checks that verify compliance with that 
policy prior to [release].


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | S-B-14, S-B-15, A-B-3, A-B-8 |
| [CRA] | 1,2a, 1.2c, 2.2, 2.3 |
| CSF | GV.PO-01, GV.PO-02, ID.RA-01, ID.RA-08 |
| OC | 4.1.5 |
| OCRE | 486-813, 833-442, 611-158, 207-435, 088-377 |
| SSDF | PW8.1 |



---

### OSPS-VM-03

**Criterion:** The [project documentation] MUST include a
policy for coordinated vulnerability
reporting, with a clear timeframe for
response.


**Maturity Level:** 2

**Rationale:** Establish a process for reporting and
addressing vulnerabilities in the project,
ensuring that security issues are handled
promptly and transparently.


**Details:** Create a SECURITY.md file at the root of the
directory, outlining the project's policy
for coordinated [vulnerability reporting].
Include a method for reporting
vulnerabilities. Set expectations for the
how the project will respond and address
reported issues.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | R-B-6, R-B-8, R-S-2, S-B-14, S-B-15 |
| [CRA] | 2.1, 2.3, 2.6, 2.7, 2.8 |
| CSF | GV.PO-01, GV.PO-02, ID.RA-01, ID.RA-08 |
| OC | 4.1.5, 4.2.1, 4.3.2 |
| OCRE | 887-750 |
| SSDF | RV1.3 |



---

### OSPS-VM-04

**Criterion:** All proposed [changes] to the project's
[codebase] must be automatically evaluated 
against a documented policy for known
vulnerabilities and blocked in the
event of violations except when declared
and supressed as non exploitable.


**Maturity Level:** 3

**Rationale:** Identify and address [defects] and security weaknesses 
in the project's [codebase] early in the 
development process, reducing the risk of 
shipping insecure software.


**Details:** Create a [status check] in the project's
[version control system] that runs a Static
Application Security Testing (SAST) tool on
all [changes] to the [codebase]. Require that the
[status check] passes before [changes] can be
merged.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | CC-B-9, A-B-1, A-S-1 |
| [CRA] | 1.2a, 1.2b |
| OC | 4.1.5 |
| OCRE | 486-813, 124-564, 757-271 |
| SSDF | PO4.1, RV1.2, RV2.1, RV2.2 |



---

### OSPS-VM-05

**Criterion:** The project publishes contacts and process
for reporting vulnerabilities.


**Maturity Level:** 1

**Rationale:** Reports from researchers and users are an important source for
identifying vulnerabilities in a project. People with
vulnerabilities to report should have a clear understanding of
the process so that they can quickly submit the report to the
project.


**Details:** Create a security.md (or similarly-named) file that contains security 
contacts for the project and provide project's
process for handling vulnerabilities in the project or dependencies.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB | B-S-8 |
| [CRA] | 2.5 |
| CSF | GV.PO-01, GV.PO-02, ID.RA-01 |
| OC | 4.1.1, 4.1.3, 4.1.5, 4.2.2 |
| OCRE | 464-513 |
| SSDF | RV1.3 |



---

### OSPS-VM-06

**Criterion:** The project MUST provide a means for reporting 
security vulnerabilities privately to the security
contacts within the project.


**Maturity Level:** 2

**Rationale:** Security vulnerabilities should not be shared with 
the public until such time the project has been 
provided time to analyze and prepare remediations
to protect users of the project.


**Details:** Enable private bug reporting through [VCS] or other
infrastrucuture.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| BPB |  |
| [CRA] | 1.2a, 1.2b, 2.1, 2.4, 2.6 |
| OCRE | 308-514 |



---

### OSPS-VM-07

**Criterion:** The project MUST publicly publish data about vulnerabilities
discovered within the [codebase].


**Maturity Level:** 2

**Rationale:** Consumers of the project must be informed about
[known vulnerabilities] found within the project.


**Details:** Provide information about [known vulnerabilities] in a predictable
public channel, such as a CVE entry, blog post, or other
medium. To the degree possible, this information should include
affected version(s), how a consumer can determine if they are
vulnerable, and instructions for mitigation or remediation.


| Catalog | Potential Mappings |
| ------- | ------------------ |
| [CRA] | 1.2a, 1.2b, 2.1, 2.4, 2.6 |



---


## Lexicon


### Arbitrary Code

Code provided by an external source that is
executed by a system without validation or
restriction.




### Attack Surface Analysis

Attack Surface Analysis is about mapping out what parts of a system need to
be reviewed and tested for security vulnerabilities. The point of Attack
Surface Analysis is to understand the risk areas in an application, to make
developers and security specialists aware of what parts of the application
are open to attack, to find ways of minimizing this, and to notice when and
how the Attack Surface changes and what this means from a risk perspective.

See OWASP's Attack Surface Analysis Cheat Sheet for more information.



**References:**

  - https://cheatsheetseries.owasp.org/cheatsheets/Attack_Surface_Analysis_Cheat_Sheet.html


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




### Cyber Resilience Act

Regulation (EU) 2024/2847 (Cyber Resilience Act, CRA).
2024 European cybersecurity law that goes into full effect
December 2027.  Focuses on products sold within the European
Union and the cybersecurity and vulnerability management
practices used to create and support the product.



**References:**

  - https://eur-lex.europa.eu/eli/reg/2024/2847/oj


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




### Threat Modeling

Threat modeling is an activity where the project
looks at the codebase, associated processes and 
infrastructure, interfaces, key components and 
"thinks like a hacker" to brainstorm how the 
system be be broken or compromised.

Each identified threat is listed out so the project
can then think about how to proactively avoid or 
close off any gaps/vulnerabilities that could arise.

Examples of threat modeling methodologies include
the Shostack "4 Questions" model, STRIDE, and
tools such as the Elevation of Privilege Threat
Modeling Card Game or Threat Dragon.



**References:**

  - https://github.com/adamshostack/4QuestionFrame

  - https://owasp.org/www-community/Threat_Modeling_Process

  - https://www.microsoft.com/en-us/download/details.aspx?id=20303

  - https://owasp.org/www-project-threat-dragon


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
[Attack Surface Analysis]: #attack-surface-analysis
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
[Cyber Resilience Act]: #cyber-resilience-act
[CRA]: #cyber-resilience-act
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
[Threat Modeling]: #threat-modeling
[Version Identifier]: #version-identifier
[Version Control System]: #version-control-system
[VCS]: #version-control-system
[version control systems]: #version-control-system
[Vulnerability Reporting]: #vulnerability-reporting
[Coordinated Vulnerability Reporting]: #vulnerability-reporting
