package iabconsent

import (
	gc "github.com/go-check/check"
)

type BitStringSuite struct{}

func (p *BitStringSuite) TestParseBytes(c *gc.C) {
	b := []byte{1, 2, 3, 4}
	s := "00000001000000100000001100000100"
	bs := parseBytes(b)

	c.Check(bs.value, gc.Equals, s)
}

func (p *BitStringSuite) TestParseInt(c *gc.C) {
	i, err := dummyBitString().parseInt(4, 100)
	c.Check(err, gc.Equals, errOutOfRange)

	i, err = dummyBitString().parseInt(4, 8)
	c.Check(err, gc.IsNil)
	c.Check(i, gc.Equals, 78)
}

func (p *BitStringSuite) TestParseInt64(c *gc.C) {
	i, err := dummyBitString().parseInt64(4, 100)
	c.Check(err, gc.Equals, errOutOfRange)

	i, err = dummyBitString().parseInt64(4, 8)
	c.Check(err, gc.IsNil)
	c.Check(i, gc.Equals, int64(78))
}

func (p *BitStringSuite) TestParseBitList(c *gc.C) {
	i, err := dummyBitString().parseBitList(4, 100)
	c.Check(err, gc.Equals, errOutOfRange)

	i, err = dummyBitString().parseBitList(4, 8)
	c.Check(err, gc.IsNil)
	c.Check(i, gc.DeepEquals, map[int]bool{
		2: true,
		5: true,
		6: true,
		7: true,
	})
}

func (p *BitStringSuite) TestParseBit(c *gc.C) {
	i, err := dummyBitString().parseBool(100)
	c.Check(err, gc.Equals, errOutOfRange)

	i, err = dummyBitString().parseBool(4)
	c.Check(err, gc.IsNil)
	c.Check(i, gc.Equals, false)
}

func (p *BitStringSuite) TestParseString(c *gc.C) {
	i, err := dummyBitString().parseString(4, 11)
	c.Check(err, gc.Equals, errWrongLength)

	i, err = dummyBitString().parseString(4, 13)
	c.Check(err, gc.Equals, errOutOfRange)

	i, err = dummyBitString().parseString(4, 12)
	c.Check(err, gc.IsNil)
	c.Check(i, gc.Equals, "Td")
}

// dummyBitString returns a *bitString to test on.
func dummyBitString() *bitString {
	return &bitString{"00000100111000110"}
}

var _ = gc.Suite(&BitStringSuite{})
