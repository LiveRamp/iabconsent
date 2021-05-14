package iabconsent

import (
	"strconv"
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

// GetDeniedPurposes returns an array of all the purposes
// which this request did not receive consent.
func (p *ParsedConsent) GetDeniedPurposes(ps []int) []string {
	var deniedPurposes []string
	for _, rp := range ps {
		if !p.PurposesAllowed[rp] {
			deniedPurposes = append(deniedPurposes, strconv.Itoa(rp))
		}
	}
	return deniedPurposes
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

// SuitableToProcess is the union of EveryPurposeAllowed(ps) and
// VendorAllowed(v). It's used to make possible an interface that
// can be used to check whether its legal to process v1 & v2 strings.
func (p *ParsedConsent) SuitableToProcess(ps []int, v int) bool {
	return p.EveryPurposeAllowed(ps) && p.VendorAllowed(v)
}

// RangeEntry defines an inclusive range of vendor IDs from StartVendorID to
// EndVendorID.
type RangeEntry struct {
	StartVendorID int
	EndVendorID   int
}
