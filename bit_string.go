package iabconsent

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Constants used to help do time conversions.
const (
	dsPerS  = 10
	nsPerDs = int64(time.Millisecond * 100)
)

// Common errors for BitString methods.
var (
	errOutOfRange  = errors.New("index out of range")
	errWrongLength = errors.New("bit string length must be multiple of 6")
)

// BitString is a simple struct which has only one field, value.
// The value is expected to be a string containing only ones and zeros
// which represent a Vendor Consent String as defined by the IAB spec
// found here: https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#vendor-consent-string-format-.
// The type is defined to enable common operations needed to parse
// the string which are defined below.
type BitString struct {
	value string
}

// ParseBytes takes in a []byte |b| and returns a BitString |bs|
// whose value is the concatenation of the 8 bit binary representation
// of each element of |b|. Given that the consent string is not necessarily
// a multiple of 8 bits, we pad the end of the string with 0s.
func ParseBytes(b []byte) BitString {
	var buffer bytes.Buffer

	for _, s := range b {
		buffer.WriteString(fmt.Sprintf("%08b", s))
	}

	return BitString{value: buffer.String()}
}

// ParseInt64 takes a bit offset and size and converts the binary
// number produced from that substring slice into an int64.
func (b BitString) ParseInt64(offset, size int) (int64, error) {
	if len(b.value)-1 < offset+size {
		return 0, errOutOfRange
	}
	return strconv.ParseInt(b.value[offset:(offset+size)], 2, 64)
}

// ParseInt takes a bit offset and size and converts the binary
// number produced from that substring slice into an int.
func (b BitString) ParseInt(offset, size int) (int, error) {
	var s, err = b.ParseInt64(offset, size)
	if err != nil {
		return 0, err
	}
	return int(s), nil
}

// ParseTime takes a bit offset and size and converts the binary
// number produced from that substring slice into a time.Time.
func (b BitString) ParseTime(offset, size int) (time.Time, error) {
	var s, err = b.ParseInt64(offset, size)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(s/dsPerS, s%dsPerS*nsPerDs), nil
}

// ParseBitList takes a bit offset and size which specify a range
// of bits in the BitString's value which represent an ordered list
// of bits representing purposes as defined in the IAB spec.
// More on the purposes here: https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#purposes-features.
// The resulting map's keys represent the purposes allowed for this user.
func (b BitString) ParseBitList(offset, size int) (map[int]bool, error) {
	if len(b.value)-1 < offset+size {
		return nil, errOutOfRange
	}
	var purposes = make(map[int]bool)
	for i, v := range b.value[offset:(offset + size)] {
		if v == '1' {
			purposes[i+1] = true
		}
	}
	return purposes, nil
}

// ParseBool returns a bool representing the bit at the
// passed offset.
func (b BitString) ParseBool(offset int) (bool, error) {
	if len(b.value)-1 < offset {
		return false, errOutOfRange
	}
	return b.value[offset] == '1', nil
}

// ParseString take a bit offset and size which should represent
// size / 6 characters to be parsed. Each six bits is parsed into
// a letter and returned in a final string. ParseString will error
// if size is not divisible by 6.
func (b BitString) ParseString(offset, size int) (string, error) {
	if len(b.value)-1 < offset+size {
		return "", errOutOfRange
	}

	var numChars = size / 6
	var retString []string

	if size%6 != 0 {
		return "", errWrongLength
	}
	for i := 0; i < numChars; i++ {
		str, _ := b.ParseInt64(offset+6*i, 6)
		retString = append(retString, string(str+65))
	}
	return strings.Join(retString, ""), nil
}
