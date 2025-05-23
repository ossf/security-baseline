# OSPS Baseline checklist, version: 2025-02-25

## Level 1

- [ ] **OSPS-AC-01.01**: When a user attempts to access a sensitive resource in the project&#39;s version control system, the system MUST require the user to complete a multi-factor authentication process. 
- [ ] **OSPS-AC-02.01**: When a new collaborator is added, the version control system MUST require manual permission assignment, or restrict the collaborator permissions to the lowest available privileges by default. 
- [ ] **OSPS-AC-03.01**: When a direct commit is attempted on the project&#39;s primary branch, an enforcement mechanism MUST prevent the change from being applied. 
- [ ] **OSPS-AC-03.02**: When an attempt is made to delete the project&#39;s primary branch, the version control system MUST treat this as a sensitive activity and require explicit confirmation of intent. 
- [ ] **OSPS-BR-01.01**: When a CI/CD pipeline accepts an input parameter, that parameter MUST be sanitized and validated prior to use in the pipeline. 
- [ ] **OSPS-BR-03.01**: When the project lists a URI as an official project channel, that URI MUST be exclusively delivered using encrypted channels. 
- [ ] **OSPS-DO-01.01**: When the project has made a release, the project documentation MUST include user guides for all basic functionality. 
- [ ] **OSPS-DO-02.01**: When the project has made a release, the project documentation MUST include a guide for reporting defects. 
- [ ] **OSPS-GV-02.01**: While active, the project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles. 
- [ ] **OSPS-GV-03.01**: While active, the project documentation MUST include an explanation of the contribution process. 
- [ ] **OSPS-LE-02.01**: While active, the license for the source code MUST meet the OSI Open Source Definition or the FSF Free Software Definition. 
- [ ] **OSPS-LE-02.02**: While active, the license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition. 
- [ ] **OSPS-LE-03.01**: While active, the license for the source code MUST be maintained in the corresponding repository&#39;s LICENSE file, COPYING file, or LICENSE/ directory. 
- [ ] **OSPS-LE-03.02**: While active, the license for the released software assets MUST be included in the released source code, or in a LICENSE file, COPYING file, or LICENSE/ directory alongside the corresponding release assets. 
- [ ] **OSPS-QA-01.01**: While active, the project&#39;s source code repository MUST be publicly readable at a static URL. 
- [ ] **OSPS-QA-01.02**: The version control system MUST contain a publicly readable record of all changes made, who made the changes, and when the changes were made. 
- [ ] **OSPS-QA-02.01**: When the package management system supports it, the source code repository MUST contain a dependency list that accounts for the direct language dependencies. 
- [ ] **OSPS-QA-04.01**: While active, the project documentation MUST contain a list of any codebases that are considered subprojects or additional repositories. 
- [ ] **OSPS-QA-05.01**: While active, the version control system MUST NOT contain generated executable artifacts. 
- [ ] **OSPS-VM-02.01**: While active, the project documentation MUST contain security contacts. 

## Level 2

- [ ] **OSPS-AC-04.01**: When a CI/CD task is executed with no permissions specified, the project&#39;s version control system MUST default to the lowest available permissions for all activities in the pipeline. 
- [ ] **OSPS-BR-02.01**: When an official release is created, that release MUST be assigned a unique version identifier. 
- [ ] **OSPS-BR-04.01**: When an official release is created, that release MUST contain a descriptive log of functional and security modifications. 
- [ ] **OSPS-BR-05.01**: When a build and release pipeline ingests dependencies, it MUST use standardized tooling where available. 
- [ ] **OSPS-BR-06.01**: When an official release is created, that release MUST be signed or accounted for in a signed manifest including each asset&#39;s cryptographic hashes. 
- [ ] **OSPS-DO-06.01**: When the project has made a release, the project documentation MUST include a description of how the project selects, obtains, and tracks its dependencies. 
- [ ] **OSPS-GV-01.01**: While active, the project documentation MUST include a list of project members with access to sensitive resources. 
- [ ] **OSPS-GV-01.02**: While active, the project documentation MUST include descriptions of the roles and responsibilities for members of the project. 
- [ ] **OSPS-GV-03.02**: While active, the project documentation MUST include a guide for code contributors that includes requirements for acceptable contributions. 
- [ ] **OSPS-LE-01.01**: While active, the version control system MUST require all code contributors to assert that they are legally authorized to make the associated contributions on every commit. 
- [ ] **OSPS-QA-03.01**: When a commit is made to the primary branch, any automated status checks for commits MUST pass or be manually bypassed. 
- [ ] **OSPS-QA-06.01**: Prior to a commit being accepted, the project&#39;s CI/CD pipelines MUST run at least one automated test suite to ensure the changes meet expectations. 
- [ ] **OSPS-SA-01.01**: When the project has made a release, the project documentation MUST include design documentation demonstrating all actions and actors within the system. 
- [ ] **OSPS-SA-02.01**: When the project has made a release, the project documentation MUST include descriptions of all external software interfaces of the released software assets. 
- [ ] **OSPS-SA-03.01**: When the project has made a release, the project MUST perform a security assessment to understand the most likely and impactful potential security problems that could occur within the software. 
- [ ] **OSPS-VM-01.01**: While active, the project documentation MUST include a policy for coordinated vulnerability reporting, with a clear timeframe for response. 
- [ ] **OSPS-VM-03.01**: While active, the project documentation MUST provide a means for reporting security vulnerabilities privately to the security contacts within the project. 
- [ ] **OSPS-VM-04.01**: While active, the project documentation MUST publicly publish data about discovered vulnerabilities. 

