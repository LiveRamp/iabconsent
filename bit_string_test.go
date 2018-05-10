package iabconsent

import (
	"github.com/go-check/check"
)

type BitStringSuite struct{}

func (p *BitStringSuite) TestParseBytes(c *check.C) {
	b := []byte{1, 2, 3, 4}
	s := "00000001000000100000001100000100"
	bs := ParseBytes(b)

	c.Check(bs.value, check.Equals, s)
}

func (p *BitStringSuite) TestParseInt(c *check.C) {
	i, err := dummyBitString().ParseInt(4, 100)
	c.Check(err, check.Equals, errOutOfRange)

	i, err = dummyBitString().ParseInt(4, 8)
	c.Check(err, check.IsNil)
	c.Check(i, check.Equals, 78)
}

func (p *BitStringSuite) TestParseInt64(c *check.C) {
	i, err := dummyBitString().ParseInt64(4, 100)
	c.Check(err, check.Equals, errOutOfRange)

	i, err = dummyBitString().ParseInt64(4, 8)
	c.Check(err, check.IsNil)
	c.Check(i, check.Equals, int64(78))
}

func (p *BitStringSuite) TestParseBitList(c *check.C) {
	i, err := dummyBitString().ParseBitList(4, 100)
	c.Check(err, check.Equals, errOutOfRange)

	i, err = dummyBitString().ParseBitList(4, 8)
	c.Check(err, check.IsNil)
	c.Check(i, check.DeepEquals, map[int]bool{
		2: true,
		5: true,
		6: true,
		7: true,
	})
}

func (p *BitStringSuite) TestParseBit(c *check.C) {
	i, err := dummyBitString().ParseBool(100)
	c.Check(err, check.Equals, errOutOfRange)

	i, err = dummyBitString().ParseBool(4)
	c.Check(err, check.IsNil)
	c.Check(i, check.Equals, false)
}

func (p *BitStringSuite) TestParseString(c *check.C) {
	i, err := dummyBitString().ParseString(4, 11)
	c.Check(err, check.Equals, errWrongLength)

	i, err = dummyBitString().ParseString(4, 13)
	c.Check(err, check.Equals, errOutOfRange)

	i, err = dummyBitString().ParseString(4, 12)
	c.Check(err, check.IsNil)
	c.Check(i, check.Equals, "Td")
}

// dummyBitString returns a *BitString to test on.
func dummyBitString() *BitString {
	return &BitString{"00000100111000110"}
}

var _ = check.Suite(&BitStringSuite{})
