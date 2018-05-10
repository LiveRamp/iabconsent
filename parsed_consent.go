/*

Package iabconsent provides structs and methods for parsing
Vendor Consent Strings as defined by the IAB Consent String 1.1 Spec.
More info on the spec here:
https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#vendor-consent-string-format-.

Copyright (c) 2018 LiveRamp. All rights reserved.

Written by Andy Day, Software Engineer @ LiveRamp
for use in the LiveRamp Pixel Server.

*/
package iabconsent

import (
	"encoding/base64"
	"time"
)

// These constants represent the bit offsets and sizes of the
// fields in the IAB Consent String 1.1 Spec.
const (
	VersionBitOffset        = 0
	VersionBitSize          = 6
	CreatedBitOffset        = 6
	CreatedBitSize          = 36
	UpdatedBitOffset        = 42
	UpdatedBitSize          = 36
	CmpIdOffset             = 78
	CmpIdSize               = 12
	CmpVersionOffset        = 90
	CmpVersionSize          = 12
	ConsentScreenSizeOffset = 102
	ConsentScreenSize       = 6
	ConsentLanguageOffset   = 108
	ConsentLanguageSize     = 12
	VendorListVersionOffset = 120
	VendorListVersionSize   = 12
	PurposesOffset          = 132
	PurposesSize            = 24
	MaxVendorIdOffset       = 156
	MaxVendorIdSize         = 16
	EncodingTypeOffset      = 172
	VendorBitFieldOffset    = 173
	DefaultConsentOffset    = 173
	NumEntriesOffset        = 174
	NumEntriesSize          = 12
	RangeEntryOffset        = 186
	VendorIdSize            = 16
)

// ParsedConsent contains all fields defined in the
// IAB Consent String 1.1 Spec.
type ParsedConsent struct {
	consentString     string
	version           int
	created           time.Time
	lastUpdated       time.Time
	cmpID             int
	cmpVersion        int
	consentScreen     int
	consentLanguage   string
	vendorListVersion int
	purposesAllowed   map[int]bool
	maxVendorID       int
	isRange           bool
	approvedVendorIDs map[int]bool
	defaultConsent    bool
	numEntries        int
	rangeEntries      []*rangeEntry
}

// ConsentString returns the consentString.
func (p *ParsedConsent) ConsentString() string {
	return p.consentString
}

// Version returns the version.
func (p *ParsedConsent) Version() int {
	return p.version
}

// Created returns the created.
func (p *ParsedConsent) Created() time.Time {
	return p.created
}

// LastUpdated returns lastUpdated.
func (p *ParsedConsent) LastUpdated() time.Time {
	return p.lastUpdated
}

// CmpID returns cmpID.
func (p *ParsedConsent) CmpID() int {
	return p.cmpID
}

// CmpVersion returns cmpVersion.
func (p *ParsedConsent) CmpVersion() int {
	return p.cmpVersion
}

// ConsentScreen returns consentScreen.
func (p *ParsedConsent) ConsentScreen() int {
	return p.consentScreen
}

// ConsentLanguage returns consentLanguage.
func (p *ParsedConsent) ConsentLanguage() string {
	return p.consentLanguage
}

// VendorListVersion returns vendorListVersion.
func (p *ParsedConsent) VendorListVersion() int {
	return p.vendorListVersion
}

// MaxVendorID returns maxVendorID.
func (p *ParsedConsent) MaxVendorID() int {
	return p.maxVendorID
}

// PurposeAllowed returns true if the consent ID |pu|
// exists in the ParsedConsent and false otherwise.
func (p *ParsedConsent) PurposeAllowed(pu int) bool {
	_, ok := p.purposesAllowed[pu]
	return ok
}

// PurposesAllowed returns true if each consent ID in |pu|
// exists in the ParsedConsent, and false if any does not.
func (p *ParsedConsent) PurposesAllowed(pu []int) bool {
	for _, rp := range pu {
		if !p.PurposeAllowed(rp) {
			return false
		}
	}
	return true
}

// VendorAllowed returns true if the ParsedConsent contains
// affirmative consent for Vendor of ID |i|.
func (p *ParsedConsent) VendorAllowed(i int) bool {
	if p.isRange {
		// defaultConsent indicates the consent for those
		// not covered by any Range Entries. Vendors covered
		// in rangeEntries have the opposite consent of
		// defaultConsent.
		for _, re := range p.rangeEntries {
			if re.StartVendorID <= i &&
				re.EndVendorID >= i {
				return !p.defaultConsent
			}
		}
	} else {
		var _, ok = p.approvedVendorIDs[i]
		return ok
	}
	return p.defaultConsent
}

