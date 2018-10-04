package iabconsent

import (
	"encoding/base64"
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

// Parse takes a base64 Raw URL Encoded string which represents a Vendor
// Consent String and returns a ParsedConsent with its fields populated with
// the values stored in the string.
//
// Example Usage:
//
//   var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
func Parse(s string) (*ParsedConsent, error) {
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
