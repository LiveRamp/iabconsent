package iabconsent

import (
	"encoding/base64"
	"fmt"
	"time"
)

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
	VendorBitfieldOffset    = 173
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

type RangeEntry struct {
	DefaultConsent bool
	NumEntries     int
	SingleOrRange  bool
	SingleVendorID int
	StartVendorID  int
	EndVendorID    int
}

// Parse...
func Parse(s string) (*ParsedConsent, error) {
	var b []byte
	var err error

	b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var cs = ParseBytes(b)
	var version, cmpID, cmpVersion, consentScreen, vendorListVersion, maxVendorID,
		numEntries, singleVendorID, startVendorID, endVendorID int
	var created, updated int64
	var isRange, defaultConsent, singleOrRange bool
	var consentLanguage string
	var purposesAllowed = make(map[int]interface{})
	var approvedVendorIDs = make(map[int]interface{})

	version, err = cs.ParseInt(VersionBitOffset, VersionBitSize)
	if err != nil {
		return nil, err
	}
	created, err = cs.ParseInt64(CreatedBitOffset, CreatedBitSize)
	if err != nil {
		return nil, err
	}
	updated, err = cs.ParseInt64(UpdatedBitOffset, UpdatedBitSize)
	if err != nil {
		return nil, err
	}
	fmt.Println(created, updated)
	cmpID, err = cs.ParseInt(CmpIdOffset, CmpIdSize)
	if err != nil {
		return nil, err
	}
	cmpVersion, err = cs.ParseInt(CmpVersionOffset, CmpVersionSize)
	if err != nil {
		return nil, err
	}
	consentScreen, err = cs.ParseInt(ConsentScreenSizeOffset, ConsentScreenSize)
	if err != nil {
		return nil, err
	}
	consentLanguage, err = cs.ParseString(ConsentLanguageOffset, ConsentLanguageSize)
	if err != nil {
		return nil, err
	}
	vendorListVersion, err = cs.ParseInt(VendorListVersionOffset, VendorListVersionSize)
	if err != nil {
		return nil, err
	}
	purposesAllowed = cs.ParseBitList(PurposesOffset, PurposesSize)
	maxVendorID, err = cs.ParseInt(MaxVendorIdOffset, MaxVendorIdSize)
	if err != nil {
		return nil, err
	}
	isRange = cs.ParseBit(EncodingTypeOffset)

	if isRange {
		defaultConsent = cs.ParseBit(DefaultConsentOffset)
		numEntries, err = cs.ParseInt(NumEntriesOffset, NumEntriesSize)
		if err != nil {
			return nil, err
		}
		singleOrRange = cs.ParseBit(SingleOrRangeOffset)

		if singleOrRange {
			singleVendorID, err = cs.ParseInt(SingleVendorIdOffset, SingleVendorIdSize)
			if err != nil {
				return nil, err
			}
		} else {
			startVendorID, err = cs.ParseInt(StartVendorIdOffset, StartVendorIdSize)
			if err != nil {
				return nil, err
			}
			endVendorID, err = cs.ParseInt(EndVendorIdOffset, EndVendorIdSize)
			if err != nil {
				return nil, err
			}
		}
	} else {
		approvedVendorIDs = cs.ParseBitList(VendorBitfieldOffset, len(cs.value)-1-VendorBitfieldOffset)
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
