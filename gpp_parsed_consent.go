package iabconsent

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type GppHeader struct {
	Type     int
	Version  int
	Sections []int
}

// GppParsedConsent is an empty interface since GPP will need to handle more consent structs
// than just the Multi-state Privacy Agreement structs.
type GppParsedConsent interface {
}

type GppSubSection struct {
	Gpc bool
}

type GppSubSectionTypes int

const (
	SubSectCore GppSubSectionTypes = iota
	SubSectGpc
)

// Each supported Section ID must have a Parsing Function and be added here to support a given section.
var parsingFunctions = map[int]func(string) (GppParsedConsent, error){
	7: ParseUsNational,
}

// ParseFunctionValue packages together the correct parsing function with the Section Value to be
// Parsed at some point in the future.
func ParseFunctionValue(f func(string) (GppParsedConsent, error), val string) func() (GppParsedConsent, error) {
	return func() (GppParsedConsent, error) { return f(val) }
}

// ParseGppHeader parses the first (and required) part of any GPP Consent String.
// It is used to read the Type, Version, and which sections are contained in the following string(s).
// Format is:
// Type	    Int(6)	Fixed to 3 as “GPP Header field”
// Version	Int(6)	Version of the GPP spec (version 1, as of Jan. 2023)
// Sections	Range(Fibonacci)	List of Section IDs that are contained in the GPP string.
func ParseGppHeader(s string) (*GppHeader, error) {
	// IAB's base64 conversion means a 6 bit grouped value can be converted to 8 bit bytes.
	// Any leftover bits <8 would be skipped in normal base64 decoding.
	// Therefore, pad with 6 '0's w/ `A` to ensure that all bits are decoded into bytes.
	var b, err = base64.RawURLEncoding.DecodeString(s + "A")
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

// ParseGpp takes a base64 Raw URL Encoded string which represents a GPP v1 string
// of the format {gpp header}~{section 1}[.{sub-section}][~{section n}]
// and returns each pair of section value and parsing function that should be used.
// The pairs are returned to allow more control over how parsing functions are applied.
func ParseGpp(s string) (map[int]func() (GppParsedConsent, error), error) {
	var gppHeader *GppHeader
	var err error
	// ~ separated fields. with the format {gpp header}~{section 1}[.{sub-section}][~{section n}]
	var segments = strings.Split(s, "~")
	if len(segments) < 2 {
		return nil, errors.New("not enough gpp segments")
	}

	gppHeader, err = ParseGppHeader(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "read gpp header")
	} else if len(segments[1:]) != len(gppHeader.Sections) {
		// Return early if sections in header do not match sections passed.
		return nil, errors.New("mismatch number of sections")
	}
	// Go through each section and add parsing function and section value to returned value.
	var gppFunctions = make(map[int]func() (GppParsedConsent, error), 0)
	for i := 1; i < len(segments); i++ {
		var parsingFunc func(string) (GppParsedConsent, error)
		var ok bool
		// Segments and Corresponding sections are off by 1.
		parsingFunc, ok = parsingFunctions[gppHeader.Sections[i-1]]
		if !ok {
			// Missing parsing function, quietly skip for now.
		} else {
			gppFunctions[gppHeader.Sections[i-1]] = ParseFunctionValue(parsingFunc, segments[i])
		}
	}
	return gppFunctions, nil
}

// ParseGppConsent takes a base64 Raw URL Encoded string which represents a GPP v1 string and
// returns a map of Section ID to ParsedConsents with consent parsed via a consecutive parsing.
func ParseGppConsent(s string) (map[int]GppParsedConsent, error) {
	var gppFuncs map[int]func() (GppParsedConsent, error)
	var err error
	gppFuncs, err = ParseGpp(s)
	if err != nil {
		return nil, err
	}
	var gppConsents = make(map[int]GppParsedConsent, len(gppFuncs))
	// Consecutively, go through each section and try to parse.
	for sId, gpp := range gppFuncs {
		var consent GppParsedConsent
		var consentErr error
		consent, consentErr = gpp()
		if consentErr != nil {
			// If an error, quietly do not add teh consent value to map.
		} else {
			gppConsents[sId] = consent
		}
	}
	return gppConsents, nil
}

// ParseGppSubSection parses the subsections that may be appended to GPP sections after a `.`
// Currently, GPC is the only subsection, so we only have a single Subsection parsing function.
// In the future, Section IDs may need their own SubSection parser.
func ParseGppSubSection(s string) (*GppSubSection, error) {
	var b, err = base64.RawURLEncoding.DecodeString(s + "A")
	if err != nil {
		return nil, errors.Wrap(err, "parse gpp subsection string")
	}
	var r = NewConsentReader(b)
	var gppSub = new(GppSubSection)
	var subType int
	subType, err = r.ReadInt(2)
	if err != nil {
		return nil, errors.Wrap(err, "parse gpp subsection type")
	}
	var gppValue bool
	if GppSubSectionTypes(subType) == SubSectGpc {
		gppValue, err = r.ReadBool()
		if err != nil {
			return nil, errors.Wrap(err, "parse gpp subsection gpc bool")
		}
		gppSub.Gpc = gppValue
	}
	return gppSub, nil
}
