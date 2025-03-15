package iabconsent

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rupertchen/go-bits"
)

const (
	// dsPerS is deciseconds per second
	dsPerS = 10
	// nsPerDs is nanoseconds per decisecond
	nsPerDs = int64(time.Millisecond * 100)
)

var PrecompiledFibonacci = map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 8: 21, 9: 34, 10: 55, 11: 89, 12: 144, 13: 233}

// ConsentReader provides additional Consent String-specific bit-reading
// functionality on top of bits.Reader.
type ConsentReader struct {
	*bits.Reader
}

// NewConsentReader returns a new ConsentReader backed by src.
func NewConsentReader(src []byte) *ConsentReader {
	return &ConsentReader{bits.NewReader(bits.NewBitmap(src))}
}

// ReadInt reads the next n bits and converts them to an int.
func (r *ConsentReader) ReadInt(n uint) (int, error) {
	if b, err := r.ReadBits(n); err != nil {
		return 0, errors.WithMessage(err, "read int")
	} else {
		return int(b), nil
	}
}

// FibonacciIndexValue is a helper function to get the Fibonacci value of an index for Fibonacci encoding.
// These are currently only used for the various consent types, which there are not many, so
// we should only expect to use the smaller indexes. Therefore, create a map, but still allow for
// new indexes to be calculated.
func FibonacciIndexValue(index int) (int, error) {
	// Due to int limitations, we cannot calculate indexes >92 with ints.
	if index > 92 {
		return 0, errors.New("fibonacci: index greater than max of 92")
	}
	if index <= len(PrecompiledFibonacci) {
		return PrecompiledFibonacci[index], nil
	} else {
		return newFib(index), nil
	}
}

// newFib calculates a new Fibonacci value for a given index.
func newFib(n int) int {
	if n <= 1 {
		return n
	} else {
		var n2, n1 = 0, 1
		for i := 2; i < n; i++ {
			n2, n1 = n1, n1+n2
		}

		return n2 + n1
	}
}

// ReadFibonacciInt reads all the bits until two consecutive `1`s, and converts the bits to an int
// using Fibonacci Encoding. More info: https://en.wikipedia.org/wiki/Fibonacci_coding
func (r *ConsentReader) ReadFibonacciInt() (int, error) {
	var previous = 0
	var total = 0
	// If there is an invalid FibonacciEncoding (no `11`), the Consent Reader will eventually
	// fail with a `bits: index out of range` error.
	for i := 0; ; i++ {
		b, err := r.ReadBits(1)
		if err != nil {
			return 0, errors.WithMessage(err, "read fibonacci int")
		} else {
			// Two bits set to 1 indicates the end of the Fibonacci encoding.
			if previous == 1 && b == 1 {
				break
			}
			// Only add the value if bit is set to 1.
			if b == 1 {
				var fibonacciValue int
				// Since FibonacciEncoding skips 0 + 1, offset index by 2.
				fibonacciValue, err = FibonacciIndexValue(i + 2)
				if err != nil {
					return 0, errors.WithMessage(err, "read fibonacci value")
				}
				total += fibonacciValue
				previous = 1
			} else {
				previous = 0
			}
		}
	}
	return total, nil
}

// ReadTime reads the next 36 bits representing the epoch time in deciseconds
// and converts it to a time.Time.
func (r *ConsentReader) ReadTime() (time.Time, error) {
	if b, err := r.ReadBits(36); err != nil {
		return time.Time{}, errors.WithMessage(err, "read time")
	} else {
		var ds = int64(b)
		return time.Unix(ds/dsPerS, (ds%dsPerS)*nsPerDs).UTC(), nil
	}
}

// ReadString returns a string of length n by reading the next 6 * n bits.
func (r *ConsentReader) ReadString(n uint) (string, error) {
	var buf = make([]byte, 0, n)
	for i := uint(0); i < n; i++ {
		if b, err := r.ReadBits(6); err != nil {
			return "", errors.WithMessage(err, "read string")
		} else {
			buf = append(buf, byte(b)+'A')
		}
	}
	return string(buf), nil
}

