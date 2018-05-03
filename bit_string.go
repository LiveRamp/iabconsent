package iabconsent

import (
	"fmt"
	"strconv"
)

// bitString is a simple struct which has only one field, value.
// The value is expected to be a string containing only ones and zeros
// which represent a Vendor Consent String as defined by the IAB spec
// found here: https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#vendor-consent-string-format-.
// The type is defined to enable common operations needed to parse
// the string which are defined below.
type bitString struct {
	value string
}

// parseBytes takes in a []byte |b| and returns a bitString |bs|
// who's value is the concatenation of the 8 bit binary representation
// of each element of |b|.
func parseBytes(b []byte) (bs bitString) {
	for _, s := range b {
		bs.value = bs.value + fmt.Sprintf("%08b", s)
	}
	return
}

// parseInt64 takes a bit offset and size and converts the binary
// number produced from that substring slice into an int64.
func (b bitString) parseInt64(offset, size int) (int64, error) {
	return strconv.ParseInt(b.value[offset:(offset+size)], 2, 64)
}

// parseInt takes a bit offset and size and converts the binary
// number produced from that substring slice into an int.
func (b bitString) parseInt(offset, size int) (int, error) {
	var s, err = b.parseInt64(offset, size)
	return int(s), err
}

// parseBitList takes a bit offset and size which specify a range
// of bits in the bitString's value which represent an ordered list
// of bits representing purposes as defined in the IAB spec.
// More on the purposes here: https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#purposes-features.
// The resulting map's keys represent the purposes allowed for this user.
func (b bitString) parseBitList(offset, size int) map[int]interface{} {
	var purposes = make(map[int]interface{})
	for i, v := range b.value[offset:(offset + size)] {
		if v == '1' {
			purposes[i+1] = true
		}
	}
	return purposes
}

// parseBit returns a bool representing the bit at the
// passed offset.
func (b bitString) parseBit(offset int) bool {
	return b.value[offset] == '1'
}

// parseString take a bit offset and size which should represent
// size / 6 characters to be parsed. Each six bits is parsed into
// a letter and returned in a final string. parseString will error
// if size is not divisible by 6.
func (b bitString) parseString(offset, size int) (string, error) {
	var numChars = size / 6
	var retString = ""

	if size%6 != 0 {
		return "", fmt.Errorf("bit string length must be multiple of 6")
	}
	for i := 0; i < numChars; i++ {
		str, _ := b.parseInt64(offset+6*i, 6)
		retString = retString + string(str+65)
	}
	return retString, nil
}
