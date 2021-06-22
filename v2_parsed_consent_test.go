package iabconsent_test

import (
	"github.com/go-check/check"

	"github.com/LiveRamp/iabconsent"
)

type V2ParsedConsentSuite struct{}

func (v *V2ParsedConsentSuite) TestParseV2(c *check.C) {
	for k, v := range v2ConsentFixtures {
		c.Log(k)

		var p, err = iabconsent.ParseV2(k)

		c.Check(err, check.IsNil)
		c.Check(p, check.DeepEquals, v)
	}
}

func (v *V2ParsedConsentSuite) TestNonV2Input(c *check.C) {
	var _, err = iabconsent.ParseV2("BONMj34ONMj34ABACDENALqAAAAAplY") // V1 string.
	c.Check(err, check.ErrorMatches, "non-v2 string passed to v2 parse method")
}

func (v  *V2ParsedConsentSuite) TestEveryPurposeAllowed(c *check.C) {
	var tcs = []struct{
		purposes []int
		consent map[int]bool
		exp bool
	}{
		{
			purposes: []int{1, 2, 3},
			consent: map[int]bool{1: true, 2: true, 3: true},
			exp: true,
		},
		{
			purposes: []int{1, 2, 3},
			consent: map[int]bool{1: true, 2: true, 3: false},
			exp: false,
		},
		{
			purposes: []int{1, 2, 3},
			consent: map[int]bool{1: true, 2: true},
			exp: false,
		},
		{
			purposes: []int{1, 2},
			consent: map[int]bool{1: true, 2: true, 3: true},
			exp: true,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.V2ParsedConsent{
			PurposesConsent: tc.consent,
		}

		c.Check(pc.EveryPurposeAllowed(tc.purposes), check.Equals, tc.exp)
	}
}

func (p  *V2ParsedConsentSuite) TestPurposeAllowed(c *check.C) {
	var tcs = []struct{
		purposes []int
		consent map[int]bool
	}{
		{
			purposes: []int{1, 2, 3},
			consent: map[int]bool{1: true, 2: true, 3: true},
		},
		{
			purposes: []int{1, 2, 3},
			consent: map[int]bool{1: true, 2: true, 3: false},
		},
		{
			purposes: []int{1, 2, 3},
			consent: map[int]bool{1: true, 2: true},
		},
		{
			purposes: []int{1, 2},
			consent: map[int]bool{1: true, 2: true, 3: true},
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.V2ParsedConsent{
			PurposesConsent: tc.consent,
		}

		for i, p := range tc.purposes{
			c.Check(pc.PurposeAllowed(p), check.Equals, tc.consent[i+1])
		}

	}
}

func (v  *V2ParsedConsentSuite) TestEverySpecialFeatureOptInGiven(c *check.C) {
	var tcs = []struct{
		specialFeatures []int
		specialFeatureOptIn map[int]bool
		exp bool
	}{
		{
			specialFeatures: []int{1, 2, 3},
			specialFeatureOptIn: map[int]bool{1: true, 2: true, 3: true},
			exp: true,
		},
		{
			specialFeatures: []int{1, 2, 3},
			specialFeatureOptIn: map[int]bool{1: true, 2: true, 3: false},
			exp: false,
		},
		{
			specialFeatures: []int{1, 2, 3},
			specialFeatureOptIn: map[int]bool{1: true, 2: true},
			exp: false,
		},
		{
			specialFeatures: []int{1, 2},
			specialFeatureOptIn: map[int]bool{1: true, 2: true, 3: true},
			exp: true,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.V2ParsedConsent{
			SpecialFeaturesOptIn: tc.specialFeatureOptIn,
		}

		c.Check(pc.EverySpecialFeaturesOptInGiven(tc.specialFeatures), check.Equals, tc.exp)
	}
}

func (v  *V2ParsedConsentSuite) TestVendorAllowed(c *check.C) {
	var tcs = []struct{
		vendor int
		isRange bool
		entries []*iabconsent.RangeEntry
		vendors map[int]bool
		exp bool
	}{
		{
			vendor: 123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 100,
					EndVendorID: 200,
				},
			},
			exp: true,
		},
		{
			vendor: 123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 50,
					EndVendorID: 60,
				},
				{
					StartVendorID: 100,
					EndVendorID: 200,
				},
				{
					StartVendorID: 250,
					EndVendorID: 260,
				},
			},
			exp: true,
		},
		{
			vendor: 123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 123,
					EndVendorID: 123,
				},
			},
			exp: true,
		},
		{
			vendor: 123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{
				{
					StartVendorID: 50,
					EndVendorID: 60,
				},
				{
					StartVendorID: 250,
					EndVendorID: 260,
				},
			},
			exp: false,
		},
		{
			vendor: 123,
			isRange: true,
			entries: []*iabconsent.RangeEntry{},
			exp: false,
		},
		{
			vendor: 123,
			isRange: false,
			vendors: map[int]bool{123: true},
			exp: true,
		},
		{
			vendor: 123,
			isRange: false,
			vendors: map[int]bool{123: true, 124: true},
			exp: true,
		},
		{
			vendor: 123,
			isRange: false,
			vendors: map[int]bool{122: true, 124: true},
			exp: false,
		},
		{
			vendor: 123,
			isRange: false,
			vendors: map[int]bool{123: false, 124: true},
			exp: false,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.V2ParsedConsent{
			IsConsentRangeEncoding: tc.isRange,
			ConsentedVendorsRange: tc.entries,
			ConsentedVendors: tc.vendors,
		}

		c.Check(pc.VendorAllowed(tc.vendor), check.Equals, tc.exp)
	}
}

