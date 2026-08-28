# OSPS Baseline release notes

<!-- Use the template below to start the release notes for the next version. -->

<!--
## TEMPLATE

### New controls

* Level 1
    * (none)
* Level 2
    * (none)
* Level 3
    * (none)

### Modified controls

(none)

### Removed controls

(none)

### New control mappings

(none)

### Other changes

-->

## 2026-08-28

### New controls

* Level 1
    * (none)
* Level 2
    * (none)
* Level 3
    * (none)

### Modified controls

* [OSPS-LE-03.01](versions/2026-08-28#osps-le-0301) now also accepts a `LICENSES/` directory as a license location

### Removed controls

(none)

### New control mappings

* [OpenSSF Scorecard](versions/2026-08-28#external-frameworks)

### Other changes

* The control mappings are now maintained as machine-readable [Gemara](https://gemara.openssf.org) mapping documents, published to [grc.store](https://grc.store/openssf) with each release alongside the compiled control catalog
* Each version now includes a generated External Framework Crosswalk page (external requirement → OSPS controls), and the External Frameworks table links each framework's mapping document
* The External Frameworks table now lists only frameworks with published mappings; the OpenChain and OpenCRE entries were renamed to ISO/IEC 18974 and OpenCRE respectively
* Baseline content and tooling migrated to the Gemara v1 schema

## 2026-02-19

### New controls

* Level 1
    * [OSPS-BR-01.03](versions/2026-02-19#osps-br-0103)
* Level 2
    * [OSPS-DO-07.01](versions/2026-02-19#osps-do-0701)
* Level 3
    * [OSPS-BR-01.04](versions/2026-02-19#osps-br-0104)


### Modified controls

* [OSPS-BR-01.01](versions/2026-02-19#osps-br-0101) updated to remove ambiguity and overlap
* [OSPS-BR-03.02](versions/2026-02-19#osps-br-0302) updated to clarify intent
* [OSPS-QA-04.01](versions/2026-02-19#osps-qa-0401) updated to clarify intent

### Removed controls

* OSPS-BR-01.02

### New control mappings

* [BSI-TR-03185-2](https://www.bsi.bund.de/SharedDocs/Downloads/EN/BSI/Publications/TechGuidelines/TR03185/BSI-TR-03185-2.pdf?__blob=publicationFile&v=5)
* Additional [UKSSCOP](https://www.gov.uk/government/publications/software-security-code-of-practice/software-security-code-of-practice) 

### Other changes

* Control titles are more brief to improve clarity


## 2025-10-10

### New controls

* Level 1
    * [OSPS-BR-01.02](versions/2025-10-10#osps-br-0102)
    * [OSPS-BR-03.02](versions/2025-10-10#osps-br-0302)
    * [OSPS-BR-07.01](versions/2025-10-10#osps-br-0701)
    * [OSPS-QA-05.02](versions/2025-10-10#osps-qa-0502)
* Level 2
    * (none)
* Level 3
    * [OSPS-BR-07.02](versions/2025-10-10#osps-br-0702)
    * [OSPS-DO-03.02](versions/2025-10-10#osps-br-0302)

### Modified controls

(none)

### Removed controls

(none)

### New control mappings

* [800-161](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-161r1-upd1.pdf)
* [PCIDSS](https://docs-prv.pcisecuritystandards.org/PCI%20DSS/Standard/PCI-DSS-v4_0_1.pdf)
* [PSSCRM](https://arxiv.org/pdf/2404.12300)
* [SAMM](https://owaspsamm.org/model/)
* [UKSSCOP](https://www.gov.uk/government/publications/software-security-code-of-practice/software-security-code-of-practice)

### Other changes

* Added definitions for several additional terms in order to improve clarity

----

## 2025-02-25

Initial release of the Open Source Project Security Baseline.
