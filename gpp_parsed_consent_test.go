package iabconsent_test

import (
	"github.com/go-check/check"
	"github.com/pkg/errors"

	"github.com/LiveRamp/iabconsent"
)

type GppParseSuite struct{}

var _ = check.Suite(&GppParseSuite{})

func (s *GppParseSuite) TestParseGppHeader(c *check.C) {
	var tcs = []struct {
		description string
		header   string
		expected *iabconsent.GppHeader
	}{
		{
			// Examples pulled from: https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Core/Consent%20String%20Specification.md#gpp-string-examples
			description: "EU TCF v2 Section",
			header: "DBABMA",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{2}},
		},
		{
			description: "EU TCF v2 + USPrivacy String Sections",
			header: "DBACNYA",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{2, 6}},
		},
		{
			description: "Canadian TCF + USPrivacy String Sections",
			header: "DBABjw",
			expected: &iabconsent.GppHeader{
				Type:     3,
				Version:  1,
				Sections: []int{5, 6}},
		},
	}

	for _, tc := range tcs {
		c.Log(tc)
		var g, err = iabconsent.ParseGppHeader(tc.header)
		c.Check(err, check.IsNil)
		c.Check(g, check.DeepEquals, tc.expected)
	}
}

func (s *GppParseSuite) TestParseGppHeaderFail(c *check.C) {
	var tcs = []struct {
		description string
		header   string
		expected error
	}{
		{
			description: "GPP Header must be 3 for now.",
			// []byte{0b00000100, 0b00010000, 0b00000010, 0b00110101, 0b10000000}
			header: "BBACNYA",
			expected: errors.New("wrong gpp header type 1"),
		},
		{
			description: "Only support GPP Version 1",
			// []byte{0b00001100, 0b00100000, 0b00000010, 0b00110101, 0b10000000}
			header: "DCACNYA",
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