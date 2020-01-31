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

var _ = check.Suite(&V2ParsedConsentSuite{})
