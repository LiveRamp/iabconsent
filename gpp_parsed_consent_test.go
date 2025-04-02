package iabconsent_test

import (
	"encoding/base64"
	"github.com/go-check/check"
	"github.com/pkg/errors"
	"strings"

	"github.com/LiveRamp/iabconsent"
)

type GppParseSuite struct{}

var _ = check.Suite(&GppParseSuite{})

func (s *GppParseSuite) TestParseGppHeader(c *check.C) {
	var tcs = []struct {
		description string
		header      string
		expected    *iabconsent.GppHeader
	}{
		{
			// Examples pulled from: https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Core/Consent%20String%20Specification.md#gpp-string-examples
			description: "EU TCF v2 Section",
			header:      "DBABM",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{2}},
		},
		{
			description: "EU TCF v2 + USPrivacy String Sections",
			header:      "DBACNY",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{2, 6}},
		},
		{
			description: "Canadian TCF + USPrivacy String Sections",
			header:      "DBABjw",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{5, 6}},
		},
		{
			description: "US National MSPA (Multi-State Privacy Agreement)",
			header:      "DBABL",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{7}},
		},
		{
			description: "US Privacy and US National MSPA (Multi-State Privacy Agreement)",
			header:      "DBABzw",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{6, 7}},
		},
		{
			description: "US CA",
			header:      "DBABBg",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{8}},
		},
		{
			header: "DBABBgAA",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{8},
			},
		},
		{
			header: "DBABzYA",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{6, 7, 8},
			},
		},
		{
			header: "DBACOaw",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{2, 5, 6, 7, 8, 9, 10, 11, 12},
			},
		},
	}

	for _, tc := range tcs {
		c.Log(tc)
		var g, err = iabconsent.ParseGppHeader(tc.header)
		c.Check(err, check.IsNil)
		c.Check(g, check.DeepEquals, tc.expected)
	}
}

func (s *GppParseSuite) TestParseGppHeaderError(c *check.C) {
	var tcs = []struct {
		description string
		header      string
		expected    error
	}{
		{
			description: "GPP Header must be 3, as of Jan. 2023.",
			// []byte{0b00000100, 0b00010000, 0b00000010, 0b00110101, 0b10000000}
			header:   "BBACNYA",
			expected: errors.New("wrong gpp header type 1"),
		},
		{
			description: "GPP Header must be 3, as of Jan. 2023, without trailing zero-padding.",
			// Six bit groupings: 000001 000001 000000 000010 001101 011000
			header:   "BBACNY",
			expected: errors.New("wrong gpp header type 1"),
		},
		{
			description: "Only support GPP Version 1, as of Jan. 2023",
			// []byte{0b00001100, 0b00100000, 0b00000010, 0b00110101, 0b10000000}
			header:   "DCACNYA",
			expected: errors.New("unsupported gpp version 2"),
		},
		{
			description: "Only support GPP Version 1, as of Jan. 2023, without trailing zero-padding.",
			// Six bit groupings: 000011 000010 000000 000010 001101 011000
			header:   "DCACNY",
			expected: errors.New("unsupported gpp version 2"),
		}}

	for _, tc := range tcs {
		c.Log(tc)
		var g, err = iabconsent.ParseGppHeader(tc.header)
		c.Check(g, check.IsNil)
		c.Check(err, check.NotNil)
		c.Check(err.Error(), check.Equals, tc.expected.Error())
	}
}

func (s *MspaSuite) TestMapGppSectionToParser(c *check.C) {
	for gppString, expectedValues := range gppParsedConsentFixtures {
		c.Log(gppString)

		var gppSections, err = iabconsent.MapGppSectionToParser(gppString)

		c.Check(err, check.IsNil)
		// Instead of checking the parsing functions, run each of them to ensure the final values match.
		c.Check(gppSections, check.HasLen, len(expectedValues))
		for _, sect := range gppSections {
			consent, err := sect.ParseConsent()
			c.Check(err, check.IsNil)
			c.Check(consent, check.DeepEquals, expectedValues[sect.GetSectionId()])
		}
	}
}

