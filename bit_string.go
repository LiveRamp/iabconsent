package iabconsent

import (
	"fmt"
	"strconv"
)

type BitString struct {
	value string
}

func (b BitString) ParseInt64(offset, size int) (int64, error) {
	return strconv.ParseInt(b.value[offset:(offset+size)], 2, 64)
}

func (b BitString) ParseInt(offset, size int) (int, error) {
	var s, err = strconv.ParseInt(b.value[offset:(offset+size)], 2, 0)
	return int(s), err
}

func (b BitString) ParseBitList(offset, size int) map[int]interface{} {
	var purposes = make(map[int]interface{})
	for i, v := range b.value[offset:(offset + size)] {
		if v == '1' {
			purposes[i+1] = true
		}
	}
	return purposes
}

func (b BitString) ParseBit(offset int) bool {
	return b.value[offset] == '1'
}

func (b BitString) ParseString(offset, size int) (string, error) {
	var numChars = size / 6
	var retString = ""

	if size%6 != 0 {
		return "", fmt.Errorf("bit string length must be multiple of 6")
	}
	for i := 0; i < numChars; i++ {
		str, _ := b.ParseInt64(offset+6*i, 6)
		retString = retString + string(str+65)
	}
	return retString, nil
}

func ParseBytes(b []byte) (bs BitString) {
	for _, s := range b {
		bs.value = bs.value + fmt.Sprintf("%08b", s)
	}
	return
}
