package iabconsent

// mspaParsedConsent represents data extract from a Multi-State Privacy Agreement (mspa) consent string.
type MspaParsedConsent struct {
	// The version of this section specification used to encode the string.
	Version int
	// Notice of the Sharing of the Consumer’s Personal Data with Third Parties.
	// 0 Not Applicable. The Business does not share Personal Data with Third Parties.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SharingNotice int
	// Notice of the Opportunity to Opt Out of the Sale of the Consumer’s Personal Data.
	// 0 Not Applicable. The Business does not Sell Personal Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SaleOptOutNotice int
	// Notice of the Opportunity to Opt Out of the Sharing of the Consumer’s Personal Data.
	// 0 Not Applicable.The Business does not Share Personal Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SharingOptOutNotice int
	// Notice of the Opportunity to Opt Out of Processing of the Consumer’s Personal Data for Targeted Advertising.
	// 0 Not Applicable.The Business does not Process Personal Data for Targeted Advertising.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	TargetedAdvertisingOptOutNotice int
	// Notice of the Opportunity to Opt Out of the Processing of the Consumer’s Sensitive Data.
	// 0 Not Applicable. The Business does not Process Sensitive Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SensitiveDataProcessingOptOutNotice int
	// Notice of the Opportunity to Limit Use or Disclosure of the Consumer’s Sensitive Data.
	// 0 Not Applicable. The Business does not use or disclose Sensitive Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SensitiveDataLimitUseNotice int
	// Opt-Out of the Sale of the Consumer’s Personal Data.
	// 0 Not Applicable. SaleOptOutNotice value was not applicable or no notice was provided
	// 1 Opted Out
	// 2 Did Not Opt Out
	SaleOptOut int
	// Opt-Out of the Sharing of the Consumer’s Personal Data.
	// 0 Not Applicable. SharingOptOutNotice value was not applicable or no notice was provided.
	// 1 Opted Out
	// 2 Did Not Opt Out
	SharingOptOut int
	// Opt-Out of Processing the Consumer’s Personal Data for Targeted Advertising.
	// 0 Not Applicable. TargetedAdvertisingOptOutNotice value was not applicable or no notice was provided
	// 1 Opted Out
	// 2 Did Not Opt Out
	TargetedAdvertisingOptOut int
	// Two bits for each Data Activity:
	// 0 Not Applicable. The Business does not Process the specific category of Sensitive Data.
	// 1 No Consent
	// 2 Consent
	// Data Activities:
	// (1) Consent to Process the Consumer’s Sensitive Data Consisting of Personal Data Revealing Racial or Ethnic Origin.
	// (2) Consent to Process the Consumer’s Sensitive Data Consisting of Personal Data Revealing Religious or Philosophical Beliefs.
	// (3) Consent to Process the Consumer’s Sensitive Data Consisting of Personal Data Concerning a Consumer’s Health (including a Mental or Physical Health Condition or Diagnosis; Medical History; or Medical Treatment or Diagnosis by a Health Care Professional).
	// (4) Consent to Process the Consumer’s Sensitive Data Consisting of Personal Data Revealing Sex Life or Sexual Orientation.
	// (5) Consent to Process the Consumer’s Sensitive Data Consisting of Personal Data Revealing Citizenship or Immigration Status.
	// (6) Consent to Process the Consumer’s Sensitive Data Consisting of Genetic Data for the Purpose of Uniquely Identifying an Individual / Natural Person.
	// (7) Consent to Process the Consumer’s Sensitive Data Consisting of Biometric Data for the Purpose of Uniquely Identifying an Individual / Natural Person.
	// (8) Consent to Process the Consumer’s Sensitive Data Consisting of Precise Geolocation Data.
	// (9) Consent to Process the Consumer’s Sensitive Data Consisting of a Consumer’s Social Security, Driver’s License, State Identification Card, or Passport Number.
	// (10) Consent to Process the Consumer’s Sensitive Data Consisting of a Consumer’s Account Log-In, Financial Account, Debit Card, or Credit Card Number in Combination with Any Required Security or Access Code, Password, or Credentials Allowing Access to an Account.
	// (11) Consent to Process the Consumer’s Sensitive Data Consisting of Union Membership.
	// (12) Consent to Process the Consumer’s Sensitive Data Consisting of the contents of a Consumer’s Mail, Email, and Text Messages unless You Are the Intended Recipient of the Communication.
	SensitiveDataProcessing map[int]int
	// Two bits for each Data Activity:
	// 0 Not Applicable. The Business does not have actual knowledge that it Processes Personal Data or Sensitive Data of a Consumer who is a known child.
	// 1 No Consent
	// 2 Consent
	// Fields:
	// (1) Consent to Process the Consumer’s Personal Data or Sensitive Data for Consumers from Age 13 to 16.
	// (2) Consent to Process the Consumer’s Personal Data or Sensitive Data for Consumers Younger Than 13 Years of Age.
	KnownChildSensitiveDataConsents map[int]int
	// Consent to Collection, Use, Retention, Sale, and/or Sharing of the Consumer’s Personal Data that Is Unrelated to or Incompatible with the Purpose(s) for which the Consumer’s Personal Data Was Collected or Processed.
	// 0 Not Applicable. The Business does not use, retain, Sell, or Share the Consumer’s Personal Data for advertising purposes that are unrelated to or incompatible with the purpose(s) for which the Consumer’s Personal Data was collected or processed.
	// 1 No Consent
	// 2 Consent
	PersonalDataConsents int
	// Publisher or Advertiser, as applicable, is a signatory to the IAB Multistate Service Provider Agreement (MSPA), as may be amended from time to time, and declares that the transaction is a “Covered Transaction” as defined in the MSPA.
	// 1 Yes
	// 2 No
	MspaCoveredTransaction int
	// Publisher or Advertiser, as applicable, has enabled “Opt-Out Option Mode” for the “Covered Transaction,” as such terms are defined in the MSPA.
	// 0 Not Applicable.
	// 1 Yes
	// 2 No
	MspaOptOutOptionMode int
	// Publisher or Advertiser, as applicable, has enabled “Service Provider Mode” for the “Covered Transaction,” as such terms are defined in the MSPA.
	// 0 Not Applicable
	// 1 Yes
	// 2 No
	MspaServiceProviderMode int
}

type MspaNotice int

const (
	NoticeNotApplicable MspaNotice = iota
	NoticeProvided
	NoticeNotProvided
)

type MspaOptout int

const (
	OptOutNotApplicable MspaOptout = iota
	OptedOut
	NotOptedOut
)

type MspaConsent int

const (
	ConsentNotApplicable MspaConsent = iota
	NoConsent
	Consent
)