func (s *MspaSuite) TestMapTcfeuSectionToParser(c *check.C) {
	for gppString, expectedValues := range tcfeuParsedConsentFixtures {
		c.Log(gppString)

		var gppSections, err = iabconsent.MapGppSectionToParser(gppString)

		c.Check(err, check.IsNil)
		// Instead of checking the parsing functions, run each of them to ensure the final values match.
		c.Check(gppSections, check.HasLen, len(expectedValues))
		for _, sect := range gppSections {
			consent, err := sect.ParseConsent()
			c.Check(err, check.IsNil)
			c.Check(consent, check.DeepEquals, expectedValues[sect.GetSectionId()])
		}
	}
}

func (s *MspaSuite) TestMapCcpaSectionToParser(c *check.C) {
	for gppString, expectedValues := range ccpaParsedConsentFixtures {
		c.Log(gppString)

		var gppSections, err = iabconsent.MapGppSectionToParser(gppString)

		c.Check(err, check.IsNil)
		// Instead of checking the parsing functions, run each of them to ensure the final values match.
		c.Check(gppSections, check.HasLen, len(expectedValues))
		for _, sect := range gppSections {
			consent, err := sect.ParseConsent()
			c.Check(err, check.IsNil)
			c.Check(consent, check.DeepEquals, expectedValues[sect.GetSectionId()])
		}
	}
}

func (s *MspaSuite) TestMapTcfcaSectionToParser(c *check.C) {
	for gppString, expectedValues := range tcfcaParsedConsentFixtures {
		c.Log(gppString)

		var gppSections, err = iabconsent.MapGppSectionToParser(gppString)

		c.Check(err, check.IsNil)
		// Instead of checking the parsing functions, run each of them to ensure the final values match.
		c.Check(gppSections, check.HasLen, len(expectedValues))
		for _, sect := range gppSections {
			consent, err := sect.ParseConsent()
			c.Check(err, check.IsNil)
			c.Check(consent, check.DeepEquals, expectedValues[sect.GetSectionId()])
		}
	}
}

