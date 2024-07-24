# **Open Source Project Security Baseline**

The SIG is under the [BEST Working Group](https://github.com/ossf/wg-best-practices-os-developers). The goal of this SIG is to evolve [OpenSSF security baseline](https://github.com/ossf/tac/blob/main/process/security_baseline.md) for OpenSSF and Linux Foundation wide adoption, make it easier to discover, adopt and contribute to open source technologies to improve Open Source Software (OSS) security. 

## Motivation

For OpenSSF adoption of the security baseline, there needs to be a home for tracking the adoption, for maintainers to raise issues to refine the security baseline, merge the baseline back to TAC lifecycle, and for OpenSSF to develop the roadmap for the security baseline. It will provide a venue for early adopters to share their findings and reusable code with other maintainers. The pilot adoption builds the foundation for wider adoption of the security baseline in OpenSSF and in Linux Foundation.

This SIG creates a venue for other participating foundations to help evolve the OpenSSF security baseline into a security baseline that can be applied to a broad range of software-based projects. The group will define the right level of risks that the security baseline is applicable for, the effectiveness measurement of the security baseline, and the adoption path of the security baseline at the minimum.

Members of this group will be from various Linux foundations and entities outside of Linux Foundation. Reducing duplicate effort and achieving a higher level of security across Linux Foundation participating foundations is one of the goals of the group.

Right now it's very difficult for new people interested in participating to determine how to participate or what to use. This SIG will provide a technology consumption framework to help open source producers quickly navigate the OSS security technology landscape, discover, adopt and contribute to technical initiatives, help end user organizations large and small to think about adopting OpenSSF/adjacent technical projects and guidance. 

## Objective

Through the effort of this group, open source software is more secure. 

## Scope
* Security adoption and maintenance in OpenSSF
* Security basline adoption and maintenance across Linux Foundation
* OpenSSF/adjacent technollogy consumption framwork 


## Prior Work

This group's work is built on top of the technical output from all the [OpenSSF WGs and Sigs](https://github.com/ossf/tac?tab=readme-ov-file#working-groups-wgs), and [projects](https://github.com/ossf/tac?tab=readme-ov-file#projects). 

## OpenSSF Security Baseline Adoption Pilot Projects

The maintainers of these projects will be the early adopters of the [OpenSSF security baseline](https://github.com/ossf/tac/blob/main/process/security_baseline.md). The maintainers of these pilot projects have committed to fully adopting the seurity baseline by meeting the security baseline requirements for their project life cycle by 9/15/2024, resolving issues 

Detailed adoption plan and tracking will be linked here once it's apprved by the adopting pilot project maintainers.


| Name                       | Repository/Home Page | Notes                                                                                                 | Sponsoring Org | Lifecycle     |Security Baseline requirements|
| ---------------------- | ---------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------- |---------- |---------- |
| GUAC                   | https://github.com/guacsec/guac                  | [Meeting Notes](https://docs.google.com/document/d/1ImSlr_t3WNZ3zWqpmfqkw1mi6_nkv3enkQ7snWDomKA/edit)  | Supply Chain Integrity WG | [Incubating](process/project-lifecycle-documents/guac_incubating.md) |[Security Baseline - Once Sandbox](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---once-sandbox)<br /> and <br />[Security Baseline - To Become Incubating](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---to-become-incubating)<br /> and <br />[Security Baseline - Once Incubating](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---to-become-incubating) |
| OpenVEX | https://github.com/openvex | [Meeting Notes](https://docs.google.com/document/d/1C-L0JDx5O35TjXb6dcyL6ioc5xWUCkdR5kEbZ1uVQto/edit) | Vulnerability Disclosures WG | [Sandbox](process/project-lifecycle-documents/openvex_for_sandbox_stage.md) |[Security Baseline - Once Sandbox](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---once-sandbox) |
| Protobom | http://github.com/bom-squad/protobom | [Meeting Notes](https://docs.google.com/document/d/1bz2BBImzSnLRiBLrA5GehQ0ckW3Vs7Gmtt8R-Olm0QY/edit)  | Security Tooling WG | [Sandbox](process/project-lifecycle-documents/protobom_sandbox_stage.md) |[Security Baseline - Once Sandbox](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---once-sandbox) |
| Repository Service for TUF | https://github.com/repository-service-tuf/repository-service-tuf |  [Meeting Notes](https://docs.google.com/document/d/13a_AtFpPK9WO4PlAN6ciD-G1jiBU3gEDtRD1OUinUFY/edit)  | Securing Software Repositories WG | [Incubating](process/project-lifecycle-documents/repository_service_for_tuf_incubation_stage.md) |[Security Baseline - Once Sandbox](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---once-sandbox) <br /> and <br />[Security Baseline - To Become Incubating](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---to-become-incubating)<br /> and <br />[Security Baseline - Once Incubating](https://github.com/ossf/tac/blob/main/process/security_baseline.md#security-baseline---to-become-incubating) |
 
## Get Involved

*   Official communications occur on the [Security Baseline Mailing List](https://lists.openssf.org/g/openssf-sig-security-baseline). 
*   [Manage your subscriptions to Open SSF mailing lists](https://lists.openssf.org/g/main/subgroups).
*   [Join our Slack](https://app.slack.com/client/T019QHUBYQ3/C07DC6TT2QY)


### Quick Start

*   Need contributors to improve documents, raise awareness of the initiative and increase adoption of the security baseline.
*   Need contributors to evangelize the security baseline for Linux Foundation wide adoption .
*   Create an issue or submit a OR to turn your ideas into reality

### Meeting times
*   Every other Tuesday @ 10:00am EST on [zoom](https://zoom-lfx.platform.linuxfoundation.org/meeting/97740884759?password=5cab7229-2324-4816-81db-517812a088a9)
*   [Meeting Minutes](https://docs.google.com/document/d/16tL1Ln7owIRXSoCKgyYHCs9-JP9iw-ouyk8koGAeHA0/edit)

## Governance

[TODO: Update this link to your specific CHARTER.md file]
The [CHARTER.md](https://github.com/ossf/project-template/blob/main/CHARTER.md) outlines the scope and governance of our group activities.

*   Lead name: Eddie Knight 
*   Co-Lead name: Michael Lieberman

## Intellectual Property

In accordance with the [OpenSSF Charter (PDF)](https://charter.openssf.org/), work produced by this group is licensed as follows:

1. Software source code
* Apache License, Version 2.0, available at https://www.apache.org/licenses/LICENSE-2.0;

## Antitrust Policy Notice

Linux Foundation meetings involve participation by industry competitors, and it is the intention of the Linux Foundation to conduct all of its activities in accordance with applicable antitrust and competition laws. It is therefore extremely important that attendees adhere to meeting agendas, and be aware of, and not participate in, any activities that are prohibited under applicable US state, federal or foreign antitrust and competition laws.

Examples of types of actions that are prohibited at Linux Foundation meetings and in connection with Linux Foundation activities are described in the Linux Foundation Antitrust Policy available at http://www.linuxfoundation.org/antitrust-policy. If you have questions about these matters, please contact your company counsel, or if you are a member of the Linux Foundation, feel free to contact Andrew Updegrove of the firm of Gesmer Updegrove LLP, which provides legal counsel to the Linux Foundation.
