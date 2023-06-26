package iabconsent_test

import (
	"encoding/base64"
	"time"

	"github.com/go-check/check"

	"github.com/StackAdapt/iabconsent"
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

func (s *ParseSuite) TestConsentReader_ReadFibonacciInt(c *check.C) {
	var tests = []struct {
		testBytes []byte
		expected  int
	}{
		{testBytes: []byte{0b11000000},
			expected: 1},
		{testBytes: []byte{0b01100000},
			expected: 2},
		{testBytes: []byte{0b00110000},
			expected: 3},
		{testBytes: []byte{0b10110000},
			expected: 4},
		{testBytes: []byte{0b00011000},
			expected: 5},
		{testBytes: []byte{0b10011000},
			expected: 6},
		{testBytes: []byte{0b01011000},
			expected: 7},
	}

	for _, t := range tests {

		var r = iabconsent.NewConsentReader(t.testBytes)
		var v, err = r.ReadFibonacciInt()
		c.Check(err, check.IsNil)
		c.Check(v, check.Equals, t.expected)
	}
}

func (s *ParseSuite) TestConsentReader_ReadFibonacciIntError(c *check.C) {
	var tests = []struct {
		testBytes []byte
	}{
		{testBytes: []byte{0b00000000}},
		{testBytes: []byte{0b0100000, 0b00000000, 0b01010101, 0b00000000}},
	}

	for _, t := range tests {

		var r = iabconsent.NewConsentReader(t.testBytes)
		var v, err = r.ReadFibonacciInt()
		c.Check(err, check.NotNil)
		c.Check(v, check.Equals, 0)
	}
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

func (s *ParseSuite) TestConsentReader_ReadNBitField(c *check.C) {
	var tests = []struct {
		testBytes []byte
		expected  map[int]int
		nBitsPer  uint
		length    uint
		unread    bool
	}{
		{testBytes: []byte{0x5a},
			expected: map[int]int{
				0: 1,
				1: 1,
				2: 2,
				3: 2,
			},
			nBitsPer: 2,
			length:   4,
			unread:   false},
		{testBytes: []byte{0x5a},
			expected: map[int]int{
				0: 0,
				1: 1,
				2: 0,
				3: 1,
				4: 1,
				5: 0,
				6: 1,
				7: 0,
			},
			nBitsPer: 1,
			length:   8,
			unread:   false},
		// Expected Fields: 00 01 10 10 01 00
		// Bytes: 00011010 01000000
		// Hex: 1a 40
		{testBytes: []byte{0x1a, 0x40},
			expected: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 2,
				4: 1,
				5: 0,
			},
			nBitsPer: 2,
			length:   6,
			unread:   true},
		// Expected Fields: 0000 0001 0010 0011	0110 0111 1000 1001 1010 1011 1100 1101 1110 1111
		// Bytes: 00000001 00100011 01000101 01100111 10001001 10101011 11001101 11101111
		{testBytes: []byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111},
			expected: map[int]int{
				0:  0,
				1:  1,
				2:  2,
				3:  3,
				4:  4,
				5:  5,
				6:  6,
				7:  7,
				8:  8,
				9:  9,
				10: 10,
				11: 11,
				12: 12,
				13: 13,
				14: 14,
				15: 15,
			},
			nBitsPer: 4,
			length:   16,
			unread:   false},
	}

	for _, t := range tests {
		var r = iabconsent.NewConsentReader(t.testBytes)
		var v, err = r.ReadNBitField(t.nBitsPer, t.length)
		c.Check(err, check.IsNil)
		c.Check(v, check.DeepEquals, t.expected)
		c.Check(r.HasUnread(), check.Equals, t.unread)
	}
}

func (s *ParseSuite) TestConsentReader_ReadFibonacciRange(c *check.C) {
	var tests = []struct {
		testBytes []byte
		expected  []int
	}{
		// Samples pulled from: https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/259d5d4d05c23b2a5d49abf936a2f732be128e1a/Core/Consent%20String%20Specification.md#section-encoding
		{testBytes: []byte{0b00000000, 0b00100001, 0b11011001, 0b10000000},
			expected: []int{3, 5, 6, 7, 8},
		},
		{testBytes: []byte{0b00000000, 0b00010011},
			expected: []int{2}},
		{testBytes: []byte{0b00000000, 0b00100011, 0b01011000},
			expected: []int{2, 6}},
		{testBytes: []byte{0b00000000, 0b00011000, 0b11110000},
			expected: []int{5, 6}},
		// Start with group to ensure it starts with 1.
		// Cannot start with 0 because there is no way to Fibonacci-encode 0.
		// TestBytes: # of items: 3, 1st item: Range, offset 1, length of 3, 2nd item: Range offset 2, length of 3, 3rd item: Single Item, Add 100 to last value.
		{testBytes: []byte{0b00000000, 0b00111110, 0b01110110, 0b01100010, 0b10000110},
			expected: []int{1, 2, 3, 4, 6, 7, 8, 9, 109},
		},
	}

	for _, t := range tests {

		var r = iabconsent.NewConsentReader(t.testBytes)
		var v, err = r.ReadFibonacciRange()
		c.Check(err, check.IsNil)
		c.Check(v, check.DeepEquals, t.expected)
	}
}