func (s *MspaSuite) TestMapGppSectionToParserErrors(c *check.C) {
	tcs := []struct {
		desc     string
		gpp      string
		expected error
	}{
		{
			desc:     "No sections.",
			gpp:      "DBABL",
			expected: errors.New("not enough gpp segments"),
		},
		// Disabled as we allow partial sections
		//{
		//	desc:     "Mismatched # of sections, header expects 1.",
		//	gpp:      "DBABL~section1~section2",
		//	expected: errors.New("mismatch number of sections"),
		//},
		{
			desc:     "Bad header.",
			gpp:      "badheader~BVVqAAEABCA.QA",
			expected: errors.New("read gpp header: wrong gpp header type 27"),
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var p, err = iabconsent.MapGppSectionToParser(t.gpp)

		c.Check(p, check.IsNil)
		c.Check(err, check.ErrorMatches, t.expected.Error())
	}
}

// Disabled as we allow partial sections
//func (s *MspaSuite) TestParseGppConsent(c *check.C) {
//	for g, e := range gppParsedConsentFixtures {
//		c.Log(g)
//
//		var p, err = iabconsent.ParseGppConsent(g)
//
//		c.Check(err, check.IsNil)
//		c.Check(p, check.HasLen, len(e))
//		for i, expected := range e {
//			parsed, found := p[i]
//			c.Check(found, check.Equals, true)
//			c.Check(parsed, check.DeepEquals, expected)
//		}
//	}
//}

// Disabled as we allow partial sections
//func (s *MspaSuite) TestParseGppConsentError(c *check.C) {
//	tcs := []struct {
//		desc string
//		gpp  string
//	}{
//		{
//			desc: "Empty Subsection.",
//			gpp:  "DBABzw~1YNN~BVVqAAEABCA.",
//		},
//		{
//			desc: "Empty Subsection.",
//			gpp:  "DBABAw~Bqqqqqqo.",
//		},
//	}
//	for _, tc := range tcs {
//		c.Log(tc.desc)
//
//		var p, err = iabconsent.ParseGppConsent(tc.gpp)
//
//		// Despite an error in the underlying parsing, we quietly do not add the bad value to the map.
//		c.Check(err, check.IsNil)
//		c.Check(p, check.HasLen, 0)
//	}
//}

func (s *GppParseSuite) TestParseGppSubSections(c *check.C) {
	var tcs = []struct {
		description        string
		subsections        string
		expectedSubsection *iabconsent.GppSubSection
		expectedError      error
	}{
		{
			description: "GPC Type, false value",
			// 01000000
			subsections: "QA",
			expectedSubsection: &iabconsent.GppSubSection{
				Gpc: false,
			},
		},
		{
			description: "GPC Type, true value.",
			// 01100000
			subsections: "YA",
			expectedSubsection: &iabconsent.GppSubSection{
				Gpc: true,
			},
		},
		{
			description: "No GPC Type.",
			// 00000000
			subsections: "AA",
			expectedSubsection: &iabconsent.GppSubSection{
				Gpc: false,
			},
		},
		{
			description: "GPC True, then GPC False, should remain True.",
			// 01100000.01000000
			subsections: "YA.QA",
			expectedSubsection: &iabconsent.GppSubSection{
				Gpc: true,
			},
		},
		{
			description: "GPC False, then GPC True, should remain True.",
			// 01000000.01100000
			subsections: "QA.YA",
			expectedSubsection: &iabconsent.GppSubSection{
				Gpc: true,
			},
		},
		{
			description: "GPC Error.",
			// Blank value
			subsections:        "",
			expectedSubsection: nil,
			expectedError:      errors.New("parse gpp subsection type: read int: read bits (index=0, length=2): bits: index out of range"),
		},
	}

	for _, tc := range tcs {
		c.Log(tc)
		// There may be >1 subsections, and func expects them as an array, so split.
		subsect := strings.Split(tc.subsections, ".")
		var g, err = iabconsent.ParseGppSubSections(subsect)
		if tc.expectedError == nil {
			c.Check(err, check.IsNil)
		} else {
			c.Check(err.Error(), check.Equals, tc.expectedError.Error())
		}
		c.Check(g, check.DeepEquals, tc.expectedSubsection)
	}
}

func (s *GppParseSuite) TestParseGpcSubSections(c *check.C) {
	var tcs = []struct {
		description string
		subsection  string
		expected    bool
	}{
		{
			description: "All 0 bits.",
			// 0000000
			subsection: "AA",
			expected:   false,
		},
		{
			description: "Second bit 1.",
			// 01000000
			subsection: "QA",
			expected:   false,
		},
		{
			description: "First bit 1.",
			// 1000000
			subsection: "gA",
			expected:   true,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)
		b, err := base64.RawURLEncoding.DecodeString(tc.subsection)
		c.Check(err, check.IsNil)
		var cr = iabconsent.NewConsentReader(b)
		g, err := iabconsent.ParseGpcSubsection(cr)
		c.Check(err, check.IsNil)
		c.Check(g, check.DeepEquals, tc.expected)
	}
}

func (s *GppParseSuite) TestParse1(c *check.C) {
	var gppSections, err = iabconsent.MapGppSectionToParser("DBABBg~BUoAAAA")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBABTYA~1YNY~000001010101010000101010000000000000000000000000000000010101")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBABM~")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBABBg~BAAAAAA")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBABRg~BVoAAA")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBABRgA~bSFgmiU")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBABFg~BVKAAAA")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 1)

	gppSections, err = iabconsent.MapGppSectionToParser("DBACOe~CPzyhEAPzyhEAEXdtAENDXCwAP_AAH_AACiQI9AB4C5EQCFDcHJNAIoUAAQDQIhAACAgAAABgYAACBoAAIwAAAAwAAAAAAoCAAAAIABAAAEAAAAAAAEAAAAAAAEAAEAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiAAAAAIAEEAAAAACAAEAAAgAABAAAgAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAQQUgSgALAAqABcADIAHAAQAAkABkADQAHIAPAAfQBEAEUAJgATwApABfADMAGgAPwAhABSwDKAMsAc8A7gDvAIHAQcBCACLAFPALqAvMBkwDLAGfANVAfuBBQAA~CPzyhEAPzyhEAEXdtAENDXCgAf-AAP-AAAj0AHgLkRAIUNwck0AihQABANAiEAAICAAAAGBgAAIGgAAjAAAADAAAAAACgIAAAAgAEAAAQAAAAAAAQAAAAAAAQAAQAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAgAQQAAAAAIAAQAACAAAEAACAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAABBBSBKAAsACoAFwAMgAcABAACQAGQANAAcgA8AB9AEQARQAmABPACkAF8AMwAaAA_ACEAFLAMoAywBzwDuAO8AgcBBwEIAIsAU8AuoC8wGTAMsAZ8A1UB-4EFA~1YN-")
	c.Check(err, check.IsNil)
	c.Check(len(gppSections), check.Equals, 3)
}