// ReadBitField reads the next n bits and converts them to a map[int]bool.
func (r *ConsentReader) ReadBitField(n uint) (map[int]bool, error) {
	var m = make(map[int]bool)
	for i := uint(0); i < n; i++ {
		if b, err := r.ReadBool(); err != nil {
			return nil, errors.WithMessage(err, "read bit field")
		} else {
			if b {
				m[int(i)+1] = true
			}
		}
	}
	return m, nil
}

// ReadNBitField reads n bits, l number of times and converts them to a map[int]int.
// This allows a variable number of bits to be read for more possible number values.
func (r *ConsentReader) ReadNBitField(n, l uint) (map[int]int, error) {
	var m = make(map[int]int)
	for f := uint(0); f < l; f++ {
		if readInt, err := r.ReadInt(n); err != nil {
			return nil, errors.WithMessage(err, "read n-bitfield")
		} else {
			// Zero-based index, as no reason for one-based.
			m[int(f)] = readInt
		}
	}
	return m, nil
}

// ReadRangeEntries reads n range entries of 1 + 16 or 32 bits.
func (r *ConsentReader) ReadRangeEntries(n uint) ([]*RangeEntry, error) {
	var ret = make([]*RangeEntry, 0, n)
	var err error
	for i := uint(0); i < n; i++ {
		var isRange bool
		if isRange, err = r.ReadBool(); err != nil {
			return nil, errors.WithMessage(err, "is-range check")
		}
		var start, end int
		if start, err = r.ReadInt(16); err != nil {
			return nil, errors.WithMessage(err, "range start")
		}
		if isRange {
			if end, err = r.ReadInt(16); err != nil {
				return nil, errors.WithMessage(err, "range end")
			}
		} else {
			end = start
		}
		ret = append(ret, &RangeEntry{StartVendorID: start, EndVendorID: end})
	}
	return ret, nil
}

// ReadFibonacciRange reads a range entries of Fibonacci encoded integers.
// Returns an array of numbers. The format of the range field always consists of:
// - int(12) - representing the amount of items to follow
// - (per item) Boolean - representing whether the item is a single ID (0/false) or a group of IDs (1/true)
// - (per item) int(Fibonacci) - representing a) the offset to a single ID or b) the offset to the start ID in case of a group (the offset is from the last seen number, or 0 for the first entry)
// - (per item + only if group) int(Fibonacci) - length of the group
func (r *ConsentReader) ReadFibonacciRange() ([]int, error) {
	var length int
	var err error
	// Get the amount of items to follow
	if length, err = r.ReadInt(12); err != nil {
		return nil, errors.WithMessage(err, "fibonacci length check")
	}
	var ret []int
	for i := uint(0); i < uint(length); i++ {
		var isRange bool
		// Check if item is a single ID (false) or group of IDs (true).
		if isRange, err = r.ReadBool(); err != nil {
			return nil, errors.WithMessage(err, "is-fibonacci-range check")
		}
		var lastSeen, offset int
		// if no values, start at 0. Otherwise, get last value in slice.
		if len(ret) == 0 {
			lastSeen = 0
		} else {
			lastSeen = ret[len(ret)-1]
		}
		if offset, err = r.ReadFibonacciInt(); err != nil {
			return nil, errors.WithMessage(err, "fibonacci range offset")
		}
		if isRange {
			// If a range, we need to get group length to add multiple values to the range.
			var groupLength int
			if groupLength, err = r.ReadFibonacciInt(); err != nil {
				return nil, errors.WithMessage(err, "fibonacci range length")
			}
			// Add offset to last seen value as starting point of range.
			lastSeen += offset
			// Keep appending integers until we reach the group length.
			for o := 0; o <= groupLength; o++ {
				ret = append(ret, lastSeen)
				lastSeen++
			}
		} else {
			// If a single ID, add value to last seen value.
			ret = append(ret, lastSeen+offset)
		}
	}
	return ret, nil
}

