![image](https://github.com/user-attachments/assets/27e12677-ab40-485f-b07f-a830694ef3d3)

---

###  **PHASE 1: Downlink Synchronization & Cell Selection**

#### 1. **Cell Search and Downlink Synchronization**

* UE synchronizes with the cell by:

  * Detecting **Primary and Secondary Synchronization Signals (PSS/SSS)**.
  * Decoding the **Cell ID** and **PBCH** (contains MIB - Master Information Block).
* Purpose:

  * Time and frequency synchronization.
  * Retrieve basic cell parameters.

---

###  **PHASE 2: Uplink Synchronization via RACH**

#### 2. **System Information Acquisition**

* UE decodes **SIB1** for:

  * RACH configuration.
  * UL synchronization info.
  * Access barring info.

#### 3. **RACH Procedure (Msg1 – Msg4)**

##### ➤ **Msg1: RACH Preamble Transmission**

* UE selects a **random preamble** (with **RAP ID**) and sends it on **PRACH**.
* Starts **T300** timer to wait for RRC Setup.

##### ➤ **Msg2: Random Access Response (RAR)**

* gNB detects Msg1 and sends:

  * **DCI Format 1\_0** on PDCCH (CRC scrambled with RA-RNTI).
  * **RAR message** on PDSCH, containing:

    * **Timing Advance (TA)**
    * **UL Grant**
    * **Temporary C-RNTI**

##### ➤ **Msg3: RRC Connection Request**

* UE sends:

  * `RRC Connection Request` on **UL\_CCCH** (using grant from Msg2).
  * Carries:

    * `ue-Identity`: Random number (0 to 2^39−1)
    * `establishmentCause`
* Used for contention resolution later in Msg4.

##### ➤ **Msg4: RRC Connection Setup**

* gNB responds with `RRC Setup` message including:

  * `SRB1` configuration
  * `masterCellGroup`:

    * `cellGroupId`
    * `rlc-BearerToAddModList`
    * `mac-CellGroupConfig`
    * `physicalCellGroupConfig`
* UE stops timer **T300**.

---

###  **PHASE 3: RRC Setup Completion & Registration**

#### 4. **RRC Setup Complete + Registration Request**

* UE sends `RRC Setup Complete` message with:

  * `dedicatedNAS-Message`:

    * **Registration Request**
    * Contains:

      * `selectedPLMN-Identity`
      * `registeredAMF`
      * `snssai-list`
      * `UE network capability`
* gNB selects **AMF** and allocates **RAN UE NGAP ID**.

---

###  **PHASE 4: NGAP Initial Context Creation**

#### 5. **Initial UE Message**

* gNB sends `Initial UE Message` to AMF with:

  * RRC Establishment Cause
  * RAN UE NGAP ID
  * NAS Registration Request message

---

###  **PHASE 5: Identity and Authentication**

#### 6. **UE Identity Transfer** *(conditional)*

* If:

  * SUCI not available, or
  * New AMF context is selected
* AMF sends **Identity Request**
* UE responds with **Identity Response** including:

  * **SUCI** (derived using HPLMN public key)

#### 7. **Authentication & NAS Security**

* AMF initiates **Authentication procedure**

  * Verifies UE legitimacy
* Sends **NAS Security Mode Command**:

  * Selects NAS encryption/integrity algorithms
  * Requests **IMEISV**
* UE responds with:

  * **NAS Security Mode Complete** (with IMEISV)

---

###  **PHASE 6: Initial Context Setup**

#### 8. **Initial Context Setup Request**

* AMF allocates:

  * `AMF UE NGAP ID`
* Sends `INITIAL CONTEXT SETUP REQUEST` to gNB, includes:

  * `Registration Accept` NAS message
  * `UE Aggregate Maximum Bit Rate`
  * `UE security capabilities`
  * Security key
  * **PDU Session setup request(s)**

---

###  **PHASE 7: AS Security and Bearer Setup**

#### 9. **AS (Access Stratum) UE Capability Transfer**

* gNB sends:

  * `UE Capability Enquiry`
* UE responds:

  * `UE Capability Information`
* gNB updates AMF with UE capability

#### 10. **AS Security Setup**

* gNB sends:

  * `Security Mode Command` to UE
* UE:

  * Derives keys
  * Enables **DL encryption**
  * Sends `Security Mode Complete`
  * Then **UL encryption** is enabled

---

###  **PHASE 8: SRB2 & DRB Setup**

#### 11. **RRC Reconfiguration**

* gNB sends `RRC Reconfiguration` to:

  * Setup **SRB2 and DRB**
* UE responds with:

  * `RRC Reconfiguration Complete`
* gNB sends `INITIAL CONTEXT SETUP RESPONSE` to AMF

---

###  **PHASE 9: Registration Complete & PDU Session Setup**

#### 12. **Registration Complete + PDU Session Establishment**

* UE sends:

  * `Registration Complete`
  * `PDU Session Establishment Request` (NAS)

#### 13. **PDU SESSION RESOURCE SETUP REQUEST (AMF → gNB)**

* AMF sends:

  * List of **PDU sessions**
  * QoS Flows and attributes
* gNB maps and activates PDU sessions and QoS flows

#### 14. **NAS: PDU SESSION ESTABLISHMENT ACCEPT**

* gNB sends NAS `PDU Session Establishment Accept` to UE

---

### 📶 **Summary: 5G Registration Call Flow in Sequence**

| **Step** | **Message/Procedure**            | **Key Elements**               |
| -------- | -------------------------------- | ------------------------------ |
| 1        | Cell Search & Synchronization    | PSS, SSS, PBCH (MIB)           |
| 2        | SIB1 Acquisition                 | RACH configs                   |
| 3        | RACH Msg1                        | PRACH preamble                 |
| 4        | RACH Msg2                        | RAR, TA, Temp C-RNTI           |
| 5        | RACH Msg3                        | RRC Connection Request         |
| 6        | RACH Msg4                        | RRC Setup                      |
| 7        | RRC Setup Complete               | + Registration Request         |
| 8        | Initial UE Message               | Sent by gNB to AMF             |
| 9        | Identity Req/Resp                | SUCI                           |
| 10       | Authentication                   | NAS Security Mode Cmd/Complete |
| 11       | Initial Context Setup Request    | From AMF to gNB                |
| 12       | UE Capability Transfer           | Enquiry + Info                 |
| 13       | AS Security Mode Cmd             | + Mode Complete                |
| 14       | RRC Reconfiguration              | SRB2, DRB setup                |
| 15       | RRC Reconfig Complete            | Sent by UE                     |
| 16       | Initial Context Setup Response   | gNB → AMF                      |
| 17       | Registration Complete            | Sent by UE                     |
| 18       | PDU Session Establishment Req    | NAS message                    |
| 19       | PDU Session Resource Setup Req   | AMF → gNB                      |
| 20       | PDU Session Establishment Accept | NAS to UE                      |

---



![image](https://github.com/user-attachments/assets/126456f6-3d8a-4363-9bb9-335113f28c21)

UE ⇄ gNB ⇄ AMF ⇄ SMF

1. Cell Search, PSS/SSS, PBCH (MIB)
2. RACH (Msg1-4)
3. RRC Setup
4. Registration Request
5. Identity (SUCI)
6. Authentication
7. NAS Security
8. Initial Context Setup
9. AS Security
10. RRC Reconfiguration
11. Registration Complete
12. PDU Session Request / Accept