func (s *ParseSuite) TestParse2_error(c *check.C) {
	var tests = []struct {
		EncodedString string
		Error         string
	}{
		{
			EncodedString: "//BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Error:         "parse v1 consent string: illegal base64 data at input byte 0",
		},
		{
			// base64.RawURLEncoding.EncodeToString([]byte("10011010110110101"))
			EncodedString: "BTAwMTEwMTAxMTAxMTAxMDE",
			Error:         ".*index out of range",
		},
		{
			EncodedString: "COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA",
			Error:         "non-v1 string passed to v1 parse method",
		},
	}

	for _, t := range tests {
		_, err := iabconsent.ParseV1(t.EncodedString)
		c.Check(err, check.ErrorMatches, t.Error)
	}
}

func (s *ParseSuite) TestConsentReader_RestrictionType(c *check.C) {
	// Enums: 0, 1, 2, 3.
	// Bits: 00, 01, 10, 11.
	// Hex: 0x1B.
	var r = iabconsent.NewConsentReader([]byte{0x1B})

	var rts = []iabconsent.RestrictionType{
		iabconsent.PurposeFlatlyNotAllowed,   // 0.
		iabconsent.RequireConsent,            // 1.
		iabconsent.RequireLegitimateInterest, // 2.
		iabconsent.Undefined,                 // 3.
	}

	for _, i := range rts {
		var rt, err = r.ReadRestrictionType()
		c.Check(err, check.IsNil)
		c.Check(rt, check.Equals, i)
	}
}

func (s *ParseSuite) TestConsentReader_SegmentType(c *check.C) {
	// Enums: 0, 1, 2, 3.
	// Bits: 000, 001, 010, 011 (, 0000 extra bits).
	// Hex: 0x05, 0x20.
	var r = iabconsent.NewConsentReader([]byte{0x05, 0x30})

	var rts = []iabconsent.SegmentType{
		iabconsent.CoreString,       // 0.
		iabconsent.DisclosedVendors, // 1.
		iabconsent.AllowedVendors,   // 2.
		iabconsent.PublisherTC,      // 3.
	}

	for _, i := range rts {
		var rt, err = r.ReadSegmentType()
		c.Check(err, check.IsNil)
		c.Check(rt, check.Equals, i)
	}
}

func (s *ParseSuite) TestParseVersion(c *check.C) {
	var tcs = []struct {
		s   string
		err string
		exp iabconsent.TCFVersion
	}{
		{
			s:   "Nonsense",
			exp: iabconsent.InvalidTCFVersion,
		},
		{
			s:   "BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA",
			exp: iabconsent.V1,
		},
		{
			s:   "COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA",
			exp: iabconsent.V2,
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		c.Check(iabconsent.TCFVersionFromTCString(tc.s), check.Equals, tc.exp)
	}
}

func (s *ParseSuite) TestFibonacciIndexValue(c *check.C) {
	var tcs = []struct {
		index    int
		expected int
	}{
		// Test in pre-compiled
		{index: 0,
			expected: 0},
		{index: 2,
			expected: 1},
		// Test last value in pre-compiled.
		{index: 13,
			expected: 233},
		{index: 52,
			expected: 32951280099},
		{index: 62,
			expected: 4052739537881},
		{index: 75,
			expected: 2111485077978050},
		{index: 92,
			expected: 7540113804746346429},
	}
	for _, tc := range tcs {
		var fibonacciValue, err = iabconsent.FibonacciIndexValue(tc.index)
		c.Check(fibonacciValue, check.Equals, tc.expected)
		c.Check(err, check.IsNil)
	}
}

func (s *ParseSuite) TestFibonacciIndexValueFail(c *check.C) {
	var tcs = []struct {
		index int
	}{
		{index: 93},
		{index: 1000000000},
	}

	for _, tc := range tcs {
		var _, err = iabconsent.FibonacciIndexValue(tc.index)
		c.Check(err, check.NotNil)
	}
}

func (s *ParseSuite) TestFibonacciPrecompiled(c *check.C) {
	for i, v := range iabconsent.PrecompiledFibonacci {
		var fibValue, err = iabconsent.FibonacciIndexValue(i)
		c.Check(err, check.IsNil)
		c.Check(v, check.Equals, fibValue)
	}
}
