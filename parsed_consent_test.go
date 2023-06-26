package iabconsent_test

import (
	"sort"

	"github.com/go-check/check"

	"github.com/StackAdapt/iabconsent"
)

type ParsedConsentSuite struct{}

func (p *ParsedConsentSuite) TestParseConsentStrings(c *check.C) {
	var cases = []struct {
		Type          consentType
		EncodedString string
	}{
		{
			Type:          BitField,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAplY",
		},
		{
			Type:          SingleRangeWithSingleID,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqABAD2AAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          SingleRangeWithRange,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqABgD2AdQAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          MultipleRangesWithSingleID,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACAD2AOoAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          MultipleRangesWithRange,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACgD2AdUBWQHIAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          MultipleRangesMixed,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACAD3AVkByAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
	}

	for _, tc := range cases {
		c.Log(tc)
		pc, err := iabconsent.ParseV1(tc.EncodedString)
		c.Check(err, check.IsNil)

		normalizeParsedConsent(pc)
		normalizeParsedConsent(v1ConsentFixtures[tc.Type])

		c.Assert(pc, check.DeepEquals, v1ConsentFixtures[tc.Type])
	}
}

func normalizeParsedConsent(p *iabconsent.ParsedConsent) {
	sort.Slice(p.RangeEntries, func(i, j int) bool {
		return p.RangeEntries[i].StartVendorID < p.RangeEntries[j].StartVendorID
	})
}

func (p *ParsedConsentSuite) TestEveryPurposeAllowed(c *check.C) {
	var tcs = []struct {
		purposes []int
		consent  map[int]bool
		exp      bool
	}{
		{
			purposes: []int{1, 2, 3},
			consent:  map[int]bool{1: true, 2: true, 3: true},
			exp:      true,
		},
		{
			purposes: []int{1, 2, 3},
			consent:  map[int]bool{1: true, 2: true, 3: false},
			exp:      false,
		},
		{
			purposes: []int{1, 2, 3},
			consent:  map[int]bool{1: true, 2: true},
			exp:      false,
		},
		{
			purposes: []int{1, 2},
			consent:  map[int]bool{1: true, 2: true, 3: true},
			exp:      true,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.ParsedConsent{
			PurposesAllowed: tc.consent,
		}

		c.Check(pc.EveryPurposeAllowed(tc.purposes), check.Equals, tc.exp)
	}
}

func (p *ParsedConsentSuite) TestPurposeAllowed(c *check.C) {
	var tcs = []struct {
		purposes []int
		consent  map[int]bool
	}{
		{
			purposes: []int{1, 2, 3},
			consent:  map[int]bool{1: true, 2: true, 3: true},
		},
		{
			purposes: []int{1, 2, 3},
			consent:  map[int]bool{1: true, 2: true, 3: false},
		},
		{
			purposes: []int{1, 2, 3},
			consent:  map[int]bool{1: true, 2: true},
		},
		{
			purposes: []int{1, 2},
			consent:  map[int]bool{1: true, 2: true, 3: true},
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.ParsedConsent{
			PurposesAllowed: tc.consent,
		}

		for i, p := range tc.purposes {
			c.Check(pc.PurposeAllowed(p), check.Equals, tc.consent[i+1])
		}

	}
}

func (p *ParsedConsentSuite) TestVendorAllowed(c *check.C) {
	var tcs = []struct {
		vendor  int
		isRange bool
		entries []*iabconsent.RangeEntry
		vendors map[int]bool
		exp     bool
	}{
		{
			vendor:  123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 100,
					EndVendorID:   200,
				},
			},
			exp: true,
		},
		{
			vendor:  123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 50,
					EndVendorID:   60,
				},
				{
					StartVendorID: 100,
					EndVendorID:   200,
				},
				{
					StartVendorID: 250,
					EndVendorID:   260,
				},
			},
			exp: true,
		},
		{
			vendor:  123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 123,
					EndVendorID:   123,
				},
			},
			exp: true,
		},
		{
			vendor:  123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 50,
					EndVendorID:   60,
				},
				{
					StartVendorID: 250,
					EndVendorID:   260,
				},
			},
			exp: false,
		},
		{
			vendor:  123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{},
			exp:     false,
		},
		{
			vendor:  123,
			isRange: false,
			vendors: map[int]bool{123: true},
			exp:     true,
		},
		{
			vendor:  123,
			isRange: false,
			vendors: map[int]bool{123: true, 124: true},
			exp:     true,
		},
		{
			vendor:  123,
			isRange: false,
			vendors: map[int]bool{122: true, 124: true},
			exp:     false,
		},
		{
			vendor:  123,
			isRange: false,
			vendors: map[int]bool{123: false, 124: true},
			exp:     false,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.ParsedConsent{
			IsRangeEncoding:  tc.isRange,
			RangeEntries:     tc.entries,
			ConsentedVendors: tc.vendors,
		}

		c.Check(pc.VendorAllowed(tc.vendor), check.Equals, tc.exp)
	}
}

func (v *ParsedConsentSuite) TestSuitableToProcess(c *check.C) {
	var tcs = []struct {
		vendor            int
		purposes          []int
		consentedPurposes map[int]bool
		consentedVendors  map[int]bool
		exp               bool
	}{
		{
			vendor:            123,
			purposes:          []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors:  map[int]bool{123: true},
			exp:               true,
		},
		{
			vendor:            123,
			purposes:          []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors:  map[int]bool{123: false},
			exp:               false,
		},
		{
			vendor:            123,
			purposes:          []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: false},
			consentedVendors:  map[int]bool{123: true},
			exp:               false,
		},
		{
			vendor:            123,
			purposes:          []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: false},
			consentedVendors:  map[int]bool{123: false},
			exp:               false,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.ParsedConsent{
			PurposesAllowed:  tc.consentedPurposes,
			ConsentedVendors: tc.consentedVendors,
		}

		c.Check(pc.SuitableToProcess(tc.purposes, tc.vendor), check.Equals, tc.exp)
	}
}

var _ = check.Suite(&ParsedConsentSuite{})
