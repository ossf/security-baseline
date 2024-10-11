# Criteria Breakdown

This document provides details for each criteria ID, including what it protects against and how it might be implemented in common situations.

## Maturity Level 1 (Sandbox)

### mfa

- **What it protects against:** Protects against unauthorized access to sensitive areas of the project's repository, reducing the risk of account compromise or insider threats.
- **How it might be implemented:** Use version control systems like GitHub or GitLab, configuring multi-factor authentication (MFA) for all users with elevated privileges. Ensure that repository settings and sensitive data (e.g., secrets) are locked behind MFA.

### restrict-privileges

- **What it protects against:** Protects against accidental or malicious changes by limiting access to sensitive areas and functionality to only those who need it.
- **How it might be implemented:** Enforce the principle of least privilege by configuring roles in GitHub, GitLab, or similar tools, so collaborators only have permissions to perform actions they are authorized for (e.g., preventing pushes to protected branches).

### direct-commits

- **What it protects against:** Prevents accidental modification of the main development branch, reducing the risk of untested or unreviewed code being integrated.
- **How it might be implemented:** Use branch protection features in your version control system to require pull requests and reviews before code is merged into the primary branch.

### branch-deletion

- **What it protects against:** Prevents accidental loss of the primary development branch, which could cause significant disruptions to ongoing work.
- **How it might be implemented:** Configure branch protection rules in your version control system to disallow deletion of the primary branch.

### arbitrary-code

- **What it protects against:** Prevents the execution of untrusted or malicious code, reducing risks like remote code execution vulnerabilities.
- **How it might be implemented:** Ensure build pipelines only execute code from trusted sources. Use signed commits and verified pull requests to control the source of code being built.

### version-id-required

- **What it protects against:** Ensures traceability and version control, preventing confusion over which software version is being used, which can be critical during vulnerability or defect assessments.
- **How it might be implemented:** Use automated tools to generate and enforce version identifiers during the build process. Utilize semantic versioning or calendar versioning to label releases clearly.

### website-delivery

- **What it protects against:** Protects against man-in-the-middle attacks or data interception by ensuring that web resources related to the project are securely delivered.
- **How it might be implemented:** Implement HTTPS or SSH for all web services, including project websites, API endpoints, and documentation portals.

### public-discussions

- **What it protects against:** Protects against lack of transparency in decision-making and helps mitigate misunderstandings in the community about changes to the project.
- **How it might be implemented:** Use public issue trackers (e.g., GitHub Issues), mailing lists, or forums where contributors and users can discuss and propose changes openly.

### basic-docs

- **What it protects against:** Prevents confusion or misusage of the software by ensuring users have access to clear guides on how to interact with its basic functionality.
- **How it might be implemented:** Include a README and user guide in the repository, covering installation, setup, and basic usage instructions.

### contribution-docs

- **What it protects against:** Prevents potential contributors from feeling blocked or confused by a lack of guidance on how to participate in the project.
- **How it might be implemented:** Document the contribution process, including how to submit pull requests, coding standards, and review processes in a CONTRIBUTING.md file.

### public-repo

- **What it protects against:** Ensures project transparency and access, preventing the risk of closed or restricted visibility that may hinder community participation.
- **How it might be implemented:** Host your project in a publicly accessible version control platform like GitHub or GitLab, and ensure the repository settings are configured for public visibility.

### change-tracking

- **What it protects against:** Prevents lack of clarity around what changes have been made in the project, helping track contributions and ensuring accountability.
- **How it might be implemented:** Maintain a CHANGELOG.md file or use Git tags to keep a record of what changes were made, by whom, and when, making this information publicly accessible.

### contributor-assertion

- **What it protects against:** Mitigates the risk of intellectual property issues by ensuring all contributors assert their legal right to contribute the code.
- **How it might be implemented:** Integrate contributor license agreements (CLAs) or developer certificates of origin (DCOs) that must be agreed to before contributors can make a commit.

### license-standard

- **What it protects against:** Prevents legal ambiguity or misuse of the project's code by ensuring it is licensed under recognized open-source terms.
- **How it might be implemented:** Choose a license from OSI- or FSF-approved licenses and place the LICENSE file at the root of the repository.

### license-location

- **What it protects against:** Ensures that the project's licensing terms are easily accessible, reducing the risk of unclear or unenforceable licensing.
- **How it might be implemented:** Store the project's license in a clearly visible location (e.g., root of the repository, LICENSE directory).

### release-license

- **What it protects against:** Ensures that distributed software follows proper licensing terms, reducing the risk of violating open-source or proprietary licenses.
- **How it might be implemented:** Clearly state the licensing terms for distributed binaries or software artifacts, ensuring they align with the source code license or have a distinct license if necessary.

## Maturity Level 2 (Incubating)

### pipeline-permissions

- **What it protects against:** Reduces the risk of privilege escalation or unauthorized access within the CI/CD pipeline.
- **How it might be implemented:** Limit permissions in CI/CD tools to the minimum required for each step of the pipeline. Elevate permissions only for critical operations (e.g., deployment).

### contributor-review

- **What it protects against:** Ensures that only trusted contributors have access to sensitive resources, reducing the risk of insider threats or unintentional damage.
- **How it might be implemented:** Set up a review process for any contributor requesting escalated privileges. Ensure all access to sensitive resources is logged and reviewed periodically.

