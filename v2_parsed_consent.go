package iabconsent

import (
	"time"
)

// V2ParsedConsent represents data extracted from an v2 TCF Consent String.
type V2ParsedConsent struct {
	// Version number of the encoding format.
	Version int
	// Epoch deciseconds when this TC String was first created (should not be changed
	// unless a new TCString is created from scratch).
	Created time.Time
	// Epoch deciseconds when TC String was last updated (Must be updated any time a
	// value is changed).
	LastUpdated time.Time
	// Consent Management Platform ID that last updated the TC String.
	// A unique ID will be assigned to each Consent Management Platform.
	CMPID int
	// Consent Management Platform version of the CMP that last updated this TC String.
	// Each change to a CMP should increment their internally assigned version number as
	// a record of which version the user gave consent and transparency was established.
	CMPVersion int
	// CMP Screen number at which consent was given for a user with the CMP that last
	// updated this TC String. The number is a CMP internal designation and is CMPVersion
	// specific. The number is used for identifying on which screen a user gave consent
	// as a record.
	ConsentScreen int
	// Two-letter ISO 639-1 language code in which the CMP UI was presented.
	ConsentLanguage string
	// Number corresponds to the Global Vendor List (GVL) vendorListVersion.
	VendorListVersion int
	// Version of policy used within GVL.
	TCFPolicyVersion int
	// Whether the signals encoded in this TC String were from service-specific storage
	// (true) versus ‘global’ consensu.org shared storage (false).
	IsServiceSpecific bool
	// Setting this to 1 means that a publisher-run CMP – that is still IAB Europe
	// registered – is using customized Stack descriptions and not the standard stack
	// descriptions defined in the Policies. A CMP that services multiple publishers sets
	// this value to 0.
	UseNonStandardStacks bool
	// The TCF Policies designates certain Features as “special” which means a CMP must
	// afford the user a means to opt in to their use. These “Special Features” are
	// published and numerically identified in the Global Vendor List separately from
	// normal Features.
	SpecialFeaturesOptIn map[int]bool
	// The user’s consent value for each Purpose established on the legal basis of consent.
	// The Purposes are numerically identified and published in the Global Vendor List.
	// From left to right, Purpose 1 maps to the 0th bit, purpose 24 maps to the bit at
	// index 23. Special Purposes are a different ID space and not included in this field.
	PurposesConsent map[int]bool
	// The Purpose’s transparency requirements are met for each Purpose on the legal basis
	// of legitimate interest and the user has not exercised their “Right to Object” to that
	// Purpose. By default or if the user has exercised their “Right to Object” to a Purpose,
	// the corresponding bit for that Purpose is set to 0. From left to right, Purpose 1 maps
	// to the 0th bit, purpose 24 maps to the bit at index 23. Special Purposes are a
	// different ID space and not included in this field.
	PurposesLITransparency map[int]bool
	// CMPs can use the PublisherCC field to indicate the legal jurisdiction the publisher is
	// under to help vendors determine whether the vendor needs consent for Purpose 1. In a
	// globally-scoped TC string, this field must always have a value of 0. When a CMP
	// encounters a globally-scoped TC String with PurposeOneTreatment=1 then it is considered
	// invalid and the CMP must discard it and re-establish transparency and consent.
	PurposeOneTreatment bool
	// The country code of the country that determines legislation of reference. Commonly,
	// this corresponds to the country in which the publisher’s business entity is established.
	PublisherCC string

	// The maximum Vendor ID that is represented in the following bit field or range encoding.
	MaxConsentVendorID int
	// The encoding scheme used to encode the IDs in the section – Either a BitField Section or
	// Range Section follows. Encoding logic should choose the encoding scheme that results in
	// the smaller output size for a given set.
	IsConsentRangeEncoding bool
	// The consent value for each Vendor ID.
	ConsentedVendors map[int]bool
	// Number of RangeEntry sections to follow.
	NumConsentEntries int
	// A single or range of Vendor ID(s) who have received consent. If a Vendor ID is not within
	// the bounds of the ranges then the vendor is assumed to have “No Consent”.
	ConsentedVendorsRange []*RangeEntry

	// The maximum Vendor ID that is represented in the following bit field or range encoding.
	MaxInterestsVendorID int
	// The encoding scheme used to encode the IDs in the section – Either a BitField Section or
	// Range Section follows. Encoding logic should encode with the encoding scheme that results
	// in the smaller output size for a given set.
	IsInterestsRangeEncoding bool
	// The legitimate interest value for each Vendor ID from 1 to MaxVendorId. Set the bit
	// corresponding to a given vendor to 1 if the CMP has established transparency for a vendor's
	// legitimate interest disclosures. If a user exercises their “Right To Object” to a vendor’s
	// processing based on a legitimate interest, then that vendor’s bit must be set to 0.
	InterestsVendors map[int]bool
	// 	Number of RangeEntry sections to follow.
	NumInterestsEntries int
	// A single or range of Vendor ID(s) who have established transparency for their legitimate
	// interest disclosures with the user. If a Vendor ID is not within the bounds of the ranges
	// then they have not established that transparency.
	InterestsVendorsRange []*RangeEntry

	// Number of restriction records to follow.
	NumPubRestrictions int
	// Each Publisher Restriction Entry is made up of three parts: Purpose ID, Restriction Type and,
	// List of Vendor IDs under that Purpose restriction.
	PubRestrictionEntries []*PubRestrictionEntry

	// The DisclosedVendors is a TC String segment that signals which vendors have been disclosed
	// to a given user by a CMP.
	OOBDisclosedVendors *OOBVendorList
	// Signals which vendors the publisher permits to use OOB legal bases.
	OOBAllowedVendors *OOBVendorList
	// Publishers may need to establish transparency and consent for a set of personal data processing
	// purposes for their own use. For example, a publisher that wants to set a frequency-capping
	// first-party cookie should request user consent for Purpose 1 "Store and/or access information on
	// a device" in jurisdictions where it is required.
	//
	// The Publisher TC segment in the TC string represents publisher purposes transparency & consent
	// signals which is different than the other TC String segments; they are used to collect consumer
	// purposes transparency & consent for vendors. This segment supports the standard list of purposes
	// defined by the TCF as well as Custom Purposes defined by the publisher if they so choose.
	*PublisherTCEntry
}