// ReadPubRestrictionEntries reads n publisher restriction entries.
func (r *ConsentReader) ReadPubRestrictionEntries(n uint) ([]*PubRestrictionEntry, error) {
	var ret = make([]*PubRestrictionEntry, 0, n)
	var err error

	for i := uint(0); i < n; i++ {
		var purpose int
		if purpose, err = r.ReadInt(6); err != nil {
			return nil, errors.WithMessage(err, "purpose")
		}
		var rt RestrictionType
		if rt, err = r.ReadRestrictionType(); err != nil {
			return nil, errors.WithMessage(err, "restriction type")
		}
		var num int
		if num, err = r.ReadInt(12); err != nil {
			return nil, errors.WithMessage(err, "num entries")
		}
		var rr []*RangeEntry
		if rr, err = r.ReadRangeEntries(uint(num)); err != nil {
			return nil, errors.WithMessage(err, "range entries")
		}
		ret = append(ret, &PubRestrictionEntry{
			PurposeID:         purpose,
			RestrictionType:   rt,
			NumEntries:        num,
			RestrictionsRange: rr,
		})
	}
	return ret, nil
}

// ReadRestrictionType reads two bits and returns an enum |RestrictionType|.
func (r *ConsentReader) ReadRestrictionType() (RestrictionType, error) {
	var rt, err = r.ReadInt(2)
	return RestrictionType(rt), err
}

// ReadSegmentType reads three bits and returns an enum |SegmentType|.
func (r *ConsentReader) ReadSegmentType() (SegmentType, error) {
	var rt, err = r.ReadInt(3)
	return SegmentType(rt), err
}

// ReadVendors reads in a vendor list representing either disclosed or allowed vendor lists.
// It's assumed that the segment type bit has already been read, despite those bits being
// grouped together in the spec. This is done because the logic for the remaining bits differs
// based on the segment type.
func (r *ConsentReader) ReadVendors(t SegmentType) (*OOBVendorList, error) {
	var v = &OOBVendorList{
		SegmentType: t,
	}
	var err error
	if v.MaxVendorID, err = r.ReadInt(16); err != nil {
		return nil, errors.WithMessage(err, "reading vendor ID")
	}
	if v.IsRangeEncoding, err = r.ReadBool(); err != nil {
		return nil, errors.WithMessage(err, "reading is range flag")
	}
	if v.IsRangeEncoding {
		if v.NumEntries, err = r.ReadInt(12); err != nil {
			return nil, errors.WithMessage(err, "reading num entries")
		}
		if v.VendorEntries, err = r.ReadRangeEntries(uint(v.NumEntries)); err != nil {
			return nil, errors.WithMessage(err, "reading vendor range entries")
		}
	} else {
		if v.Vendors, err = r.ReadBitField(uint(v.MaxVendorID)); err != nil {
			return nil, errors.WithMessage(err, "reading vendor bit field")
		}
	}
	return v, nil
}

// ReadPublisherTCEntry reads in a publisher TC entry. It's assumed that the segment type bit
// has already been read, despite those bits being grouped together in the spec. This is done
// because the logic for the remaining bits differs based on the segment type.
func (r *ConsentReader) ReadPublisherTCEntry() (*PublisherTCEntry, error) {
	var ptc = &PublisherTCEntry{
		SegmentType: PublisherTC,
	}
	var err error
	if ptc.PubPurposesConsent, err = r.ReadBitField(24); err != nil {
		return nil, errors.WithMessage(err, "reading purposes bit field")
	}
	if ptc.PubPurposesLITransparency, err = r.ReadBitField(24); err != nil {
		return nil, errors.WithMessage(err, "reading lit transparency bit field")
	}
	if ptc.NumCustomPurposes, err = r.ReadInt(6); err != nil {
		return nil, errors.WithMessage(err, "reading num custom purposes")
	}
	if ptc.CustomPurposesConsent, err = r.ReadBitField(uint(ptc.NumCustomPurposes)); err != nil {
		return nil, errors.WithMessage(err, "reading custom purposes bitfield")
	}
	if ptc.CustomPurposesLITransparency, err = r.ReadBitField(uint(ptc.NumCustomPurposes)); err != nil {
		return nil, errors.WithMessage(err, "reading lit transparency bitfield")
	}
	return ptc, nil
}

