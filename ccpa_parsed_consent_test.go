package iabconsent

import (
	"github.com/go-check/check"
)

type CcpaSuite struct{}

var _ = check.Suite(&CcpaSuite{})

func (s *CcpaSuite) TestParseCcpa(c *check.C) {
	var ccpaParsedConset, err = ParseCCPA("1YYY")
	c.Assert(err, check.IsNil)
	c.Assert(ccpaParsedConset, check.Not(check.IsNil))
	c.Assert(ccpaParsedConset.Version, check.Equals, 1)
	c.Assert(ccpaParsedConset.Notice, check.Equals, uint8('Y'))
	c.Assert(ccpaParsedConset.OptOutSale, check.Equals, uint8('Y'))
	c.Assert(ccpaParsedConset.LSPACoveredTransaction, check.Equals, uint8('Y'))

	ccpaParsedConset, err = ParseCCPA("1---")
	c.Assert(err, check.IsNil)
	c.Assert(ccpaParsedConset, check.Not(check.IsNil))
	c.Assert(ccpaParsedConset.Version, check.Equals, 1)
	c.Assert(ccpaParsedConset.Notice, check.Equals, uint8('-'))
	c.Assert(ccpaParsedConset.OptOutSale, check.Equals, uint8('-'))
	c.Assert(ccpaParsedConset.LSPACoveredTransaction, check.Equals, uint8('-'))

	ccpaParsedConset, err = ParseCCPA("1NNN")
	c.Assert(err, check.IsNil)
	c.Assert(ccpaParsedConset, check.Not(check.IsNil))
	c.Assert(ccpaParsedConset.Version, check.Equals, 1)
	c.Assert(ccpaParsedConset.Notice, check.Equals, uint8('N'))
	c.Assert(ccpaParsedConset.OptOutSale, check.Equals, uint8('N'))
	c.Assert(ccpaParsedConset.LSPACoveredTransaction, check.Equals, uint8('N'))

	ccpaParsedConset, err = ParseCCPA("1NNNN")
	c.Assert(err, check.Not(check.IsNil))

	ccpaParsedConset, err = ParseCCPA("2NYN")
	c.Assert(err, check.Not(check.IsNil))

	ccpaParsedConset, err = ParseCCPA("1ABC")
	c.Assert(err, check.Not(check.IsNil))
}