### automated-pipelines

- **What it protects against:** Prevents inconsistencies or errors in the release process by ensuring that all software releases follow a standardized, repeatable process.
- **How it might be implemented:** Use tools like Jenkins, GitLab CI, or GitHub Actions to automate builds and releases, ensuring that the process is identical each time.

### standardized-deps

- **What it protects against:** Mitigates the risk of introducing vulnerabilities or incompatible software by standardizing dependency management across builds.
- **How it might be implemented:** Use package managers that enforce version constraints (e.g., NPM, Maven, Go modules) and run vulnerability scanning on dependencies during the build process.

### change-log

- **What it protects against:** Ensures that functional and security-related changes are clearly documented, aiding in transparency and security assessments.
- **How it might be implemented:** Update the CHANGELOG.md with details of each release, particularly focusing on security fixes and functional modifications.

## Maturity Level 2 (Incubating)

### vuln-report

- **What it protects against:** Ensures vulnerabilities are reported in a structured, coordinated way, reducing the likelihood of undisclosed vulnerabilities remaining in the system and allowing maintainers to address issues proactively.
- **How it might be implemented:** Create a SECURITY.md file that includes instructions for vulnerability reporting, setting expectations for response times, and providing contact details for security-related issues.

### defect-reporting

- **What it protects against:** Ensures that issues or bugs in the software are tracked and managed effectively, preventing defects from being overlooked or left unresolved.
- **How it might be implemented:** Set up an issue tracker in GitHub, GitLab, or other project management tools to allow users and contributors to report defects. Define the process for triaging, prioritizing, and resolving defects.

### contributor-guide

- **What it protects against:** Helps prevent confusion and mistakes from new contributors by clearly outlining expectations and requirements for contributing code.
- **How it might be implemented:** Document contribution guidelines, including how to write quality code, testing standards, and how contributions are reviewed and merged into the project. Store this information in a CONTRIBUTING.md file.

### design-documentation

- **What it protects against:** Prevents misunderstandings and misalignment on how the software system operates, reducing the risk of introducing changes that conflict with the project's design principles.
- **How it might be implemented:** Maintain up-to-date design documentation that includes architectural diagrams, descriptions of key components, and the interactions between different modules of the system. This can be included as part of the project documentation.

### dependency-list

- **What it protects against:** Helps mitigate the risk of using vulnerable or outdated dependencies by providing a clear, up-to-date list of all external components.
- **How it might be implemented:** Use dependency management tools (e.g., npm, pip, Maven) to generate a list of dependencies. Ensure this list is included in the release artifacts, ideally in a machine-readable format (e.g., JSON or YAML).

### status-checks

- **What it protects against:** Prevents broken or insecure code from being merged into the main branch by ensuring that automated checks are run on every commit or pull request.
- **How it might be implemented:** Configure status checks in your CI/CD system (e.g., GitHub Actions, Jenkins) that automatically run tests, security scans, or linting before code can be merged. Require these checks to pass before a merge is allowed.

### repo-security

- **What it protects against:** Ensures that additional repositories related to the project are equally secure, preventing lax security practices from being applied to related codebases.
- **How it might be implemented:** Enforce security requirements (e.g., MFA, branch protection, vulnerability scanning) for all repositories associated with the project, not just the primary code repository.

### no-executables

- **What it protects against:** Reduces the risk of including generated or compiled artifacts that may be outdated or could obscure the origin of the code, ensuring that only source code is versioned.
- **How it might be implemented:** Add rules to `.gitignore` files that prevent executables or generated files (e.g., `.exe`, `.jar`, `.dll`) from being committed to the repository. Ensure these artifacts are generated through the build process instead.

## Maturity Level 3 (Graduated)

### mfa-sms

- **What it protects against:** Strengthens security by preventing the use of SMS-based MFA, which is vulnerable to attacks like SIM swapping or interception.
- **How it might be implemented:** Configure your version control system to use stronger MFA methods such as TOTP (Time-based One-Time Password) apps or hardware security keys. Disable SMS-based MFA in the security settings.

### sca-policy

- **What it protects against:** Prevents the use of insecure or outdated third-party components by ensuring that Software Composition Analysis (SCA) is conducted regularly and issues are addressed before each release.
- **How it might be implemented:** Use SCA tools like OWASP Dependency-Check, Snyk, or GitHub's Dependabot to scan for vulnerabilities in dependencies. Define a policy that vulnerabilities must be remediated or approved before a release can proceed.

### vuln-evaluation

- **What it protects against:** Ensures that known vulnerabilities are consistently monitored and evaluated, reducing the risk of exposure to attacks based on unaddressed vulnerabilities.
- **How it might be implemented:** Define a process in the project documentation for regularly reviewing known vulnerabilities (e.g., using CVE databases or vulnerability scanning tools) and ensure that all vulnerabilities are either fixed or deemed non-exploitable.

### interface-descriptions

- **What it protects against:** Prevents misunderstanding of how software interacts with other components, reducing the risk of integration issues or incorrect usage.
- **How it might be implemented:** Document all input/output interfaces, including APIs, configuration options, and data formats. Ensure these descriptions are clear and kept up to date, ideally as part of the project documentation or README.