// Parse takes a base64 Raw URL Encoded string which represents a Vendor
// Consent String and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//	var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
//
// Deprecated: Use ParseV1 to parse V1 consent strings.
func Parse(s string) (*ParsedConsent, error) {
	return ParseV1(s)
}

// ParseV1 takes a base64 Raw URL Encoded string which represents a TCF v1.1
// string and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//	var pc, err = iabconsent.ParseV1("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
func ParseV1(s string) (*ParsedConsent, error) {
	var b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "parse v1 consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	var p = &ParsedConsent{}
	p.Version, _ = r.ReadInt(6)
	if p.Version != int(V1) {
		return nil, errors.New("non-v1 string passed to v1 parse method")
	}
	p.Created, _ = r.ReadTime()
	p.LastUpdated, _ = r.ReadTime()
	p.CMPID, _ = r.ReadInt(12)
	p.CMPVersion, _ = r.ReadInt(12)
	p.ConsentScreen, _ = r.ReadInt(6)
	p.ConsentLanguage, _ = r.ReadString(2)
	p.VendorListVersion, _ = r.ReadInt(12)
	p.PurposesAllowed, _ = r.ReadBitField(24)
	p.MaxVendorID, _ = r.ReadInt(16)

	p.IsRangeEncoding, _ = r.ReadBool()
	if p.IsRangeEncoding {
		p.DefaultConsent, _ = r.ReadBool()
		p.NumEntries, _ = r.ReadInt(12)
		p.RangeEntries, _ = r.ReadRangeEntries(uint(p.NumEntries))
	} else {
		p.ConsentedVendors, _ = r.ReadBitField(uint(p.MaxVendorID))
	}

	return p, r.Err
}