## Level 3

- [ ] **OSPS-AC-04.02**: When a job is assigned permissions in a CI/CD pipeline, the source code or configuration MUST only assign the minimum privileges necessary for the corresponding activity. 
- [ ] **OSPS-BR-02.01**: When an official release is created, all assets within that release MUST be clearly associated with the release identifier or another unique identifier for the asset. 
- [ ] **OSPS-DO-03.01**: When the project has made a release, the project documentation MUST contain instructions to verify the integrity and authenticity of the release assets. 
- [ ] **OSPS-DO-04.01**: When the project has made a release, the project documentation MUST include a descriptive statement about the scope and duration of support for each release. 
- [ ] **OSPS-DO-05.01**: When the project has made a release, the project documentation MUST provide a descriptive statement when releases or versions will no longer receive security updates. 
- [ ] **OSPS-GV-04.01**: While active, the project documentation MUST have a policy that code contributors are reviewed prior to granting escalated permissions to sensitive resources. 
- [ ] **OSPS-QA-02.02**: When the project has made a release, all compiled released software assets MUST be delivered with a software bill of materials. 
- [ ] **OSPS-QA-04.02**: When the project has made a release comprising multiple source code repositories, all subprojects MUST enforce security requirements that are as strict or stricter than the primary codebase. 
- [ ] **OSPS-QA-06.02**: While active, project&#39;s documentation MUST clearly document when and how tests are run. 
- [ ] **OSPS-QA-06.03**: While active, the project&#39;s documentation MUST include a policy that all major changes to the software produced by the project should add or update tests of the functionality in an automated test suite. 
- [ ] **OSPS-QA-07.01**: When a commit is made to the primary branch, the project&#39;s version control system MUST require at least one non-author approval of the changes before merging. 
- [ ] **OSPS-SA-03.02**: When the project has made a release, the project MUST perform a threat modeling and attack surface analysis to understand and protect against attacks on critical code paths, functions, and interactions within the system. 
- [ ] **OSPS-VM-04.02**: While active, any vulnerabilities in the software components not affecting the project MUST be accounted for in a VEX document, augmenting the vulnerability report with non-exploitability details. 
- [ ] **OSPS-VM-05.01**: While active, the project documentation MUST include a policy that defines a threshold for remediation of SCA findings related to vulnerabilities and licenses. 
- [ ] **OSPS-VM-05.02**: While active, the project documentation MUST include a policy to address SCA violations prior to any release. 
- [ ] **OSPS-VM-05.03**: While active, all changes to the project&#39;s codebase MUST be automatically evaluated against a documented policy for malicious dependencies and known vulnerabilities in dependencies, then blocked in the event of violations, except when declared and suppressed as non-exploitable. 
- [ ] **OSPS-VM-06.01**: While active, the project documentation MUST include a policy that defines a threshold for remediation of SAST findings. 
- [ ] **OSPS-VM-06.02**: While active, all changes to the project&#39;s codebase MUST be automatically evaluated against a documented policy for security weaknesses and blocked in the event of violations except when declared and suppressed as non-exploitable. 
