package iabconsent

import (
	"time"
)

// ParsedConsent represents data extracted from an IAB Consent String, v1.1.
type ParsedConsent struct {
	Version           int
	Created           time.Time
	LastUpdated       time.Time
	CMPID             int
	CMPVersion        int
	ConsentScreen     int
	ConsentLanguage   string
	VendorListVersion int
	PurposesAllowed   map[int]bool
	MaxVendorID       int
	IsRangeEncoding   bool
	ConsentedVendors  map[int]bool
	DefaultConsent    bool
	NumEntries        int
	RangeEntries      []*RangeEntry
}

// EveryPurposeAllowed returns true iff every purpose number in ps exists in
// the ParsedConsent, otherwise false.
func (p *ParsedConsent) EveryPurposeAllowed(ps []int) bool {
	for _, rp := range ps {
		if !p.PurposesAllowed[rp] {
			return false
		}
	}
	return true
}

// VendorAllowed returns true if the ParsedConsent contains affirmative consent
// for VendorID v.
func (p *ParsedConsent) VendorAllowed(v int) bool {
	if p.IsRangeEncoding {
		for _, re := range p.RangeEntries {
			if re.StartVendorID <= v && v <= re.EndVendorID {
				return !p.DefaultConsent
			}
		}
		return p.DefaultConsent
	}

	return p.ConsentedVendors[v]
}

// RangeEntry defines an inclusive range of vendor IDs from StartVendorID to
// EndVendorID.
type RangeEntry struct {
	StartVendorID int
	EndVendorID   int
}