// RestrictionType is an enum type of publisher restriction types.
type RestrictionType int

const (
	// Purpose Flatly Not Allowed by Publisher (regardless of Vendor declarations).
	PurposeFlatlyNotAllowed RestrictionType = iota
	// Require Consent (if Vendor has declared the Purpose IDs legal basis as Legitimate
	// Interest and flexible)
	RequireConsent
	// Require Legitimate Interest (if Vendor has declared the Purpose IDs legal basis as Consent and flexible).
	RequireLegitimateInterest
	// Undefined (not used).
	Undefined
)

// PubRestrictionEntry is made up of three parts: Purpose ID, Restriction Type and, List
// of Vendor IDs under that Purpose restriction.
type PubRestrictionEntry struct {
	// The Vendor’s declared Purpose ID that the publisher has indicated that they are overriding.
	PurposeID int
	// The restriction type.
	RestrictionType RestrictionType
	// Number of RangeEntry sections to follow.
	NumEntries int
	// A single or range of Vendor ID(s) who the publisher has designated as restricted under the
	// Purpose ID in this PubRestrictionsEntry.
	RestrictionsRange []*RangeEntry
}

// PublisherTCEntry represents Publisher Purposes Transparency and Consent.
type PublisherTCEntry struct {
	// Enum type
	SegmentType SegmentType
	// The user's consent value for each Purpose established on the legal basis of consent, for the publisher.
	// The Purposes are numerically identified and published in the Global Vendor List.
	PubPurposesConsent map[int]bool
	// The Purpose’s transparency requirements are met for each Purpose established on the legal basis of legitimate
	// interest and the user has not exercised their “Right to Object” to that Purpose. By default or if the user has
	// exercised their “Right to Object to a Purpose, the corresponding bit for that purpose is set to 0.
	PubPurposesLITransparency map[int]bool
	// The number of Custom Purposes.
	NumCustomPurposes int
	// The consent value for each CustomPurposeId from 1 to NumberCustomPurposes,
	CustomPurposesConsent map[int]bool
	// The legitimate Interest disclosure establishment value for each CustomPurposeId from 1 to NumberCustomPurposes.
	CustomPurposesLITransparency map[int]bool
}

// SegmentType is an enum type of possible Out-of-Band (OOB) legal bases.
type SegmentType int

const (
	// The core string.
	CoreString SegmentType = iota
	// The DisclosedVendors is a TC String segment that signals which vendors have been disclosed to a given user
	// by a CMP. This segment is required when saving a global-context TC String. When a CMP updates a globally-scoped
	// TC String, the CMP MUST retain the existing values and only add new disclosed Vendor IDs that had not been added
	// by other CMPs in prior interactions with this user.
	DisclosedVendors
	// Signals which vendors the publisher permits to use OOB legal bases.
	AllowedVendors
	// Publishers may need to establish transparency and consent for a set of personal data processing
	// purposes for their own use. For example, a publisher that wants to set a frequency-capping first-party
	// cookie should request user consent for Purpose 1 "Store and/or access information on a device" in
	// jurisdictions where it is required.
	//
	// The Publisher TC segment in the TC string represents publisher purposes transparency & consent signals
	// which is different than the other TC String segments; they are used to collect consumer purposes transparency
	// & consent for vendors. This segment supports the standard list of purposes defined by the TCF as well as
	// Custom Purposes defined by the publisher if they so choose.
	PublisherTC
)