func (v  *V2ParsedConsentSuite) TestPublisherRestricted(c *check.C) {
	var tcs = []struct{
		purposes []int
		vendor int
		numRestrictions int
		restrictions []*iabconsent.PubRestrictionEntry
		exp bool
	}{
		{
			purposes: []int{1, 2, 3},
			vendor: 123,
			numRestrictions: 0,
			exp: false,
		},
		{
			purposes: []int{1, 2, 3},
			vendor: 123,
			numRestrictions: 1,
			restrictions: []*iabconsent.PubRestrictionEntry{
				{
					PurposeID: 4,
					RestrictionType: iabconsent.PurposeFlatlyNotAllowed,
					NumEntries: 1,
					RestrictionsRange: []*iabconsent.RangeEntry{
						{
							StartVendorID: 123,
							EndVendorID: 123,
						},
					},
				},
			},
			exp: false,
		},
		{
			purposes: []int{1, 2, 3},
			vendor: 123,
			numRestrictions: 1,
			restrictions: []*iabconsent.PubRestrictionEntry{
				{
					PurposeID: 3,
					RestrictionType: iabconsent.PurposeFlatlyNotAllowed,
					NumEntries: 1,
					RestrictionsRange: []*iabconsent.RangeEntry{
						{
							StartVendorID: 123,
							EndVendorID: 123,
						},
					},
				},
			},
			exp: true,
		},
		{
			purposes: []int{1, 2, 3},
			vendor: 123,
			numRestrictions: 1,
			restrictions: []*iabconsent.PubRestrictionEntry{
				{
					PurposeID: 3,
					RestrictionType: iabconsent.RequireConsent,
					NumEntries: 1,
					RestrictionsRange: []*iabconsent.RangeEntry{
						{
							StartVendorID: 123,
							EndVendorID: 123,
						},
					},
				},
			},
			exp: false,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.V2ParsedConsent{
			NumPubRestrictions: tc.numRestrictions,
			PubRestrictionEntries: tc.restrictions,
		}

		c.Check(pc.PublisherRestricted(tc.purposes, tc.vendor), check.Equals, tc.exp)
	}
}

func (v  *V2ParsedConsentSuite) TestSuitableToProcess(c *check.C) {
	var tcs = []struct{
		vendor int
		purposes []int
		consentedPurposes map[int]bool
		consentedVendors map[int]bool
		numRestrictions int
		restrictions []*iabconsent.PubRestrictionEntry
		exp bool
	}{
		{
			vendor: 123,
			purposes: []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors: map[int]bool{123: true},
			numRestrictions: 0,
			exp: true,
		},
		{
			vendor: 123,
			purposes: []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors: map[int]bool{123: true},
			numRestrictions: 1,
			restrictions: []*iabconsent.PubRestrictionEntry{
				{
					PurposeID: 3,
					RestrictionType: iabconsent.PurposeFlatlyNotAllowed,
					NumEntries: 1,
					RestrictionsRange: []*iabconsent.RangeEntry{
						{
							StartVendorID: 100,
							EndVendorID: 110,
						},
						{
							StartVendorID: 130,
							EndVendorID: 140,
						},
					},
				},
			},
			exp: true,
		},
		{
			vendor: 123,
			purposes: []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors: map[int]bool{123: true},
			numRestrictions: 1,
			restrictions: []*iabconsent.PubRestrictionEntry{
				{
					PurposeID: 3,
					RestrictionType: iabconsent.PurposeFlatlyNotAllowed,
					NumEntries: 1,
					RestrictionsRange: []*iabconsent.RangeEntry{
						{
							StartVendorID: 123,
							EndVendorID: 123,
						},
					},
				},
			},
			exp: false,
		},
		{
			vendor: 123,
			purposes: []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors: map[int]bool{123: false},
			numRestrictions: 0,
			exp: false,
		},
		{
			vendor: 123,
			purposes: []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: false},
			consentedVendors: map[int]bool{123: true},
			numRestrictions: 0,
			exp: false,
		},
		{
			vendor: 123,
			purposes: []int{1, 2, 3},
			consentedPurposes: map[int]bool{1: true, 2: true, 3: true},
			consentedVendors: map[int]bool{123: false},
			numRestrictions: 1,
			restrictions: []*iabconsent.PubRestrictionEntry{
				{
					PurposeID: 3,
					RestrictionType: iabconsent.PurposeFlatlyNotAllowed,
					NumEntries: 1,
					RestrictionsRange: []*iabconsent.RangeEntry{
						{
							StartVendorID: 123,
							EndVendorID: 123,
						},
					},
				},
			},
			exp: false,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var pc = &iabconsent.V2ParsedConsent{
			PurposesConsent: tc.consentedPurposes,
			ConsentedVendors: tc.consentedVendors,
			NumPubRestrictions: tc.numRestrictions,
			PubRestrictionEntries: tc.restrictions,
		}

		c.Check(pc.SuitableToProcess(tc.purposes, tc.vendor), check.Equals, tc.exp)
	}
}

var _ = check.Suite(&V2ParsedConsentSuite{})
