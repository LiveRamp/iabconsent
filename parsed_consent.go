/*

Package iabconsent provides structs and methods for parsing
Vendor Consent Strings as defined by the IAB Consent String 1.1 Spec.
More info on the spec here:
https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#vendor-consent-string-format-.

Copyright (c) 2018 Andy Day. All rights reserved.

Written by Andy Day, Software Engineer @ LiveRamp
for use in the LiveRamp Pixel Server.

*/
package iabconsent

import (
	"encoding/base64"
	"fmt"
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
	SingleOrRangeOffset     = 186
	SingleVendorIdOffset    = 187
	SingleVendorIdSize      = 16
	StartVendorIdOffset     = 187
	StartVendorIdSize       = 16
	EndVendorIdOffset       = 203
	EndVendorIdSize         = 16
)

// ParsedConsent contains all fields defined in the
// IAB Consent String 1.1 Spec.
type ParsedConsent struct {
	ConsentString     string
	Version           int
	Created           time.Time
	LastUpdated       time.Time
	CmpID             int
	CmpVersion        int
	ConsentScreen     int
	ConsentLanguage   string
	VendorListVersion int
	PurposesAllowed   map[int]interface{}
	MaxVendorID       int
	IsRange           bool
	ApprovedVendorIDs map[int]interface{}
	RangeEntry        *RangeEntry
}

// RangeEntry contains all fields in the Range Entry
// portion of the Vendor Consent String. This portion
// of the consent string is only populated when the
// EncodingType field is set to 1.
type RangeEntry struct {
	DefaultConsent bool
	NumEntries     int
	SingleOrRange  bool
	SingleVendorID int
	StartVendorID  int
	EndVendorID    int
}

// Parse takes a base64 Raw URL Encoded string which represents
// a Vendor Consent String and returns a ParsedConsent with
// it's fields populated with the values stored in the string.
func Parse(s string) (*ParsedConsent, error) {
	var b []byte
	var err error

	b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var cs = parseBytes(b)
	var version, cmpID, cmpVersion, consentScreen, vendorListVersion, maxVendorID,
		numEntries, singleVendorID, startVendorID, endVendorID int
	var created, updated int64
	var isRange, defaultConsent, singleOrRange bool
	var consentLanguage string
	var purposesAllowed = make(map[int]interface{})
	var approvedVendorIDs = make(map[int]interface{})

	version, err = cs.parseInt(VersionBitOffset, VersionBitSize)
	if err != nil {
		return nil, err
	}
	created, err = cs.parseInt64(CreatedBitOffset, CreatedBitSize)
	if err != nil {
		return nil, err
	}
	updated, err = cs.parseInt64(UpdatedBitOffset, UpdatedBitSize)
	if err != nil {
		return nil, err
	}
	fmt.Println(created, updated)
	cmpID, err = cs.parseInt(CmpIdOffset, CmpIdSize)
	if err != nil {
		return nil, err
	}
	cmpVersion, err = cs.parseInt(CmpVersionOffset, CmpVersionSize)
	if err != nil {
		return nil, err
	}
	consentScreen, err = cs.parseInt(ConsentScreenSizeOffset, ConsentScreenSize)
	if err != nil {
		return nil, err
	}
	consentLanguage, err = cs.parseString(ConsentLanguageOffset, ConsentLanguageSize)
	if err != nil {
		return nil, err
	}
	vendorListVersion, err = cs.parseInt(VendorListVersionOffset, VendorListVersionSize)
	if err != nil {
		return nil, err
	}
	purposesAllowed = cs.parseBitList(PurposesOffset, PurposesSize)
	maxVendorID, err = cs.parseInt(MaxVendorIdOffset, MaxVendorIdSize)
	if err != nil {
		return nil, err
	}
	isRange = cs.parseBit(EncodingTypeOffset)

	if isRange {
		defaultConsent = cs.parseBit(DefaultConsentOffset)
		numEntries, err = cs.parseInt(NumEntriesOffset, NumEntriesSize)
		if err != nil {
			return nil, err
		}
		singleOrRange = cs.parseBit(SingleOrRangeOffset)

		if singleOrRange {
			singleVendorID, err = cs.parseInt(SingleVendorIdOffset, SingleVendorIdSize)
			if err != nil {
				return nil, err
			}
		} else {
			startVendorID, err = cs.parseInt(StartVendorIdOffset, StartVendorIdSize)
			if err != nil {
				return nil, err
			}
			endVendorID, err = cs.parseInt(EndVendorIdOffset, EndVendorIdSize)
			if err != nil {
				return nil, err
			}
		}
	} else {
		approvedVendorIDs = cs.parseBitList(VendorBitFieldOffset, len(cs.value)-1-VendorBitFieldOffset)
	}

	return &ParsedConsent{
		ConsentString:     cs.value,
		Version:           version,
		Created:           time.Unix(created/10, created%10),
		LastUpdated:       time.Unix(updated/10, updated%10),
		CmpID:             cmpID,
		CmpVersion:        cmpVersion,
		ConsentScreen:     consentScreen,
		ConsentLanguage:   consentLanguage,
		VendorListVersion: vendorListVersion,
		PurposesAllowed:   purposesAllowed,
		MaxVendorID:       maxVendorID,
		IsRange:           isRange,
		ApprovedVendorIDs: approvedVendorIDs,
		RangeEntry: &RangeEntry{
			DefaultConsent: defaultConsent,
			NumEntries:     numEntries,
			SingleOrRange:  singleOrRange,
			SingleVendorID: singleVendorID,
			StartVendorID:  startVendorID,
			EndVendorID:    endVendorID,
		},
	}, nil
}