// OOBVendorList is represents either a DisclosedVendors or AllowedVendors list.
type OOBVendorList struct {
	// Enum type.
	SegmentType SegmentType
	// The maximum Vendor ID that is included.
	MaxVendorID int
	// The encoding scheme used to encode the IDs in the section – Either a BitField Section or Range Section follows.
	// Encoding logic should choose the encoding scheme that results in the smaller output size for a given set.
	IsRangeEncoding bool

	// The value for each Vendor ID from 1 to MaxVendorId. Set the bit corresponding to a given Vendor ID to 1 if the
	// Publisher permits the vendor to use OOB legal bases.
	Vendors map[int]bool

	// Number of RangeEntry sections to follow.
	NumEntries int
	// A single or range of Vendor ID(s) of Vendor(s) who are allowed to use OOB legal bases on the given publisher’s
	// digital property. If a Vendor ID is not within the bounds of the ranges then they are not allowed to use OOB
	// legal bases on the given publisher's digital property.
	VendorEntries []*RangeEntry
}

// SpecialFeature is an enum type for special features. The TCF Policies designates certain Features as “special” which
// means a CMP must afford the user a means to opt in to their use. These “Special Features” are published and
// numerically identified in the Global Vendor List separately from normal Features.
type SpecialFeature int

const (
	InvalidSpecialFeature SpecialFeature = iota
	// Vendors can:
	// - Collect and process precise geolocation data in support of one or more purposes.
	// - N.B. Precise geolocation means that there are no restrictions on the precision of a user’s location;
	//   this can be accurate to within several meters.
	UsePreciseGeolocation
	// Vendors can:
	// - Create an identifier using data collected via actively scanning a device for specific characteristics, e.g.
	//   installed fonts or screen resolution.
	// - Use such an identifier to re-identify a device.
	ActivelyScanDevice
)

// SpecialPurpose is an enum type for special purposes.
type SpecialPurpose int

const (
	InvalidSpecialPurpose SpecialPurpose = iota
	// To ensure security, prevent fraud and debug vendors can:
	// - Ensure data are securely transmitted
	// - Detect and prevent malicious, fraudulent, invalid, or illegal activity.
	// - Ensure correct and efficient operation of systems and processes, including to monitor and enhance the
	//    performance of systems and processes engaged in permitted purposes.
	// Vendors cannot:
	// - Conduct any other data processing operation allowed under a different purpose under this purpose
	EnsureSecurity
	// To deliver information and respond to technical requests vendors can:
	// - Use a user’s IP address to deliver an ad over the internet
	// - Respond to a user’s interaction with an ad by sending the user to a landing page
	// - Use a user’s IP address to deliver content over the internet
	// - Respond to a user’s interaction with content by sending the user to a landing page
	// - Use information about the device type and capabilities for delivering ads or content, for example, to deliver
	// the right size ad creative or video file in a format supported by the device
	// Vendors cannot:
	// - Conduct any other data processing operation allowed under a different purpose under this purpose
	TechnicallyDeliverAds
)

// EveryPurposeAllowed returns true iff every purpose number in ps exists in
// the V2ParsedConsent, otherwise false. This explicitly checks that
// the vendor has opted in, and does not cover legitimate interest.
// This is vendor agnostic, and should not be used without checking if
// there are any Publisher Restrictions for a given vendor or vendors
// (which can be done with a call of p.PublisherRestricted).
func (p *V2ParsedConsent) EveryPurposeAllowed(ps []int) bool {
	for _, rp := range ps {
		if !p.PurposesConsent[rp] {
			return false
		}
	}
	return true
}

// PurposeAllowed returns true if the passed purpose number exists in
// the V2ParsedConsent, otherwise false.
func (p *V2ParsedConsent) PurposeAllowed(ps int) bool {
	if !p.PurposesConsent[ps] {
		return false
	}
	return true
}

// VendorAllowed returns true if the ParsedConsent contains affirmative consent
// for VendorID |v|.
func (p *V2ParsedConsent) VendorAllowed(v int) bool {
	if p.IsConsentRangeEncoding {
		return inRangeEntries(v, p.ConsentedVendorsRange)
	}

	return p.ConsentedVendors[v]
}

// PublisherRestricted returns true if any purpose in |ps| is
// Flatly Not Allowed and |v| is covered by that restriction.
func (p *V2ParsedConsent) PublisherRestricted(ps []int, v int) bool {
	// Map-ify ps for use in checking pub restrictions.
	var pm = make(map[int]bool)
	for _, p := range ps {
		pm[p] = true
	}

	if p.NumPubRestrictions > 0 {
		for _, re := range p.PubRestrictionEntries {
			if pm[re.PurposeID] &&
				re.RestrictionType == PurposeFlatlyNotAllowed &&
				inRangeEntries(v, re.RestrictionsRange) {

				return true
			}
		}
	}
	return false
}

// inRangeEntries returns whether |v| is found within |entries|.
func inRangeEntries(v int, entries []*RangeEntry) bool {
	for _, re := range entries {
		if re.StartVendorID <= v && v <= re.EndVendorID {
			return true
		}
	}
	return false
}

// SuitableToProcess evaluates if its suitable for a vendor (with a set of
// required purposes allowed on the basis of consent) to process a given request.
func (p *V2ParsedConsent) SuitableToProcess(ps []int, v int) bool {
	return p.VendorAllowed(v) &&
		p.EveryPurposeAllowed(ps) &&
		!p.PublisherRestricted(ps, v)
}
