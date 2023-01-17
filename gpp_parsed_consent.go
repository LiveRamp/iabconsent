package iabconsent

import (
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
)

type GppHeader struct {
	Type     int
	Version  int
	Sections []int
}

// ParseGppHeader parses the first (and required) part of any GPP Consent String.
// It is used to read the Type, Version, and which sections are contained in the following string(s).
// Format is:
// Type	    Int(6)	Fixed to 3 as “GPP Header field”
// Version	Int(6)	Version of the GPP spec (version 1, as of Jan. 2023)
// Sections	Range(Fibonacci)	List of Section IDs that are contained in the GPP string.
func ParseGppHeader(s string) (*GppHeader, error) {
	var b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "parse gpp header consent string")
	}

	var r = NewConsentReader(b)

	var g = &GppHeader{}
	g.Type, _ = r.ReadInt(6)
	if g.Type != 3 {
		return nil, errors.New("wrong gpp header type " + fmt.Sprint(g.Type))
	}
	g.Version, _ = r.ReadInt(6)
	if g.Version != 1 {
		return nil, errors.New("unsupported gpp version " + fmt.Sprint(g.Version))
	}
	g.Sections, _ = r.ReadFibonacciRange()
	return g, r.Err
}
