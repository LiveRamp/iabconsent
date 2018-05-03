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
		Type consentType
		EncodedString string
	}{
		{
			Type:         BitField,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAplY",
		},
		{
			Type:         SingleRangeWithSingleID,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqABAD2AAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:         SingleRangeWithRange,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqABgD2AdQAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:         MultipleRangesWithSingleID,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACAD2AOoAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:         MultipleRangesWithRange,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACgD2AdUBWQHIAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:         MultipleRangesMixed,
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

func normalizeParsedConsent(p *ParsedConsent) {
	sort.Slice(p.RangeEntries, func(i, j int) bool {
		return p.RangeEntries[i].SingleOrRange
	})
	sort.Slice(p.RangeEntries, func(i, j int) bool {
		return p.RangeEntries[i].SingleVendorID < p.RangeEntries[j].SingleVendorID
	})
	sort.Slice(p.RangeEntries, func(i, j int) bool {
		return p.RangeEntries[i].StartVendorID < p.RangeEntries[j].StartVendorID
	})
}

var _ = gc.Suite(&ParsedConsentSuite{})
