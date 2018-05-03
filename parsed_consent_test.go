package iabconsent

import (
	gc "github.com/go-check/check"
)

type ParsedConsentSuite struct{}

func (p *ParsedConsentSuite) TestErrorCases(c *gc.C) {
	var cases = []struct {
		EncodedString string
		Error         string
	}{
		{
			EncodedString: "BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAAzzz",
			Error:         "FTP",
		},
	}

	for _, tc := range cases {
		c.Log(tc.EncodedString)
		_, err := Parse(tc.EncodedString)
		c.Check(err, gc.Equals, tc.Error)
		c.Assert(err, gc.Equals, tc.Error)
	}
}

func (p *ParsedConsentSuite) TestFail(c *gc.C) {
	c.Assert(true, gc.IsNil)
}

var _ = gc.Suite(&ParsedConsentSuite{})
