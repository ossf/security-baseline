# Introduction
## Context and Scope
This document captures technologies that are hosted in OpenSSF. The document covers some of the Open Source Software(OSS) security tools and frameworks that make these OpenSSF technologies possible. The technical stack in this document demonstrates how OSS security technologies make software supply chain more secure. 

The goals of this document are:
1. To help open source producers quickly navigate the OSS security technology landscape, discover, adopt and contribute to technical initiitves. 
2. To provide an easy model for our end user organizations large and small to have a framework/reference architecture to help them think about adopting OpenSSF technical projects and guidance. 

This document is based on the prio work conducted in OpenSSF [BEST Working Group](https://github.com/ossf/wg-best-practices-os-developers)
* [Open Source Software Supply Chain Threats, Threats Definition](https://github.com/ossf/toolbelt/tree/main/threats). 
* [The Security Toolbelt](https://github.com/ossf/Diagrammers-Society/blob/main/drawings/Sterling%20Toolchain%20Patterns.svg).

This document serves as a map that takes readers from here to OpenSSF and adjacent technologies. Details of the technologies are out of the scope of this document. The details of how consumers manage open source software dependencies are out of the scope of this document. 

## Intended Audience
Open source software maintainers, contributors and consumers are the intended audience of this document.

The document answers these questions:
As an open source software producer, I'd like to adopt open source security technologies to produce more secure open source software. What resources does OpenSSF provide? 
As an open source software consumer, I'd like to adopt open source security technologies to consume secure open source software. What resources OpenSSF provide? 
As an open source software producer and a consumer, I'd like to contribute to open source security technologies, which OpenSSF projects need contributors? 

## Nomenclature
The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in RFC 2119.

# Open Source Software Dependency Management 
Open source software is the foundation of the digital world we live in. It powers critical infrastructure and impacts people's everyday lives. 

[SonaType latest state of the software supply chain report](https://www.sonatype.com/state-of-the-software-supply-chain/Introduction) has show among the monitored OSS ecosystems:
* An average 29% of year-over year growth of OSS projects
* Demands for open source ecosystems are still growing, measured by the numder of download requests
* The number of malicious packages tripled in 2023 compared to 2022. 
* In September 2023, 245,032 malicious packages discovered, 2x all previous years (2019-2022) combinedthe 
  
[Synopsys latest open source security and risk analysis report](https://www.synopsys.com/software-integrity/resources/analyst-reports/open-source-security-risk-analysis.html?utm_source=google&utm_medium=cpc&utm_term=open_source_security_tool&utm_campaign=G_S_OSSRA_BMM&cmp=ps-SIG-G_S_OSSRA_BMM&gad_source=1&gclid=CjwKCAjwko21BhAPEiwAwfaQCKkvVOJOS1uu4S7zF1p9uqhcdwrCNXNbMl7EcD5L0SIgS58-jUj_OxoC7FwQAvD_BwE#introMenu) has shown that among the surveyed comemrcial software:
* 96% of the total codebase contained open source.
* 77% of all code in the total code base originated from open source
* 84% of the codebases contained at least one open source vulnerability
* 54% increase in codedebases containing high-risk vulnerabilities in the past year

Open source software dependency management is critical to sustain software security and reliability. 

## Open Source Software Supply Chain Threats

Security threats exist in every link of the OSS supply chain, from upstream maintainers to consumers. OSS prodcuers are also OSS consumers. For detailes of the threats, check out the threat modeling that OpenSSF [BEST Working Group](https://github.com/ossf/wg-best-practices-os-developers) and [End User Working Group](https://github.com/ossf/wg-endusers) have conducted:
* [Threats against OSS producers](https://github.com/ossf/toolbelt/blob/main/threats/Developer_Threats.md)
* [Threats against OSS producers' development environments](https://github.com/ossf/toolbelt/blob/main/threats/Developer_Environment_Threats.md)
* [Threats against OSS Source Code Management(SCM) systems](https://github.com/ossf/toolbelt/blob/main/threats/Source_Code_Management_Threats.md)
* [Threats against OSS CI/CD systems](https://github.com/ossf/toolbelt/blob/main/threats/Build%2BCI_Threats.md)
* [Threats against OSS publication/distribution](https://github.com/ossf/toolbelt/blob/main/threats/Publication_Threats.md)
* [Threat Model of Enterprise Open Source Supply Chains](https://docs.google.com/document/d/1kNCETEfm2_Pm9dFwmDJfph0TtUCK1-3ERL4OjA9UKYI/edit)

Every OSS threat category requires OSS dependency management to reduce the risks from the OSS supply chain threats. The diagram below demonstrates the significance of choosing safe dependencies by OSS producers when start developing their software, securing the development environment, securing the SCM systems to reduce the risks of unauthorized code which leads to malicious dependencies, ingesting safe dependencies into the CI process from package registries, and securing the package registries to prevent malicioud packages from being uploaded and legit packages from being compromised. OSS consumers need to manage depenendencies to ensure safe OSS consumption to sustain their software/product/service security and reliability.    

![OSS Supply Chain Threats](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_OSS_Supply_Chain_Threats.jpg)

## Open source software Security Technologies for OSS Dependency Management 

This section provides an overview of the technology stack that provides OSS dependency management , and dives deeper into each of the below areas. Each area depends on the areas below it to provide technology inputs.  
* [Dependency Ingestion Policy and Enforcement](#dependency-ingestion-policy-and-enforcement)
* [Security Insights into OSS Projects](#security-insights-into-oss-projects)
* [Generating Verifiable Cryptographically Signed Artifacts Attestation, Metadata](#generating-verifiable-cryptographically-signed-artifacts-attestation-metadata)
* [Ecosystem Support](#ecosystem-support)
* [Frameworks, Specifications, Standards, Education and Training](#frameworks-specifications-standards-education-and-training)
  

### Overview

OpenSSF and its sister/brother foundations provide OSS technologies to raise security awareness through education and training, establish best software development practices, guides, security standards, frameworks, specifications, provide security tools to improve open source software dependency management. The security tools are built on top of the ecosystem support to make open source software ecosystems more secure.

Every layer in the following diagram provides support to the layers above it, relies on the underneath layers for foundational support. 

![overview](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework%20_Overview.jpg)

### Dependency Ingestion Policy and Enforcement
To sustain OSS dependency management at scale, with accuracy and efficiency, we need technical solutions that can define software ingestion policies and enforce those policies. For example, if an open source software has not been maintained actively, the software should not be introduced into your code base. 

The following diagram has illustrated that dependency ingestion policy and enforcement is a very green field on OSS producer side (This may change soon with Minder joining OpenSSF as an OSS project). The techologies in this layer will evolve as we gain more [Security Insights into OSS Projects](#security-insights-into-oss-projects). On the customer side, organizations have their own governance and technologies in managing OSS dependencies. The foundational layer 

[giituf](https://github.com/gittuf/gittuf) is an OpenSSF [sandbox](https://github.com/ossf/tac/blob/main/process/project-lifecycle.md#stages---definitions--expectations) project hosted in the [Supply Chain Integrity Working Group](https://github.com/ossf/wg-supply-chain-integrity). It provides a security layer over Git. It's in active development phase and conducts policy validation. 

[Witness](https://github.com/in-toto/witness) is a CNCF project under the [in-toto](https://github.com/in-toto) umbrella that protects software supply chain integrity. Witness comes with an OPA policy engine that verifies the software attestation and ensures that your dependency was handled safely from source to deployment. Witness in-toto attestation feature is available for CI systems on OSS producer side and consumer side  

[Repository Service for TUF(RSTUF](https://github.com/repository-service-tuf/repository-service-tuf) is an OpenSSF [incubating](https://github.com/ossf/tac/blob/main/process/project-lifecycle.md#stages---definitions--expectations) project. The verification component of RSTUF provides metadata verification to ensure the software that's being ingested is as expected. RSTUF metadata verification can be used by both CI systems on OSS producer side and consumer side.    

The target state of achieving end to end OSS supply chain security is to have policy enforcement at every link of the supply chain. As software flows from OSS producers' development environment into the SCM system, a Policy Enforcement Point (PEP) should only allow the software into the SCM system when the software meets pre-defined policies, CI system should fail the integration that's triggered by an SCM event if PEP policy validation fails, package registry should block a package being published or prevent a package from being distributed when PEP identified policy violation. Similar policy enforcement should be adopted on the consumer side as well.

![State](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_state.jpg)
![policy and enforcement](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework%20_Ingestion_Policy_and_Enforcement.jpg)

### Security Insights into OSS Projects
Gaining insights into OSS Projects is the collective effort of multiple technologies across multiple layers:
* Security Metadata Data Aggregation and Synthesis ( Health Metrics, SBOM, VEX, OSV,  in-toto Attestation, SLSA Provenance, TUF Metadata )
  * [GUAC](https://github.com/guacsec/guac) is an OpenSSF [sandbox](https://github.com/ossf/tac/blob/main/process/project-lifecycle.md#stages---definitions--expectations) project hosted in the [Supply Chain Integrity Working Group](https://github.com/ossf/wg-supply-chain-integrity). GUAC provides both API and UI functionalities. The software aggregates health metrics information via Scorecard API, Scorecard Structured Results API in the future, deps.dev API and OSV API. It also aggregates SBOM from projects from a centralized storage where OSS projects publish SBOM, VEX, OSV, in-toto Attestation, SLSA Provenance, TUF Metadata to, or pull these information from OSS projects. It's intended to be used by both OSS producers and consumers to make informed decisions on choosing an OSS dependency. deps.dev is provided by Google as a free service for pubic consumption. 
* Dependency Graph, Licenses, Vulnerabilities, Privileged Access
  * deps.dev is a free service hosted by Google. The service provides details of an OSS software dependencies, dependents, vulnerabilities and license information. It also provides Google assured open source software. [OSV](https://osv.dev/) is an open source database for open source. It provides vulnerability information to deps.dev. OSV ingests vulnerability data from 20+ security advisories and provide vulnerability information at the package level, more granular than NVD. [OSV schema](https://github.com/ossf/osv-schema) specification provides data model for OSV database, and is hosted in OpenSSF [Vulnerability Disclosures Working Group](https://github.com/ossf/wg-vulnerability-disclosures). 
* OSS Security Health Assessment
  * [Scorecard](https://github.com/ossf/scorecard) is a widely adopted OpenSSF project in OSS ecosystems. It assesses an open source project's security health posture, reports findings at various risk levels and provides recommendations on improving the project's security posture. For an OpenSSF project to enter into incubating stage, the project MUST adopt Scorecard per OpenSSF [TAC requirements](https://github.com/ossf/tac/blob/main/process/TI-Gives%2BGets.md#givesrequirements-1). Scorecard has [launched V5](https://github.com/ossf/scorecard/releases/tag/v5.0.0) focusing on structured results which will accelerate the development of security policy enforcement.  
* OSS Security Policy Enforcement
  * [Allstar](https://github.com/ossf/allstar) is a sibling project of Scorecard, hosted in [Securing Critical Projects Working Group](https://github.com/ossf/wg-securing-critical-projects). It's in the process of consolidating into the same working group as Scorecard. It's capable of enforcing rules base on Scorecard checks at GitHub organization level. It features raising alerts once it detects risky behaviors or malpractices and creates a pull request to signal maintainers to take action.
  * [Trusted Publisher](https://repos.openssf.org/trusted-publishers-for-all-package-repositories) is an initiative in OpenSSF [Securing Software Repositories Working Group](https://github.com/ossf/wg-securing-software-repos). It allows a package repository like the Python Package Index (PyPI) to authenticate an identity using a short lived OIDC token. It works in conjunction with a trusted publishing policy on the package repository/registry side. Trusted Publishers ensures a package is coming from a trusted CI system, workflow, hosted machine or build pipeline.  
   
![State](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_state.jpg)
![Security Insights](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_Security_Insights.jpg)

###  Generating Verifiable Cryptographically Signed Artifacts Attestation, Metadata
![State](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_state.jpg)
![Security Metadata](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework%20_Attestation.jpg)

### Ecosystem Support
![Ecosystem Support](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_Ecosystem_Support.jpg)

### Frameworks, Specifications, Standards, Education and Training
![State](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_state.jpg)
![Frameworks, Specifications, Standards, Education and Training](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_Framework_Specifications_Standards_Education.jpg)

