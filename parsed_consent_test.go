package iabconsent

import (
	"sort"

	gc "github.com/go-check/check"
)

type ParsedConsentSuite struct{}

func (p *ParsedConsentSuite) TestErrorCases(c *gc.C) {
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
		c.Check(err.Error(), gc.Equals, tc.Error)
	}
}

func (p *ParsedConsentSuite) TestParseConsentStrings(c *gc.C) {
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
		c.Check(err, gc.IsNil)

		normalizeParsedConsent(pc)
		normalizeParsedConsent(consentFixtures[tc.Type])

		c.Assert(pc, gc.DeepEquals, consentFixtures[tc.Type])
	}
}

func (p *ParsedConsentSuite) TestGetMethods(c *gc.C) {
	pc := consentFixtures[BitField]
	c.Check(pc.ConsentString(), gc.Equals, pc.consentString)
	c.Check(pc.Version(), gc.Equals, pc.version)
	c.Check(pc.Created(), gc.Equals, pc.created)
	c.Check(pc.LastUpdated(), gc.Equals, pc.lastUpdated)
	c.Check(pc.CmpID(), gc.Equals, pc.cmpID)
	c.Check(pc.CmpVersion(), gc.Equals, pc.cmpVersion)
	c.Check(pc.ConsentScreen(), gc.Equals, pc.consentScreen)
	c.Check(pc.ConsentLanguage(), gc.Equals, pc.consentLanguage)
	c.Check(pc.VendorListVersion(), gc.Equals, pc.vendorListVersion)
	c.Check(pc.MaxVendorID(), gc.Equals, pc.maxVendorID)
}

func (p *ParsedConsentSuite) TestPurposeAllowed(c *gc.C) {
	pc := consentFixtures[BitField]
	c.Check(pc.PurposeAllowed(1), gc.Equals, true)
	c.Check(pc.PurposeAllowed(2), gc.Equals, false)
}

func (p *ParsedConsentSuite) TestPurposesAllowed(c *gc.C) {
	pc := consentFixtures[BitField]

	c.Check(pc.PurposesAllowed([]int{1, 3}), gc.Equals, true)
	c.Check(pc.PurposesAllowed([]int{1, 4}), gc.Equals, false)
}

func (p *ParsedConsentSuite) TestVendorAllowed(c *gc.C) {
	pc := consentFixtures[BitField]

	c.Check(pc.VendorAllowed(1), gc.Equals, true)
	c.Check(pc.VendorAllowed(3), gc.Equals, false)

	pc = consentFixtures[MultipleRangesMixed]

	c.Check(pc.VendorAllowed(123), gc.Equals, true)
	c.Check(pc.VendorAllowed(345), gc.Equals, true)
	c.Check(pc.VendorAllowed(400), gc.Equals, true)
	c.Check(pc.VendorAllowed(456), gc.Equals, true)

	c.Check(pc.VendorAllowed(1), gc.Equals, false)
	c.Check(pc.VendorAllowed(150), gc.Equals, false)
	c.Check(pc.VendorAllowed(500), gc.Equals, false)
}

func normalizeParsedConsent(p *ParsedConsent) {
	sort.Slice(p.rangeEntries, func(i, j int) bool {
		return p.rangeEntries[i].StartVendorID < p.rangeEntries[j].StartVendorID
	})
}

var _ = gc.Suite(&ParsedConsentSuite{})
