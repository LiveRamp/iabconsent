package iabconsent_test

import (
	"encoding/base64"
	"time"

	"github.com/go-check/check"

	"github.com/LiveRamp/iabconsent"
)

type ParseSuite struct{}

var _ = check.Suite(&ParseSuite{})

func (s *ParseSuite) TestConsentReader_ReadInt(c *check.C) {
	var tests = []struct {
		expected int
		n        uint
	}{
		{1, 1},
		{0, 1},
		{5, 3},
		{2, 3},
	}

	var r = iabconsent.NewConsentReader([]byte{0xaa})
	for _, t := range tests {
		var v, err = r.ReadInt(t.n)
		c.Check(err, check.IsNil)
		c.Check(v, check.Equals, t.expected)
	}
	c.Check(r.HasUnread(), check.Equals, false)
}

func (s *ParseSuite) TestConsentReader_ReadTime(c *check.C) {
	// 2018-05-18 17:48:31.5 +0000 UTC
	// 1526665711.5 s
	// 15266657115 deci-seconds
	// 0x38df6b35b deci-seconds (hex)
	var r = iabconsent.NewConsentReader([]byte{0x38, 0xdf, 0x6b, 0x35, 0xB0})
	var v, err = r.ReadTime()
	c.Check(err, check.IsNil)
	c.Check(v, check.DeepEquals, time.Unix(1526665711, int64(500*time.Millisecond)).UTC())
	c.Check(r.NumUnread(), check.Equals, 4)
}

func (s *ParseSuite) TestConsentReader_ReadString(c *check.C) {
	// A four character base64 string is the shortest string that decodes to a
	// multiple of 8 bits.
	var b64 = "ABCD"
	var b, err = base64.RawURLEncoding.DecodeString(b64)
	c.Assert(err, check.IsNil)

	var r = iabconsent.NewConsentReader(b)
	var checkString = func(n uint, expected string) {
		var v, err = r.ReadString(n)
		c.Check(err, check.IsNil)
		c.Check(v, check.Equals, expected)
	}
	checkString(1, "A")
	checkString(2, "BC")
	checkString(1, "D")
	c.Check(r.HasUnread(), check.Equals, false)
}

func (s *ParseSuite) TestConsentReader_ReadBitField(c *check.C) {
	var tests = []struct {
		expected map[int]bool
		n        uint
	}{
		{map[int]bool{
			2: true,
		}, 2},
		{map[int]bool{
			2: true,
			3: true,
			5: true,
		}, 6},
	}

	var r = iabconsent.NewConsentReader([]byte{0x5a})
	for _, t := range tests {
		var v, err = r.ReadBitField(t.n)
		c.Check(err, check.IsNil)
		c.Check(v, check.DeepEquals, t.expected)
	}
	c.Check(r.HasUnread(), check.Equals, false)
}

func (s *ParseSuite) TestParse2_error(c *check.C) {
	var tests = []struct {
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
			Error:         ".*index out of range",
		},
	}

	for _, t := range tests {
		_, err := iabconsent.Parse(t.EncodedString)
		c.Check(err, check.ErrorMatches, t.Error)
	}
}