// rangeEntry contains all fields in the Range Entry
// portion of the Vendor Consent String. This portion
// of the consent string is only populated when the
// EncodingType field is set to 1.
type rangeEntry struct {
	StartVendorID int
	EndVendorID   int
}

// Parse takes a base64 Raw URL Encoded string which represents
// a Vendor Consent String and returns a ParsedConsent with
// it's fields populated with the values stored in the string.
// Example Usage:
//	var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
func Parse(s string) (*ParsedConsent, error) {
	var b []byte
	var err error

	b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var bs = parseBytes(b)
	var version, cmpID, cmpVersion, consentScreen, vendorListVersion, maxVendorID, numEntries int
	var created, updated time.Time
	var isRangeEntries, defaultConsent, isIDRange bool
	var consentLanguage string
	var purposesAllowed = make(map[int]bool)
	var approvedVendorIDs = make(map[int]bool)

	version, err = bs.parseInt(VersionBitOffset, VersionBitSize)
	if err != nil {
		return nil, err
	}
	created, err = bs.parseTime(CreatedBitOffset, CreatedBitSize)
	if err != nil {
		return nil, err
	}
	updated, err = bs.parseTime(UpdatedBitOffset, UpdatedBitSize)
	if err != nil {
		return nil, err
	}
	cmpID, err = bs.parseInt(CmpIdOffset, CmpIdSize)
	if err != nil {
		return nil, err
	}
	cmpVersion, err = bs.parseInt(CmpVersionOffset, CmpVersionSize)
	if err != nil {
		return nil, err
	}
	consentScreen, err = bs.parseInt(ConsentScreenSizeOffset, ConsentScreenSize)
	if err != nil {
		return nil, err
	}
	consentLanguage, err = bs.parseString(ConsentLanguageOffset, ConsentLanguageSize)
	if err != nil {
		return nil, err
	}
	vendorListVersion, err = bs.parseInt(VendorListVersionOffset, VendorListVersionSize)
	if err != nil {
		return nil, err
	}
	purposesAllowed, err = bs.parseBitList(PurposesOffset, PurposesSize)
	if err != nil {
		return nil, err
	}
	maxVendorID, err = bs.parseInt(MaxVendorIdOffset, MaxVendorIdSize)
	if err != nil {
		return nil, err
	}
	isRangeEntries, err = bs.parseBool(EncodingTypeOffset)
	if err != nil {
		return nil, err
	}

	var rangeEntries []*rangeEntry

	if isRangeEntries {
		defaultConsent, err = bs.parseBool(DefaultConsentOffset)
		if err != nil {
			return nil, err
		}
		numEntries, err = bs.parseInt(NumEntriesOffset, NumEntriesSize)
		if err != nil {
			return nil, err
		}

		// Track how many range entry bits we've parsed since it's variable.
		var parsedBits = 0

		for i := 0; i < numEntries; i++ {
			var startVendorID, endVendorID int

			isIDRange, err = bs.parseBool(RangeEntryOffset + parsedBits)
			parsedBits++

			if isIDRange {
				startVendorID, err = bs.parseInt(RangeEntryOffset+parsedBits, VendorIdSize)
				if err != nil {
					return nil, err
				}
				parsedBits += VendorIdSize
				endVendorID, err = bs.parseInt(RangeEntryOffset+parsedBits, VendorIdSize)
				if err != nil {
					return nil, err
				}
				parsedBits += VendorIdSize
			} else {
				startVendorID, err = bs.parseInt(RangeEntryOffset+parsedBits, VendorIdSize)
				if err != nil {
					return nil, err
				}
				endVendorID = startVendorID
				parsedBits += VendorIdSize
			}

			rangeEntries = append(rangeEntries, &rangeEntry{
				StartVendorID: startVendorID,
				EndVendorID:   endVendorID,
			})
		}
	} else {
		approvedVendorIDs, err = bs.parseBitList(VendorBitFieldOffset, maxVendorID)
		if err != nil {
			return nil, err
		}
	}

	return &ParsedConsent{
		consentString:     bs.value,
		version:           version,
		created:           created,
		lastUpdated:       updated,
		cmpID:             cmpID,
		cmpVersion:        cmpVersion,
		consentScreen:     consentScreen,
		consentLanguage:   consentLanguage,
		vendorListVersion: vendorListVersion,
		purposesAllowed:   purposesAllowed,
		maxVendorID:       maxVendorID,
		isRange:           isRangeEntries,
		approvedVendorIDs: approvedVendorIDs,
		defaultConsent:    defaultConsent,
		numEntries:        numEntries,
		rangeEntries:      rangeEntries,
	}, nil
}
