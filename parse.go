package iabconsent

import (
	"encoding/base64"
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

// ReadRestrictionType reads two bits and returns an enum |RestrictionType|.
func (r *ConsentReader) ReadRestrictionType() (RestrictionType, error) {
	var rt, err = r.ReadInt(2)
	return RestrictionType(rt), err
}

// ReadSegmentType reads three bits and returns an enum |OOBSegmentType|.
func (r *ConsentReader) ReadSegmentType() (OOBSegmentType, error) {
	var rt, err = r.ReadInt(3)
	return OOBSegmentType(rt), err
}

// Parse takes a base64 Raw URL Encoded string which represents a Vendor
// Consent String and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//   var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
//
// Deprecated: Use ParseV1 to parse V1 consent strings.
func Parse(s string) (*ParsedConsent, error) {
	return ParseV1(s)
}

// Parse takes a base64 Raw URL Encoded string which represents a TCF v1.1
// string and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//   var pc, err = iabconsent.ParseV1("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
func ParseV1(s string) (*ParsedConsent, error) {
	var b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "parse consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	var p = &ParsedConsent{}
	p.Version, _ = r.ReadInt(6)
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

// Parse takes a base64 Raw URL Encoded string which represents a TCF v2
// string and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//   var pc, err = iabconsent.ParseV2("COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA")
func ParseV2(s string) (*V2ParsedConsent, error) {
	var ss = strings.Split(s, ".")

	var b, err = base64.RawURLEncoding.DecodeString(ss[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	var p = &V2ParsedConsent{}
	p.Version, _ = r.ReadInt(6)
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

	p.NumPubRestrictions, _ = r.ReadInt(16)

	for i := 0; i < p.NumPubRestrictions; i++ {
		p.PubRestrictionEntries[i].PurposeID, _ = r.ReadInt(6)
		p.PubRestrictionEntries[i].RestrictionType, _ = r.ReadRestrictionType()
		p.PubRestrictionEntries[i].NumEntries, _ = r.ReadInt(12)
		p.PubRestrictionEntries[i].RestrictionsRange, _ = r.ReadRangeEntries(uint(p.PubRestrictionEntries[i].NumEntries))
	}

	for _, oob := range ss[1:] {
		b, err = base64.RawURLEncoding.DecodeString(oob)
		if err != nil {
			return nil, errors.Wrap(err, "parse consent string")
		}

		r = NewConsentReader(b)
		var st, _ = r.ReadSegmentType()
		switch st {
		case DisclosedVendors:
			p.OOBDisclosedVendors = parseVendors(r, st)
		case AllowedVendors:
			p.OOBAllowedVendors = parseVendors(r, st)
		case PublisherTC:
			var ptc = &PublisherTCEntry{
				SegmentType: st,
			}
			ptc.PubPurposesConsent, _ = r.ReadBitField(24)
			ptc.PubPurposesLITransparency, _ = r.ReadBitField(24)
			ptc.NumCustomPurposes, _ = r.ReadInt(6)
			ptc.CustomPurposesConsent, _ = r.ReadBitField(uint(ptc.NumCustomPurposes))
			ptc.CustomPurposesLITransparency, _ = r.ReadBitField(uint(ptc.NumCustomPurposes))
			p.PublisherTCEntry = ptc
		default:
			// panic?
		}
	}

	return p, r.Err
}

func parseVendors(r *ConsentReader, t OOBSegmentType) *OOBVendorList {
	var v = &OOBVendorList{
		SegmentType: t,
	}
	v.MaxVendorID, _ = r.ReadInt(16)
	v.IsRangeEncoding, _ = r.ReadBool()
	if v.IsRangeEncoding {
		v.NumEntries, _ = r.ReadInt(12)
		v.VendorEntries, _ = r.ReadRangeEntries(uint(v.NumEntries))
	} else {
		v.Vendors, _ = r.ReadBitField(uint(v.MaxVendorID))
	}
	return v
}

type StringVersion int

const (
	// Invalid represents an invalid version.
	Invalid StringVersion = iota
	// V1 represents a TCF v1.1 string.
	V1
	// V2 represents a TCF v2 string.
	V2
)

// ParseWithoutVersion allows the caller to pass any valid consent string
// and get back a Consent, which will allow them to determine if required
// purposes are allowed and whether a given vendor is allowed. This can be
// useful as currently both specs use the same query parameter to pass the
// strings.
func ParseVersion(s string) (StringVersion, error) {
	var ss = strings.Split(s, ".")

	var b, err = base64.RawURLEncoding.DecodeString(ss[0])
	if err != nil {
		return Invalid, errors.Wrap(err, "decoding string")
	}

	var r = NewConsentReader(b)
	var v int
	v, err = r.ReadInt(6)
	if err != nil {
		return Invalid, errors.Wrap(err, "parsing version")
	}

	switch v {
	case 1:
		return V1, nil
	case 2:
		return V2, nil
	default:
		return Invalid, errors.New("not valid version")
	}
}
