# Open Source Project Security Lexicon

## Definitions

### Arbitrary Code

Code provided by an external source that is executed by a system without validation or restriction.

### Build and Release Pipeline

A series of automated processes that compile and deploy software. Similar to the generic term [CI/CD Pipelines](#cicd-pipelines), but this term excludes some pipelines, such as pre-merge status checks.

### Changes

Any alteration of the project's [codebase](#codebase), [CI/CD Pipelines](#cicd-pipelines), or [documentation](#project-documentation). This may include addition, deletion, or modification of content.

### CI/CD Pipelines

Automated pipelines for Continuous Integration and Continuous Delivery. Responsible for building, testing, and delivering [changes](#changes). These pipelines integrate contributions frequently, enabling rapid and reliable software delivery. CI focuses on testing and building code, while CD delivers software to users.

In the context of the Open Source Project Security Baseline, CI/CD Pipelines does not involve Continuous Deployment.

### Code Contributors

Entities who [commit](#commit) code or [documentation](#project-documentation) to the project. Code contributors include [collaborators](#collaborator) or external participants who submit [changes](#changes).

In the context of the Open Source Project Security Baseline, code contributors does not address non-code contributions such as designing, triaging, reviewing, or testing.

### Codebase

The collection of source code and related assets that make up the project. The codebase includes all files necessary to build and test the software. Lives in the [repository](#repository), sometimes alongside [documentation](#project-documentation) and [CI/CD pipelines](#cicd-pipelines). The contents of the codebase are the primary deliverable in a [release](#release).

### Collaborator

A user with access to the project's [version control system](#version-control-system) who can [commit](#commit) to the [codebase](#codebase), approve [changes](#changes), or manage the [repository](#repository) settings. Collaborators may have varying permission levels based on their role in the project.

### Commit

A record of a single [change](#changes) submitted to the [version control system](#version-control-system). Each commit includes details such as the modifications made, the contributor who made them, and the timestamp of the [change](#changes).

### Defects

Errors or flaws in the software that cause it to produce incorrect or unintended results, or to behave in an unintended way. Defects can include bugs, [vulnerabilities](#exploitable-vulnerabilities), or other issues that impact the software's functionality or security. Defects may have originally been intentional, but a change in environment or understanding has made them undesirable.

### Exploitable Vulnerabilities

[Defects](#defects) in the software that can be leveraged by attackers to gain unauthorized access, execute [arbitrary code](#arbitrary-code), or cause other undesired outcomes.

### Known Vulnerabilities

Publicly acknowledged [exploitable vulnerabilities](#exploitable-vulnerabilities) that have been identified within the software. These vulnerabilities often have associated advisories, patches, or recommended mitigations.

### Multi-factor Authentication (MFA)

An authentication method that requires two or more verification factors (e.g., a password and a token) to gain access to a resource. This method strengthens security by requiring multiple forms of identification.

### Primary Branch

The main development branch in the [version control system](#version-control-system), representing the latest stable codebase. [Releases](#release) are typically made from this branch. Commonly named `main` or `master`. In some situations where branches are not featured, a [repository](#repository) with forked repositories will have the original repo acting as as an equivalent to the primary branch.

### Project Documentation

Written materials related to the project, such as user guides, developer guides, and contribution guidelines. These documents help users and developers understand, contribute to, and interact with the software. At release time, this may include [provenance] information, licensing details, and other metadata.

### Provenance

Information about the origin and history of the [released software assets](#released-software-assets). This may include details about its development, dependencies, vulnerabilities, contributors, and licensing.

### Release

- _(verb)_ The process of making a version-controlled bundle of assets available to users.
- _(noun)_ A version-controlled bundle of code, [documentation](#project-documentation), and other assets made available to users. A release often includes release notes that describe the changes.

### Released Software Assets

Deliverables provided to users as part of a [release](#release). These assets can include binaries, libraries, or containers.

### Releases

Specific versions of the software that have been made available to users, typically accompanied by a [unique version identifier](#unique-version-identifier). Each release includes notes on updates and changes.

### Repository

A storage location managed by a [version control system](#version-control-system) where the project's [code](#commit), [documentation](#project-documentation), and other resources are stored. It tracks [changes](#changes), manages collaborator permissions, and includes configuration options such as branch protection and access controls.

### Software Composition Analysis (SCA)

The process of identifying and cataloging all components and dependencies in a software codebase. SCA is essential for managing security vulnerabilities and ensuring compliance with organizational policies.

### Status Checks

Automated tests or validations that run on [commits](#commit) before they are merged. Status checks ensure that any [changes](#changes) meet the project's quality and security standards.

### Unique Version Identifier

A label assigned to a specific [release](#release) of the software, such as `v1.2.3`. Commonly recommended formats are [Semantic Versioning](https://semver.org/) or [Calendar Versioning](https://calver.org).

### Version Control System (VCS)

A tool that tracks [changes](#changes) to files over time and facilitates collaboration among contributors. Examples of version control systems include Git, Subversion, and Mercurial.

### Vulnerability Reporting

The act of identifying and documenting [exploitable vulnerabilities](#known-vulnerabilities) in [released software assets](#released-software-assets). This may include privately or openly reporting vulnerabilities to maintainers, security teams, or the public, as well as tracking the resolution of these vulnerabilities.
