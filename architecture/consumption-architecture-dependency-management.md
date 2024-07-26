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

Security threats exist in every link of the OSS supply chain, from upstream maintainers to consumers. OSS prodcuers are also OSS consumers. For detailes of the threats, check out the threat modeling that OpenSSF [BEST WG](https://github.com/ossf/wg-best-practices-os-developers) and [End User WG](https://github.com/ossf/wg-endusers) have conducted:
* [Threats against OSS producers](https://github.com/ossf/toolbelt/blob/main/threats/Developer_Threats.md)
* [Threats against OSS producers' development environments](https://github.com/ossf/toolbelt/blob/main/threats/Developer_Environment_Threats.md)
* [Threats against OSS Source Code Management(SCM) systems](https://github.com/ossf/toolbelt/blob/main/threats/Source_Code_Management_Threats.md)
* [Threats against OSS CI/CD systems](https://github.com/ossf/toolbelt/blob/main/threats/Build%2BCI_Threats.md)
* [Threats against OSS publication/distribution](https://github.com/ossf/toolbelt/blob/main/threats/Publication_Threats.md)
* [Threat Model of Enterprise Open Source Supply Chains](https://docs.google.com/document/d/1kNCETEfm2_Pm9dFwmDJfph0TtUCK1-3ERL4OjA9UKYI/edit)

Every OSS threat category requires OSS depenednecy management to reduce the risks from the OSS supply chain threats.      

![OSS Supply Chain Threats](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_OSS_Supply_Chain_Threats.jpg)

## Open source software Security Technologies for OSS Dependency Management 

This section provides an overview of the technology stack that provides OSS dependency management , and dives deeper into each of the below areas. Each area depends on the areas below it to provide technology inputs.  
* [Dependency Ingestion Policy and Enforcement](#dependency-ingestion-policy-and-enforcement)
* [Security Insights into OSS Projects](#security-insights-into-oss-projects)
* [Generating Verifiable Cryptographically Signed Artifacts Attestation, Metadata](#generating-verifiable-cryptographically-signed-artifacts-attestation-metadata)
* [Ecosystem Support](#ecosystem-support)
* [Frameworks, Specifications, Standards, Education and Training](#frameworks-specifications-standards-education-and-training)
  

### Overview

OpenSSF and its sister/brother foundations provide OSS technologies to raise security awareness through education and training, establish best practices and standards,  

![overview](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework%20_Overview.jpg)

### Dependency Ingestion Policy and Enforcement
![State](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework_state.jpg)
![policy and enforcement](https://github.com/Danajoyluck/security-baseline/blob/Danajoyluck-patch-1/architecture/images/OpenSSF_Practitioner_Framework%20_Ingestion_Policy_and_Enforcement.jpg)

### Security Insights into OSS Projects
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

