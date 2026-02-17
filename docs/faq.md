# Frequently Asked Questions

* Contents
{:toc}

## What’s the purpose of the OSPS Baseline?

The OSPS Baseline is designed to establish a minimum set of security-related best practices for open source software projects, relative to their maturity level.
It aims to help maintainers, contributors, and users quickly understand and adopt fundamental security steps — like enabling secure workflows, setting up responsible disclosure policies, and maintaining basic project hygiene.
By meeting this “baseline,” a project signals that it has taken essential measures to reduce the risk of common vulnerabilities and improve overall trustworthiness of the project to its adopters and contributors.

## What problem is the OSPS Baseline attempting to solve?

Open source project security, despite being transparent and observable by nature, is often obscure for adopters.
It can be difficult to understand what a project has done to secure itself and its development as well as being difficult for maintainers to understand how to get started in improving the security posture of their project.
There are a number of different commercial or industry frameworks, standards, and ‘practices’ that overlap, are pay-to-play, or otherwise not designed for the flexibility and differentiation inherent to all open source projects.

The OSPS Baseline is an open source approach to solving these problems.
By taking an open source first approach in reconciling these different practices and frameworks, we can provide both the intent, examples, and reasoning for why these security controls should be applied and how they could be applied to projects of varying sizes, complexities, and source forges. 

## How is the OSPS Baseline different from OpenSSF Scorecard and the Best Practices Badge? How is it different from other open source security initiatives?

The OSPS Baseline is a minimal security checklist – you can think of it like the minimum set of things you do to make sure your identity isn't compromised or that you get your kid to school on time every day with no issues.
It is intended to account for gaps and challenges open source projects experience with achieving [OpenSSF Scorecard](https://scorecard.dev), an automated grading tool.
OSPS baseline covers security-impacting criteria not reflected in the [Best Practices Badges](https://www.bestpractices.dev/en); a holistic self-assessment covering more than just security.

There may be some overlap in the OSPS Baseline with other open source security initiatives, and we cannot possibly know what all is out there (CISA has one, NIST has the SSDF).
However, the OSPS Baseline was developed with special attention and consideration by open source contributors, maintainers, and technical leaders who have been working in and alongside open source projects for decades, facilitating feedback, and driving these criteria through open governance and practices in order for them to be achievable by all projects.
We are partnering with all manner of projects to test these criteria for viability, and continually refining and clarifying the verbiage for everyone’s benefit.
We have also provided a crosswalk with mappings to existing standards, frameworks, and other control schemas to make it easier to understand how fulfilling criteria are applicable across multiple sets of requirements and how a project with an OSPS Baseline level meets those controls.
These mappings are intended to show a relationship between OSPS Baseline and external control catalogs.
*Mappings are not a guarantee of compliance with listed external catalogs and are not a substitute for audits.*

## How does the baseline relate to established frameworks such as NIST Cybersecurity Framework, ISO/IEC 27001, and LFX GenAI Framework? Is it a substitute for formal regulatory compliance or certification processes (ie. IEEE has a certification 'badge')?

The OSPS Baseline can be used as a reference point to existing established frameworks when crosswalked to them.
Since open source projects rarely undergo formal regulatory compliance or pursue a certification, the OSPS Baseline level a project achieves should not be considered a substitute.
Certification programs often have financial impediments for open source projects (typically addressed by a sponsor or fund) due to the independent verification process conducted by those organizations.

## What should the Baseline not be used for?

The baseline is not intended to be used as a security comparison tool between projects and it is not intended to be used as a scoring or grading mechanism about a project.
Every project who wishes to implement a level of the Baseline is doing so to improve their security posture and as a result may meet the criteria in their own way.
Not all validation tooling is capable of discovering implementation of the criteria for all open source projects, and therefore would be an unreliable source of accuracy into the true security posture of a project.

Adopters of projects are strongly encouraged to understand what the project has done to improve their security, by reviewing _how_ they’ve implemented the baseline, and participating in supporting the project to “level up”.  

## As a maintainer of a project why should I bother with a baseline?

If you struggle with understanding what security expectations other projects or adopters may have on open source projects, the baselines are a good starting point to learn about those expectations and why applying those controls mitigates security risks to the project.

Meeting a level of the baseline would also help new contributors and maintainers in ensuring the continued security posture of the project because it captures what has been done and how, making it easier to understand, discover, and continue to apply those to the project.

Additionally, the baseline can serve as a marketing opportunity for a project, to differentiate itself, incentivizing more adoption, and potentially more contributions.

## What OSPS Baseline level am I required to meet?

You are not required to meet any Baseline controls, unless you have a sponsoring organization (e.g. an open source foundation) that imposes a requirement.
However, all projects are encouraged to adhere to Baseline level 1 at minimum because it establishes a “universal security floor” for all open source, capturing many modern-day expectations for software development in an internationally connected ecosystem amidst modern-day online threats (refer to CNCF’s [software supply chain compromise catalog](https://tag-security.cncf.io/community/catalog/compromises/) for examples of successful attacks).
If you are a foundation that has some level or maturity criteria, we recommend you evaluate your lowest criteria tier for security and adjust to match (where reasonable and appropriate) the level 1 baseline.

## How can I prove my project complies with the OSPS Baseline?

Projects can self-attest OSPS Baseline compliance.
As tooling is developed to evaluate projects, this entry will be updated to provide links.

## How can I verify an upstream project’s compliance with the OSPS Baseline?

Many of the OSPS Baseline controls are publicly observable.
However, some of the controls are concerned with privileged settings.
Since the Baseline is designed for the developers of a project, not the consumers, if you need to verify the security posture of an upstream project, you can accept their self-attestation or make an alternate arrangement with the project that meets your specific needs.

## Does OSPS Baseline compliance expire?

OSPS Baseline compliance is a point-in-time status.
We encourage projects using the OSPS Baseline to say something like “As of April 31, 2025, this project complies with OSPS Baseline version 2025-02-30 level 2.”

## What is in scope for OSPS Baseline?

OSPS Baseline seeks to address security hygiene elements — those which secure the ways of working, delivering the product, and equipping its users to adopt it safely.

To use an analogy, Baseline is similar to health department guidelines that require a food processing plant to have practices that ensure safe food products: wearing hair nets, regularly cleaning and sanitizing equipment, monitoring refrigerator temperatures, handling recalls, etc.
	
## How can I get involved in the OSPS Baseline project?
The OSPS Baseline project welcomes contributions in the [GitHub repository](https://github.com/ossf/security-baseline/pull/24/files).
For discussion, join us in [#sig-security-baseline](https://openssf.slack.com/archives/C07DC6TT2QY) in the OpenSSF Slack instance.