// ParseV2 takes a base64 Raw URL Encoded string which represents a TCF v2
// string and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//	var pc, err = iabconsent.ParseV2("COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA")
func ParseV2(s string) (*V2ParsedConsent, error) {
	var segments = strings.Split(s, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse v2 consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/47b45ab362515310183bb3572a367b8391ef4613/TCFv2/IAB%20Tech%20Lab%20-%20Consent%20string%20and%20vendor%20list%20formats%20v2.md#about-the-transparency--consent-string-tc-string
	var p = &V2ParsedConsent{}
	p.Version, _ = r.ReadInt(6)
	if p.Version != int(V2) {
		return nil, errors.New("non-v2 string passed to v2 parse method")
	}
	p.Created, _ = r.ReadTime()
	p.LastUpdated, _ = r.ReadTime()
	p.CMPID, _ = r.ReadInt(12)
	p.CMPVersion, _ = r.ReadInt(12)
	p.ConsentScreen, _ = r.ReadInt(6)
	p.ConsentLanguage, _ = r.ReadString(2)
	p.VendorListVersion, _ = r.ReadInt(12)
	p.TCFPolicyVersion, _ = r.ReadInt(6)
	p.IsServiceSpecific, _ = r.ReadBool()
	p.UseNonStandardStacks, _ = r.ReadBool()
	p.SpecialFeaturesOptIn, _ = r.ReadBitField(12)
	p.PurposesConsent, _ = r.ReadBitField(24)
	p.PurposesLITransparency, _ = r.ReadBitField(24)
	// Check for specific 2.2 Requirements and exit early.
	// From IAB Docs: https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/TCFv2/IAB%20Tech%20Lab%20-%20Consent%20string%20and%20vendor%20list%20formats%20v2.md#the-core-string
	// "With TCF v2.2 support for legitimate interest for purpose 3 to 6 has been deprecated. Bits 2 to 5 are required to be set to 0."
	// All future versions will also have the requirement.
	// if mv, _ := p.MinorVersion(); mv >= 2 {
	// 	// Bitfield uses 1-indexing, so we need to check for purposes 3-6 (not bit positions 2-5).
	// 	for lit := 3; lit <= 6; lit++ {
	// 		if p.PurposesLITransparency[lit] != false {
	// 			return nil, errors.Errorf("TCF String Version 2.2 or higher has invalid PurposesLIT %d not set to 0.", lit)
	// 		}
	// 	}
	// }
	p.PurposeOneTreatment, _ = r.ReadBool()
	p.PublisherCC, _ = r.ReadString(2)

	p.MaxConsentVendorID, _ = r.ReadInt(16)
	p.IsConsentRangeEncoding, _ = r.ReadBool()
	if p.IsConsentRangeEncoding {
		p.NumConsentEntries, _ = r.ReadInt(12)
		p.ConsentedVendorsRange, _ = r.ReadRangeEntries(uint(p.NumConsentEntries))
	} else {
		p.ConsentedVendors, _ = r.ReadBitField(uint(p.MaxConsentVendorID))
	}

	p.MaxInterestsVendorID, _ = r.ReadInt(16)
	p.IsInterestsRangeEncoding, _ = r.ReadBool()
	if p.IsInterestsRangeEncoding {
		p.NumInterestsEntries, _ = r.ReadInt(12)
		p.InterestsVendorsRange, _ = r.ReadRangeEntries(uint(p.NumInterestsEntries))
	} else {
		p.InterestsVendors, _ = r.ReadBitField(uint(p.MaxInterestsVendorID))
	}

	p.NumPubRestrictions, _ = r.ReadInt(12)
	p.PubRestrictionEntries, _ = r.ReadPubRestrictionEntries(uint(p.NumPubRestrictions))

	// Parse remaining non-core string segments if they exist.
	for i, segment := range segments[1:] {
		b, err = base64.RawURLEncoding.DecodeString(segment)
		if err != nil {
			return p, errors.Wrap(err, "parsing segment "+strconv.Itoa(i+1))
		}

		r = NewConsentReader(b)
		var st, _ = r.ReadSegmentType()
		switch st {
		case DisclosedVendors:
			if p.OOBDisclosedVendors != nil {
				return p, errors.New("multiple disclosed vendors segments passedg")
			}
			p.OOBDisclosedVendors, _ = r.ReadVendors(st)
		case AllowedVendors:
			if p.OOBAllowedVendors != nil {
				return p, errors.New("multiple allowed vendors segments passed")
			}
			p.OOBAllowedVendors, _ = r.ReadVendors(st)
		case PublisherTC:
			if p.PublisherTCEntry != nil {
				return p, errors.New("multiple publisher TC segments passed")
			}
			p.PublisherTCEntry, _ = r.ReadPublisherTCEntry()
		default:
			return p, errors.New("unrecognized segment type")
		}
	}

	return p, r.Err
}

// TCFVersion is an enum type used for easily identifying which version
// a consent string is.
type TCFVersion int

const (
	// InvalidTCFVersion represents an invalid version.
	InvalidTCFVersion TCFVersion = iota
	// V1 represents a TCF v1.1 string.
	V1
	// V2 represents a TCF v2 string.
	V2
)

// TCFVersionFromTCString allows the caller to pass any valid consent string to
// determine which parse method is appropriate to call or otherwise
// return InvalidTCFVersion (0).
func TCFVersionFromTCString(s string) TCFVersion {
	var ss = strings.SplitN(s, ".", 2)

	var b, err = base64.RawURLEncoding.DecodeString(ss[0])
	if err != nil {
		return InvalidTCFVersion
	}

	var r = NewConsentReader(b)
	var v int
	v, err = r.ReadInt(6)
	if err != nil {
		return InvalidTCFVersion
	}

	switch v {
	case 1:
		return V1
	case 2:
		return V2
	default:
		return InvalidTCFVersion
	}
}
