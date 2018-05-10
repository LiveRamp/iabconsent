package iabconsent

import (
	"sort"

	"github.com/go-check/check"
)

type ParsedConsentSuite struct{}

func (p *ParsedConsentSuite) TestErrorCases(c *check.C) {
	var cases = []struct {
		EncodedString string
		Error         string
	}{
		{
			EncodedString: "//BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Error:         "illegal base64 data at input byte 0",
		},
		{
			// base64.RawURLEncoding.EncodeToString([]byte("10011010110110101"))
			EncodedString: "MTAwMTEwMTAxMTAxMTAxMDE",
			Error:         "index out of range",
		},
	}

	for _, tc := range cases {
		c.Log(tc.EncodedString)
		_, err := Parse(tc.EncodedString)
		c.Check(err.Error(), check.Equals, tc.Error)
	}
}

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
		pc, err := Parse(tc.EncodedString)
		c.Check(err, check.IsNil)

		normalizeParsedConsent(pc)
		normalizeParsedConsent(consentFixtures[tc.Type])

		c.Assert(pc, check.DeepEquals, consentFixtures[tc.Type])
	}
}

func (p *ParsedConsentSuite) TestGetMethods(c *check.C) {
	pc := consentFixtures[BitField]
	c.Check(pc.ConsentString(), check.Equals, pc.consentString)
	c.Check(pc.Version(), check.Equals, pc.version)
	c.Check(pc.Created(), check.Equals, pc.created)
	c.Check(pc.LastUpdated(), check.Equals, pc.lastUpdated)
	c.Check(pc.CmpID(), check.Equals, pc.cmpID)
	c.Check(pc.CmpVersion(), check.Equals, pc.cmpVersion)
	c.Check(pc.ConsentScreen(), check.Equals, pc.consentScreen)
	c.Check(pc.ConsentLanguage(), check.Equals, pc.consentLanguage)
	c.Check(pc.VendorListVersion(), check.Equals, pc.vendorListVersion)
	c.Check(pc.MaxVendorID(), check.Equals, pc.maxVendorID)
}

func (p *ParsedConsentSuite) TestPurposeAllowed(c *check.C) {
	pc := consentFixtures[BitField]
	c.Check(pc.PurposeAllowed(1), check.Equals, true)
	c.Check(pc.PurposeAllowed(2), check.Equals, false)
}

func (p *ParsedConsentSuite) TestPurposesAllowed(c *check.C) {
	pc := consentFixtures[BitField]

	c.Check(pc.PurposesAllowed([]int{1, 3}), check.Equals, true)
	c.Check(pc.PurposesAllowed([]int{1, 4}), check.Equals, false)
}

func (p *ParsedConsentSuite) TestVendorAllowed(c *check.C) {
	pc := consentFixtures[BitField]

	c.Check(pc.VendorAllowed(1), check.Equals, true)
	c.Check(pc.VendorAllowed(3), check.Equals, false)

	pc = consentFixtures[MultipleRangesMixed]

	c.Check(pc.VendorAllowed(123), check.Equals, true)
	c.Check(pc.VendorAllowed(345), check.Equals, true)
	c.Check(pc.VendorAllowed(400), check.Equals, true)
	c.Check(pc.VendorAllowed(456), check.Equals, true)

	c.Check(pc.VendorAllowed(1), check.Equals, false)
	c.Check(pc.VendorAllowed(150), check.Equals, false)
	c.Check(pc.VendorAllowed(500), check.Equals, false)
}

func normalizeParsedConsent(p *ParsedConsent) {
	sort.Slice(p.rangeEntries, func(i, j int) bool {
		return p.rangeEntries[i].StartVendorID < p.rangeEntries[j].StartVendorID
	})
}

var _ = check.Suite(&ParsedConsentSuite{})
